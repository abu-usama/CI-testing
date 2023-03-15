package api

import (
	"fmt"
	"net/http"

	"github.com/google/go-github/v50/github"
)

func handleGithubWh(req *http.Request, resp *http.ResponseWriter) {
	payload, err := github.ValidatePayload(req, []byte(""))
	if err != nil {
		fmt.Printf("error :%v", err)
	}
	fmt.Printf("payload :%v", string(payload))
}
