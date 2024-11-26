package routes

import (
	"api-gateway/domain"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/spec"
)

// LoadSwaggerSpec loads and parses Swagger 2.0 specs into routes.
func LoadSwaggerSpec(serviceName, specPath, prefix string) domain.ServiceConfig {
	data, err := os.ReadFile(specPath)
	if err != nil {
		log.Fatalf("Error reading Swagger spec for %s: %v", serviceName, err)
	}

	// Parse the Swagger spec using go-openapi/spec
	swaggerSpec := &spec.Swagger{}
	err = swaggerSpec.UnmarshalJSON(data)
	if err != nil {
		log.Fatalf("Error parsing Swagger spec for %s: %v", serviceName, err)
	}

	var routes []domain.Route

	// Iterate over paths in the Swagger 2.0 spec
	for path, pathItem := range swaggerSpec.Paths.Paths {
		// Manually check for each HTTP method in the PathItem
		if pathItem.Get != nil {
			routes = append(routes, domain.Route{Path: path, Method: http.MethodGet})
			log.Printf("Loaded endpoint %s service: [%s]\t%s", serviceName, "GET", path)
		}
		if pathItem.Post != nil {
			routes = append(routes, domain.Route{Path: path, Method: http.MethodPost})
			log.Printf("Loaded endpoint %s service: [%s]\t%s", serviceName, "POST", path)
		}
		if pathItem.Put != nil {
			routes = append(routes, domain.Route{Path: path, Method: http.MethodPut})
			log.Printf("Loaded endpoint %s service: [%s]\t%s", serviceName, "PUT", path)
		}
		if pathItem.Delete != nil {
			routes = append(routes, domain.Route{Path: path, Method: http.MethodDelete})
			log.Printf("Loaded endpoint %s service: [%s]\t%s", serviceName, "DELETE", path)
		}
	}

	// Set scheme
	scheme := "http://"
	for _, sch := range swaggerSpec.Schemes {
		if sch == "https" {
			scheme = "https://"
			break
		}
	}

	// Return the routes to the service configuration
	return domain.ServiceConfig{
		BaseURL: fmt.Sprintf("%s%s", scheme, swaggerSpec.Host),
		Prefix:  prefix,
		Routes:  routes,
	}
}
