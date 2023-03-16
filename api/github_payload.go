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

	checkRun1 := &github.CreateCheckRunOptions{
		Name:    "Check Run 1",
		HeadSHA: "",
		Status:  github.String("in_progress"),
	}
	checkRun2 := &github.CreateCheckRunOptions{
		Name:    "Check Run 2",
		HeadSHA: "headSHA",
		Status:  github.String("in_progress"),
	}
}
