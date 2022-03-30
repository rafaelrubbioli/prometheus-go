package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// define you own custom metrics to be exported
var (
	myCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "access_counter",
		Help: "The total number of times /count was accessed",
	})
)

func main() {
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		myCounter.Inc()
		_, err := w.Write([]byte("incrementing..."))
		if err != nil {
			log.Fatal(err)
		}
	})

	fmt.Println("Listening on http://localhost:1234")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}
