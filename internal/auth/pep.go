package auth

// Policy Enforcement Point — the authentication floor of the cross-service
// authorization contract (zenflows docs/cross-service-authz-contract.md §2).
//
// Decide is a pure function: its inputs are the *outcomes* of VerifyDid and
// IsAuth, not the network/Zenroom calls themselves. A thin gin layer
// (AuthenticatedPrincipal) computes those outcomes from the real calls, and
// the shared test vectors (testdata/authn.json, mirror of the zenflows
// docs/cross-service-authz/authn.json) drive them as stubs. That keeps the
// policy unit-testable without Mongo or zencode-exec.
//
// This is authentication only: it proves the caller holds a registered DID
// key and that the signature is bound to the object in the path. It does NOT
// authorize the action against the resource controller — that is the PDP step
// (contract §4), which needs the Elixir side and is not wired yet.

// Principal is the subject proven by a verified signature for a request. It is
// never derived from a header such as x-user-id.
type Principal struct {
	Did string
	Pk  string
}

// Reason codes mirror the "reason" values in the shared vectors.
const (
	ReasonOK             = "ok"
	ReasonMissingSig     = "missing_signature"
	ReasonBadSig         = "bad_signature"
	ReasonDidUnresolved  = "did_unresolved"
	ReasonObjectMismatch = "object_id_mismatch"
)

// DidContextPrefix is prepended to the public key to form the DID. It matches
// the default DID_CONTEXT_PATH used by VerifyDid ("did:dyne:ifacer").
const DidContextPrefix = "did:dyne:ifacer:"

// AuthInput is the decision input for the authentication floor.
type AuthInput struct {
	SigPresent     bool   // did-sign AND did-pk both present on the request
	Pk             string // the did-pk header value
	DidResolves    bool   // VerifyDid() == nil
	SignatureValid bool   // IsAuth() == nil
	SignedObjectID string // the object id the signature actually covers
	PathObjectID   string // the object id taken from the route
}

// AuthResult is the decision. Accept is true only when every check passes.
type AuthResult struct {
	Accept    bool
	Principal *Principal
	Reason    string
}

// Decide applies the authentication floor, failing closed: a missing or
// invalid signature, an unresolvable DID, or a signature bound to a different
// object than the path all yield a reject with no principal.
func Decide(in AuthInput) AuthResult {
	if !in.SigPresent {
		return AuthResult{Reason: ReasonMissingSig}
	}
	if !in.DidResolves {
		return AuthResult{Reason: ReasonDidUnresolved}
	}
	if !in.SignatureValid {
		return AuthResult{Reason: ReasonBadSig}
	}
	if in.SignedObjectID != in.PathObjectID {
		return AuthResult{Reason: ReasonObjectMismatch}
	}
	return AuthResult{
		Accept:    true,
		Principal: &Principal{Did: DidContextPrefix + in.Pk, Pk: in.Pk},
		Reason:    ReasonOK,
	}
}
