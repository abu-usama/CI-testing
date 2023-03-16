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

	// owner := "abu-usama"
	// repo := "CI-testing"
	// headSHA := "1c5f316edd33bf68a0a8008f3efa77e3d6fa3163"
	// checkSuiteRequest := github.CreateCheckSuiteOptions{
	// 	HeadSHA: headSHA,
	// }
	// checkSuite, _, err := client.Checks.CreateCheckSuite(ctx, owner, repo, checkSuiteRequest)
	// if err != nil {
	// 	fmt.Printf("Error creating check suite: %v\n", err)
	// }
	// fmt.Printf("Check suite created: %v\n", checkSuite)

	// checkRun1 := &github.CreateCheckRunOptions{
	// 	Name:    "Check Run 1",
	// 	HeadSHA: "",
	// 	Status:  github.String("in_progress"),
	// }
	// checkRun2 := &github.CreateCheckRunOptions{
	// 	Name:    "Check Run 2",
	// 	HeadSHA: "headSHA",
	// 	Status:  github.String("completed"),
	// }

	// event, err := github.ParseWebHook(github.WebHookType(req.Request), payload)
	// if err != nil {
	// 	c.L.Errorf("error :%v", err)
	// }
	// switch event := event.(type) {
	// case *github.PushEvent:
	// 	c.L.Infof("action: %v", event.GetAction())
	// 	commitId := event.GetHeadCommit().ID
	// 	c.L.Infof("head-commit : %v", *commitId)
	// default:
	// 	c.L.Infof("unknown event : %V", event)
	// }

	// listCheckSuiteReq := github.ListCheckSuiteOptions{}
	// 	checkSuiteList, _, err := client.Checks.ListCheckSuitesForRef(ctx, owner, repo, ref, &listCheckSuiteReq)
	// 	if err != nil {
	// 		c.L.Errorf("Error listing check suite: %v\n", err)
	// 	}
	// 	jsonCS, _ := json.Marshal(checkSuiteList)
	// 	c.L.Infof("Check suite created: %v", string(jsonCS))

}
