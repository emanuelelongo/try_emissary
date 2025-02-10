package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

var bucketToServer = map[string]string{
	"bucket1": "pizza",
	"bucket2": "sushi",
	"bucket3": "pizza",
	"bucket4": "sushi",
	"bucket5": "pizza",
	"bucket6": "sushi",
	"bucket7": "pizza",
	"bucket8": "sushi",
}

type AuthResponse struct {
	Allowed bool              `json:"allowed"`
	Headers map[string]string `json:"headers,omitempty"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Printf("Error dumping request: %v", err)
		}

		log.Printf("Received request: %v", string(dump))

		host := r.Host
		bucketName := strings.Split(host, ".")[0]

		resp := AuthResponse{
			Allowed: true,
		}

		s3Server, exists := bucketToServer[bucketName]
		if exists {
			log.Printf("Routing to %v", s3Server)
			resp.Headers = map[string]string{
				"x-route-to": s3Server,
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	log.Printf("Starting server on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
