package main

import (
	learningK8s02 "learning-k8s/internal/learning-k8s-02"
	"learning-k8s/internal/learning-k8s-02/config"
)

func main() {
	
	learningK8s02.NewK8s02App(
		config.GetEnv("PORT", "8080"),
		config.GetEnv("VERSION", "no version"),
	).Start()

}
