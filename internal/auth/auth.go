package auth

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

func (data *ZenroomData) IsAuth() error {
	var err error

	jsonData, _ := json.Marshal(data)

	log.Printf("Calling Zenroom with data length: %d", len(jsonData))

	type zenResult struct {
		result  zenroom.ZenResult
		success bool
	}
	resultChan := make(chan zenResult, 1)

	go func() {
		result, success := zenroom.ZencodeExec(VERIFY, "", string(jsonData), "")
		resultChan <- zenResult{result, success}
	}()

	select {
	case res := <-resultChan:
		log.Printf("Zenroom completed, success: %v", res.success)
		if !res.success {
			log.Printf("Zenroom logs: %s", res.result.Logs)
			return errors.New(res.result.Logs)
		}

		var zenroomResult ZenroomResult
		err = json.Unmarshal([]byte(res.result.Output), &zenroomResult)
		if err != nil {
			log.Printf("Error unmarshaling zenroom output: %v", err)
			return err
		}
		if len(zenroomResult.Output) == 0 || zenroomResult.Output[0] != "1" {
			log.Printf("Signature not authentic, output: %v", zenroomResult.Output)
			return errors.New("signature is not authentic")
		}
		log.Println("Signature verified successfully")
		return nil

	case <-time.After(30 * time.Second):
		log.Println("Zenroom execution timed out after 30 seconds")
		return errors.New("signature verification timed out")
	}
}
