package authz

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/interfacerproject/interfacer-dpp/internal/auth"
)

func TestDecidePostsRequestAndParsesResponse(t *testing.T) {
	var gotPath string
	var gotReq Request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		_ = json.NewDecoder(r.Body).Decode(&gotReq)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Response{Decision: "allow", ReasonCode: "ok"})
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, HTTP: srv.Client()}
	req := BuildRequest("rid", "dpp.delete", "R1",
		&auth.Principal{Pk: "PK"}, Proof{Signature: "S", PayloadB64: "cGF5"})
	resp, err := c.Decide(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotPath != "/authz/decide" {
		t.Fatalf("path: got %q", gotPath)
	}
	if gotReq.Action != "dpp.delete" || gotReq.Binding.ParentID != "R1" {
		t.Fatalf("server received wrong request: %+v", gotReq)
	}
	if resp.Decision != "allow" {
		t.Fatalf("decision: got %q", resp.Decision)
	}
}

func TestDecideNon2xxIsError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	c := &Client{BaseURL: srv.URL, HTTP: srv.Client()}
	_, err := c.Decide(context.Background(), BuildRequest("r", "dpp.delete", "R", &auth.Principal{Pk: "P"}, Proof{}))
	if err == nil {
		t.Fatal("expected an error on 5xx so Interpret fails closed")
	}
}
