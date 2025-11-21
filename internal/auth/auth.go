package auth

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	zenroom "github.com/dyne/Zenroom/bindings/golang/zenroom"
)

//go:embed zenflows-crypto/src/verify_graphql.zen
var VERIFY string

// Input and output of sign_graphql.zen
type ZenroomData struct {
	Gql            string `json:"gql"`
	EdDSASignature string `json:"eddsa_signature"`
	EdDSAPublicKey string `json:"eddsa_public_key"`
}

type ZenroomResult struct {
	Output []string `json:"output"`
}

func (data *ZenroomData) VerifyDid() error {
	baseUrl := os.Getenv("BASE_DID_URL")
	if baseUrl == "" {
		baseUrl = "https://explorer.did.dyne.org/details/"
	}
	context := os.Getenv("DID_CONTEXT_PATH")
	if context == "" {
		context = "did:dyne:ifacer"
	}
	url := fmt.Sprintf("%s%s:%s", baseUrl, context, data.EdDSAPublicKey)
	log.Printf("Fetching %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("problem fetching DID, status: %d", resp.StatusCode)
	}
	return nil
}

// Used to verify the signature with `zenflows-crypto`
func (data *ZenroomData) IsAuth() error {
	var err error

	jsonData, _ := json.Marshal(data)

	// Verify the signature
	result, success := zenroom.ZencodeExec(VERIFY, "", string(jsonData), "")
	if !success {
		return errors.New(result.Logs)
	}
	var zenroomResult ZenroomResult
	err = json.Unmarshal([]byte(result.Output), &zenroomResult)
	if err != nil {
		return err
	}
	if zenroomResult.Output[0] != "1" {
		return errors.New("signature is not authentic")
	}
	return nil
}