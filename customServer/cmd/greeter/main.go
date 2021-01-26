package main

import (
	"customServer/gen/restapi"
	"customServer/gen/restapi/operations"
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"log"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")


func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewGreeterAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	flag.Parse()
	server.Port = *portFlag

	api.GetGreetingHandler = operations.GetGreetingHandlerFunc(func(params operations.GetGreetingParams) middleware.Responder {
		name := swag.StringValue(params.Name)
		if name == "" {
			name = "World"
		}
		greeting := fmt.Sprintf("Hello, %s", name)
		return operations.NewGetGreetingOK().WithPayload(greeting)
	})


	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

