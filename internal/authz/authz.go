// Package authz is the DPP enforcement point's client of the cross-service
// Policy Decision Point (zenflows docs/cross-service-authz-contract.md §4).
//
// BuildRequest and Interpret are pure so the request shape and the
// fail-closed decision mapping are unit-testable without a live PDP. The
// network call lives in client.go. The action is fixed by the route and the
// parent is derived server-side from the stored object by the caller — never
// taken from client payload (contract §3, deny.parent_from_binding_not_payload).
package authz

import "github.com/interfacerproject/interfacer-dpp/internal/auth"

// Proof carries the exact bytes the PEP verified (base64) and their signature,
// so the PDP re-verifies the user's own signature instead of trusting this
// service's assertion of identity (contract §4.3 property 2).
type Proof struct {
	Signature  string `json:"signature"`
	PayloadB64 string `json:"payload_b64"`
}

// Request is the contract request plus the proof block.
type Request struct {
	RequestID string `json:"request_id"`
	Action    string `json:"action"`
	Principal struct {
		Pk string `json:"pk"`
	} `json:"principal"`
	Binding struct {
		ParentID string `json:"parent_id"`
	} `json:"binding"`
	Proof Proof `json:"proof"`
}

// Response is the contract response rendered by the PDP.
type Response struct {
	Decision        string         `json:"decision"`
	ReasonCode      string         `json:"reason_code"`
	PermissionEpoch *int           `json:"permission_epoch"`
	Obligations     map[string]any `json:"obligations"`
	RequestID       string         `json:"request_id"`
	PolicyVersion   string         `json:"policy_version"`
}

// Outcome is what the handler enforces.
type Outcome struct {
	Allow      bool
	Status     int
	ReasonCode string
}

// BuildRequest assembles the wire request. action and parentID are supplied by
// the caller (route-fixed action; parent read from the stored object).
func BuildRequest(requestID, action, parentID string, p *auth.Principal, proof Proof) Request {
	var r Request
	r.RequestID = requestID
	r.Action = action
	r.Principal.Pk = p.Pk
	r.Binding.ParentID = parentID
	r.Proof = proof
	return r
}

// Interpret maps a PDP response (or a transport error) to an enforced outcome.
// Fail closed: only an explicit "allow" proceeds; everything else is 403.
func Interpret(resp Response, transportErr error) Outcome {
	if transportErr != nil {
		return Outcome{Allow: false, Status: 403, ReasonCode: "pdp_unreachable"}
	}
	if resp.Decision == "allow" {
		return Outcome{Allow: true, Status: 200, ReasonCode: resp.ReasonCode}
	}
	reason := resp.ReasonCode
	if reason == "" {
		reason = "denied"
	}
	return Outcome{Allow: false, Status: 403, ReasonCode: reason}
}
