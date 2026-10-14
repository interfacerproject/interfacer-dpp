package auth

import (
	"encoding/json"
	"os"
	"testing"
)

// The shared vector schema (docs/cross-service-authz/README.md). Only the
// fields the authentication floor needs are decoded.
type vectorFile struct {
	Layer   string   `json:"layer"`
	Vectors []vector `json:"vectors"`
}

type vector struct {
	Name  string `json:"name"`
	Doc   string `json:"doc"`
	Input struct {
		Path    map[string]string `json:"path"`
		Headers map[string]string `json:"headers"`
		Signed  *struct {
			Object struct {
				ID string `json:"id"`
			} `json:"object"`
		} `json:"signed_payload"`
		DidResolves    bool `json:"did_resolves"`
		SignatureValid bool `json:"signature_valid"`
	} `json:"input"`
	Expect struct {
		Outcome   string `json:"outcome"`
		Principal *struct {
			Did string `json:"did"`
			Pk  string `json:"pk"`
		} `json:"principal"`
		Reason string `json:"reason"`
	} `json:"expect"`
}

// inputFromVector maps a shared vector to the pure Decide input the way the
// gin layer will: signature present iff both DID headers are set; the signed
// object id is whatever the signed envelope covers (empty when there is no
// envelope, e.g. a body-less delete with no signature).
func inputFromVector(v vector) AuthInput {
	sigPresent := v.Input.Headers["did-sign"] != "" && v.Input.Headers["did-pk"] != ""
	signedID := ""
	if v.Input.Signed != nil {
		signedID = v.Input.Signed.Object.ID
	}
	return AuthInput{
		SigPresent:     sigPresent,
		Pk:             v.Input.Headers["did-pk"],
		DidResolves:    v.Input.DidResolves,
		SignatureValid: v.Input.SignatureValid,
		SignedObjectID: signedID,
		PathObjectID:   v.Input.Path["id"],
	}
}

func TestDecide_SharedVectors(t *testing.T) {
	raw, err := os.ReadFile("testdata/authn.json")
	if err != nil {
		t.Fatalf("read vectors: %v", err)
	}
	var vf vectorFile
	if err := json.Unmarshal(raw, &vf); err != nil {
		t.Fatalf("parse vectors: %v", err)
	}
	if vf.Layer != "authn" {
		t.Fatalf("expected authn layer, got %q", vf.Layer)
	}
	if len(vf.Vectors) == 0 {
		t.Fatal("no vectors loaded")
	}

	for _, v := range vf.Vectors {
		t.Run(v.Name, func(t *testing.T) {
			got := Decide(inputFromVector(v))

			wantAccept := v.Expect.Outcome == "accept"
			if got.Accept != wantAccept {
				t.Fatalf("%s: accept = %v, want %v (%s)", v.Name, got.Accept, wantAccept, v.Doc)
			}
			if got.Reason != v.Expect.Reason {
				t.Fatalf("%s: reason = %q, want %q", v.Name, got.Reason, v.Expect.Reason)
			}

			switch {
			case v.Expect.Principal == nil && got.Principal != nil:
				t.Fatalf("%s: expected no principal, got %+v", v.Name, got.Principal)
			case v.Expect.Principal != nil && got.Principal == nil:
				t.Fatalf("%s: expected principal %+v, got none", v.Name, v.Expect.Principal)
			case v.Expect.Principal != nil:
				if got.Principal.Did != v.Expect.Principal.Did || got.Principal.Pk != v.Expect.Principal.Pk {
					t.Fatalf("%s: principal = %+v, want did=%s pk=%s",
						v.Name, got.Principal, v.Expect.Principal.Did, v.Expect.Principal.Pk)
				}
			}
		})
	}
}
