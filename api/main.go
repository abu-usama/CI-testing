package api

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func Main(init bool) {
	if init {
		flag.Parse()
	}

	listenUrl := os.Getenv("API_LISTEN_URL")
	srvr := http.NewServeMux()
	srvr.HandleFunc("/", MainHandler)
	fmt.Printf("listening on %s", listenUrl)
	http.ListenAndServe(listenUrl, srvr)
}

func MainHandler(w http.ResponseWriter, httpreq *http.Request) {

}