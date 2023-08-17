package forwarder

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"openid-connect-forwarder/internal/target"
	"openid-connect-forwarder/internal/token"
)

type OpenIdForwarder struct {
	srv *target.Server
	tm  *token.TokenManager
}

func (f *OpenIdForwarder) Handle(res http.ResponseWriter, req *http.Request) {
	req.Header.Add("Test", f.tm.Token())
	url, _ := url.Parse(f.srv.Server)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(res, req)
}

func NewOpenIdForwarder(srv *target.Server, tm *token.TokenManager) *OpenIdForwarder {
	return &OpenIdForwarder{srv: srv, tm: tm}
}
