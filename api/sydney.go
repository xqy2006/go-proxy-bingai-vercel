package api

import (
	"xqy2006/go-proxy-bingai/api/helper"
	"xqy2006/go-proxy-bingai/common"
	"net/http"
)

func Sydney(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
	common.NewSingleHostReverseProxy(common.BING_SYDNEY_URL).ServeHTTP(w, r)
}
