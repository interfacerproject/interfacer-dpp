package authz

import (
	"errors"
	"testing"

	"github.com/interfacerproject/interfacer-dpp/internal/auth"
)

func TestBuildRequestUsesRouteActionAndServerParent(t *testing.T) {
	p := &auth.Principal{Did: "did:dyne:ifacer:PK", Pk: "PK"}
	req := BuildRequest("rid-1", "dpp.delete", "RESOURCE1", p,
		Proof{Signature: "SIG", PayloadB64: "cGF5"})

	if req.Action != "dpp.delete" {
		t.Fatalf("action: got %q want dpp.delete", req.Action)
	}
	if req.Binding.ParentID != "RESOURCE1" {
		t.Fatalf("parent: got %q want RESOURCE1", req.Binding.ParentID)
	}
	if req.Principal.Pk != "PK" {
		t.Fatalf("pk: got %q want PK", req.Principal.Pk)
	}
	if req.Proof.Signature != "SIG" || req.Proof.PayloadB64 != "cGF5" {
		t.Fatalf("proof not carried: %+v", req.Proof)
	}
	if req.RequestID != "rid-1" {
		t.Fatalf("request id: got %q", req.RequestID)
	}
}

func TestInterpretAllow(t *testing.T) {
	out := Interpret(Response{Decision: "allow", ReasonCode: "ok"}, nil)
	if !out.Allow || out.Status != 200 {
		t.Fatalf("allow: got %+v", out)
	}
}

func TestInterpretDenyIs403(t *testing.T) {
	out := Interpret(Response{Decision: "deny", ReasonCode: "not_controller"}, nil)
	if out.Allow || out.Status != 403 || out.ReasonCode != "not_controller" {
		t.Fatalf("deny: got %+v", out)
	}
}

func TestInterpretTransportErrorFailsClosed(t *testing.T) {
	out := Interpret(Response{}, errors.New("connection refused"))
	if out.Allow || out.Status != 403 {
		t.Fatalf("transport error must fail closed: got %+v", out)
	}
}

func TestInterpretUnknownDecisionFailsClosed(t *testing.T) {
	out := Interpret(Response{Decision: "maybe"}, nil)
	if out.Allow || out.Status != 403 {
		t.Fatalf("unknown decision must fail closed: got %+v", out)
	}
}
