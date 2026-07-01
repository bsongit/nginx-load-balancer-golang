package main

import (
	"fmt"
	"net/http"
	"sync"
)

func createServer(port string, serverName string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Response from %s running on port %s\n", serverName, port)
	})

	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	servers := []struct {
		port string
		name string
	}{
		{"8081", "Server 1"},
		{"8082", "Server 2"},
		{"8083", "Server 3"},
	}

	for _, s := range servers {
		go func(port, name string) {
			defer wg.Done()
			srv := createServer(port, name)
			fmt.Printf("Starting %s on port %s...\n", name, port)
			if err := srv.ListenAndServe(); err != nil {
				fmt.Printf("Server on port %s closed: %v\n", port, err)
			}
		}(s.port, s.name)
	}

	wg.Wait()
}
