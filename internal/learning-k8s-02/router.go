package learning_k8s_02

import (
	"learning-k8s/internal/learning-k8s-02/controller"
	"learning-k8s/internal/learning-k8s-02/pkg/middleware"
	"net/http"
)

func LoadAPI() {
	middlewares := []middleware.Middleware{
		middleware.LogReq,
		middleware.AddReqHeader2Resp,
	}
	http.Handle("/healthz", middleware.ApplyMiddlewares(controller.Healthz, middlewares...))
}
