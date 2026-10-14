package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/interfacerproject/interfacer-dpp/internal/auth"
	"github.com/interfacerproject/interfacer-dpp/internal/authz"
)

func testCtx() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodDelete, "/dpp/01/", nil)
	return c, w
}

func TestAuthorizeWriteAllowProceeds(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(authz.Response{Decision: "allow", ReasonCode: "ok"})
	}))
	defer srv.Close()
	old := pdpClient
	pdpClient = &authz.Client{BaseURL: srv.URL, HTTP: srv.Client()}
	defer func() { pdpClient = old }()

	c, w := testCtx()
	ok := authorizeWrite(c, &auth.Principal{Pk: "PK"}, "dpp.delete", "R1", "cGF5")
	if !ok || c.IsAborted() || w.Code != http.StatusOK {
		t.Fatalf("allow should proceed: ok=%v aborted=%v code=%d", ok, c.IsAborted(), w.Code)
	}
}

func TestAuthorizeWriteDenyIs403(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(authz.Response{Decision: "deny", ReasonCode: "not_controller"})
	}))
	defer srv.Close()
	old := pdpClient
	pdpClient = &authz.Client{BaseURL: srv.URL, HTTP: srv.Client()}
	defer func() { pdpClient = old }()

	c, w := testCtx()
	ok := authorizeWrite(c, &auth.Principal{Pk: "PK"}, "dpp.delete", "R1", "cGF5")
	if ok || !c.IsAborted() || w.Code != http.StatusForbidden {
		t.Fatalf("deny must be 403+abort: ok=%v aborted=%v code=%d", ok, c.IsAborted(), w.Code)
	}
}

func TestAuthorizeWriteUnreachableFailsClosed(t *testing.T) {
	old := pdpClient
	pdpClient = &authz.Client{BaseURL: "", HTTP: http.DefaultClient} // empty URL => error
	defer func() { pdpClient = old }()

	c, w := testCtx()
	ok := authorizeWrite(c, &auth.Principal{Pk: "PK"}, "dpp.delete", "R1", "cGF5")
	if ok || w.Code != http.StatusForbidden {
		t.Fatalf("unreachable PDP must fail closed to 403: ok=%v code=%d", ok, w.Code)
	}
}
