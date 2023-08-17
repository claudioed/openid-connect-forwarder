package main

import (
	"log"
	"net/http"
	"openid-connect-forwarder/internal/forwarder"
	"openid-connect-forwarder/internal/target"
	"openid-connect-forwarder/internal/token"
	"os"
)

func main() {
	target := &target.Server{Server: os.Getenv("SERVER_URL")}
	tm := token.NewTokenManager()
	fwd := forwarder.NewOpenIdForwarder(target, tm)
	http.HandleFunc("/", fwd.Handle)
	log.Fatal(http.ListenAndServe(":"+"9999", nil))
}
