package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"api-gateway/domain"
)

var serviceConfigs map[string]domain.ServiceConfig

func SetServiceConfigs(configs map[string]domain.ServiceConfig) {
	serviceConfigs = configs
}

func Handler(w http.ResponseWriter, r *http.Request) {
	for _, service := range serviceConfigs {
		if strings.HasPrefix(r.URL.Path, service.Prefix) {
			strippedPath := strings.TrimPrefix(r.URL.Path, service.Prefix)
			for _, route := range service.Routes {
				if strippedPath == route.Path && r.Method == route.Method {
					proxyURL := fmt.Sprintf("%s%s", service.BaseURL, strippedPath)

					// Log the proxying URL
					log.Printf("[REST][%s] Endpoint: %s\n", route.Method, proxyURL)

					// Create a new HTTP request to forward the incoming request to the target service
					req, err := http.NewRequest(r.Method, proxyURL, r.Body)
					if err != nil {
						log.Printf("Error creating request: %v", err)
						http.Error(w, "Bad Gateway", http.StatusBadGateway)
						return
					}

					// Copy the headers from the incoming request to the new request
					for name, values := range r.Header {
						for _, value := range values {
							req.Header.Add(name, value)
						}
					}

					// Send the request to the backend service
					client := &http.Client{}
					resp, err := client.Do(req)
					if err != nil {
						log.Printf("Error forwarding request: %v", err)
						http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
						return
					}
					defer resp.Body.Close()

					// Copy the response headers and body from the service response
					for name, values := range resp.Header {
						for _, value := range values {
							w.Header().Add(name, value)
						}
					}
					// Set the status code and write the body
					w.WriteHeader(resp.StatusCode)
					body, _ := io.ReadAll(resp.Body)
					w.Write(body)
					return
				}
			}
		}
	}

	// If no route matched, return 404
	http.NotFound(w, r)
}
