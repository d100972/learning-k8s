package app

import (
	"fmt"
	"log"
	"net/http"
)

// SimpleBaseApp 最简单的 app 启动信息
type SimpleBaseApp struct {
	Port    string
	Version string
}

func NewSimpleBaseApp(port, version string) *SimpleBaseApp {
	return &SimpleBaseApp{
		Port:    port,
		Version: version,
	}
}

func (a *SimpleBaseApp) Start() {

	addr := fmt.Sprintf(":%s", a.Port)
	log.Printf("Starting app on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
