package main

import (
	"common-etl/pipelines"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Environment variable PORT is not set")
	}

	done := make(chan bool)

	pipeline := pipelines.NewPipeline()
	pipeline.Start()

	// health check
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "health")
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}

	<-done
	log.Println("Main goroutine is terminated")
}
