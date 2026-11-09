package handler

import (
	b64 "encoding/base64"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/interfacerproject/interfacer-dpp/internal/auth"
)

// authenticateWrite enforces the authentication floor of the cross-service
// authorization contract (zenflows docs/cross-service-authz-contract.md §2)
// on a state-changing request. It is the single place the four DPP write
// handlers — UpdateDPP, UpdateDPPStatus, DeleteDPP, DeleteAttachment — gain
// the signature check the create/upload/add-attachment handlers already have
// inline; a follow-up migrates those three onto this helper too.
//
// This is authentication only: it proves the caller holds a registered DID
// key whose signature is bound to this object. It does NOT authorize the
// action against the resource controller — that is the PDP step (contract
// §4), which needs the Zenflows side and is not wired yet. The reason is said
// plainly so no one mistakes authn for authz.
//
// On success it returns the raw request body (nil for a body-less request)
// and the verified principal. On failure it has already written 401 and
// aborted; the caller must return immediately.
//
// bodyless is true for requests with no JSON body (DELETE). The client signs
// the object id string, which the server reconstructs from the path, so the
// signature is bound to exactly this object by construction. For requests
// with a body (updates) the signature covers the body; full object-id binding
// for updates lands with the signed envelope (contract §2 shipping note).
func authenticateWrite(c *gin.Context, pathID string, bodyless bool) (body []byte, principal *auth.Principal, ok bool) {
	sig := c.Request.Header.Get("did-sign")
	pk := c.Request.Header.Get("did-pk")

	var signedBytes []byte
	if bodyless {
		signedBytes = []byte(pathID)
	} else {
		b, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not read request body"})
			c.Abort()
			return nil, nil, false
		}
		body = b
		signedBytes = b
	}

	in := auth.AuthInput{
		SigPresent:     sig != "" && pk != "",
		Pk:             pk,
		PathObjectID:   pathID,
		SignedObjectID: pathID,
	}
	if in.SigPresent {
		z := auth.ZenroomData{
			Gql:            b64.StdEncoding.EncodeToString(signedBytes),
			EdDSASignature: sig,
			EdDSAPublicKey: pk,
		}
		in.DidResolves = z.VerifyDid() == nil
		in.SignatureValid = z.IsAuth() == nil
	}

	res := auth.Decide(in)
	if !res.Accept {
		log.Printf("write auth rejected: reason=%s route=%s", res.Reason, c.FullPath())
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "authentication required for this operation",
			"reason": res.Reason,
		})
		c.Abort()
		return nil, nil, false
	}
	return body, res.Principal, true
}
