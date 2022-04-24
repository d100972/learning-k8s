package learning_k8s_02

import (
	"learning-k8s/pkg/app"
	"log"
)

type K8s02App struct {
	app.SimpleBaseApp
}

func NewK8s02App(port, version string) *K8s02App {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return &K8s02App{
		*app.NewSimpleBaseApp(port, version),
	}
}

func (k02 *K8s02App) Start() {
	LoadAPI()
	k02.SimpleBaseApp.Start()
}
