# Shared authz test vectors (mirror)

`authn.json` is a **verbatim mirror** of
`docs/cross-service-authz/authn.json` in the zenflows repo (the source of
truth for the cross-service authorization contract). It is vendored here so
`pep_test.go` can run in CI without a cross-repo checkout.

Keep it in sync: when the canonical vectors change, copy the file over. The
test `TestDecide_SharedVectors` fails loudly if a vector references a case the
PEP does not handle.
