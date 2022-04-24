package middleware

import (
	"learning-k8s/internal/learning-k8s-02/config"
	"net/http"
)

func AddReqHeader2Resp(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for k := range r.Header {
			reqHeader := r.Header[k]
			for i := range reqHeader {
				w.Header().Add(k, reqHeader[i])
			}
		}

		w.Header().Add("VERSION", config.GetEnv("VERSION", "no version"))

		wrappedHandler.ServeHTTP(w, r)
	})
}
