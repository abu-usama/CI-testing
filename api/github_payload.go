package api

import (
	"fmt"
	"net/http"

	"github.com/google/go-github/v50/github"
)

func getPayload(req *http.Request, resp *http.ResponseWriter) {
	// ctx := context.Background()
	// ts := oauth2.StaticTokenSource(
	// 	&oauth2.Token{AccessToken: ""},
	// )
	// tc := oauth2.NewClient(ctx, ts)

	// client := github.NewClient(tc)
	payload, err := github.ValidatePayload(req, []byte(""))
	if err != nil {
		fmt.Printf("error :%v", err)
	}
	fmt.Printf("payload :%v", string(payload))

	// // list all repositories for the authenticated user
	// repos, _, err := client.Repositories.List(ctx, "", nil)
	// if err != nil {
	// 	c.L.Errorf("failed to get list :%v", err)
	// 	return
	// }
	// c.L.Infof("orgs: %v", repos)
}
