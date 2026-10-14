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
)

// pdpClient is the enforcement point's client of the cross-service PDP. It is a
// package var so tests can point it at an httptest server.
var pdpClient = authz.NewClientFromEnv()

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

	resp, err := pdpClient.Decide(ctx, req)
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
