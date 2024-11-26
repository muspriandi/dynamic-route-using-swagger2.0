package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"api-gateway/constants"
	"api-gateway/domain"
	"api-gateway/handlers"
	"api-gateway/routes"
)

func main() {
	swaggerDir := "./swagger"
	serviceConfigs := make(map[string]domain.ServiceConfig)

	// Define Swagger Spec
	memberSpecPath := filepath.Join(swaggerDir, fmt.Sprintf("%s-swagger-spec.json", constants.MEMBER_SERVICE_CODE))
	memberPathPrefix := constants.MEMBER_SERVICE_PREFIX
	orderSpecPath := filepath.Join(swaggerDir, fmt.Sprintf("%s-swagger-spec.json", constants.ORDER_SERVICE_CODE))
	orderPathPrefix := constants.ORDER_SERVICE_PREFIX

	// Load Swagger specs for multiple services
	serviceConfigs[constants.MEMBER_SERVICE_CODE] = routes.LoadSwaggerSpec(
		constants.MEMBER_SERVICE_CODE,
		memberSpecPath,
		memberPathPrefix,
	)
	serviceConfigs[constants.ORDER_SERVICE_CODE] = routes.LoadSwaggerSpec(
		constants.ORDER_SERVICE_CODE,
		orderSpecPath,
		orderPathPrefix,
	)

	// Set the loaded service configurations in the handler
	handlers.SetServiceConfigs(serviceConfigs)

	// Start the API Gateway server
	http.HandleFunc("/", handlers.Handler)
	log.Printf("API Gateway running on :%d\n", constants.SERVICE_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", constants.SERVICE_PORT), nil))
}
