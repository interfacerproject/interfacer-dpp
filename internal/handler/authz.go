package handler

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/interfacerproject/interfacer-dpp/internal/auth"
	"github.com/interfacerproject/interfacer-dpp/internal/authz"
	"github.com/interfacerproject/interfacer-dpp/internal/model"
)

// keepAuthzParent pins the incoming DPP's ProductID to the stored value so an
// update cannot change the authorization parent it was just authorized
// against. ProductID is the binding the PDP decides on (contract §4.2 "parent
// from the binding, not payload"); letting the request body overwrite it would
// reparent the object on the very write being authorized.
func keepAuthzParent(incoming model.DigitalProductPassport, storedProductID string) model.DigitalProductPassport {
	incoming.ProductID = storedProductID
	return incoming
}

// pdpClient is the enforcement point's client of the cross-service PDP. It is a
// package var so tests can point it at an httptest server. It is resolved
// lazily (pdp()) rather than at package init, because ZENFLOWS_PDP_URL is only
// loaded from .env in main(), which runs after package initialization.
var pdpClient *authz.Client

func pdp() *authz.Client {
	if pdpClient == nil {
		pdpClient = authz.NewClientFromEnv()
	}
	return pdpClient
}

func newRequestID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

// authorizeWrite asks the PDP whether principal may perform action on the
// resource parentID (the stored DPP's ProductID, derived server-side).
// payloadB64 is the base64 of the exact bytes the PEP already verified, so the
// PDP re-verifies the user's own signature. On allow it returns true; on deny
// or any transport error it writes 403, aborts, and returns false (fail
// closed). This is authorization; authenticateWrite already proved identity.
func authorizeWrite(c *gin.Context, principal *auth.Principal, action, parentID, payloadB64 string) bool {
	sig := c.Request.Header.Get("did-sign")
	req := authz.BuildRequest(newRequestID(), action, parentID, principal,
		authz.Proof{Signature: sig, PayloadB64: payloadB64})

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	resp, err := pdp().Decide(ctx, req)
	out := authz.Interpret(resp, err)
	if !out.Allow {
		log.Printf("write authz denied: reason=%s action=%s parent=%s route=%s",
			out.ReasonCode, action, parentID, c.FullPath())
		c.JSON(http.StatusForbidden, gin.H{
			"error":  "not authorized for this operation",
			"reason": out.ReasonCode,
		})
		c.Abort()
		return false
	}
	return true
}
