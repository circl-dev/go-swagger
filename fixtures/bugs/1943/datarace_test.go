// +build ignore

package main

import (
	"log"
	"testing"
	"time"

	"github.com/circl-dev/go-swagger/fixtures/bugs/1943/restapi"
	"github.com/circl-dev/go-swagger/fixtures/bugs/1943/restapi/operations"
	"github.com/circl-dev/loads"
)

func Test_DataRace(t *testing.T) {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewPseudoServiceAPI(swaggerSpec)
	server := restapi.NewServer(api)

	server.ConfigureFlags()

	server.ConfigureAPI()

	go func() {
		time.Sleep(1 * time.Second)
		server.Shutdown()
	}()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
