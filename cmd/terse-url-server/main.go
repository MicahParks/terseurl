// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"fmt"
	"log"
	"os"

	"flag"

	"github.com/MicahParks/terse-URL/restapi"
	"github.com/MicahParks/terse-URL/restapi/operations"
	"github.com/go-openapi/loads"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	var server *restapi.Server // make sure init is called

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprint(os.Stderr, "  terse-url-server [OPTIONS]\n\n")

		title := "Terse URL"
		fmt.Fprint(os.Stderr, title+"\n\n")
		desc := "The Terse URL shortener."
		if desc != "" {
			fmt.Fprintf(os.Stderr, desc+"\n\n")
		}
		flag.CommandLine.SetOutput(os.Stderr)
		flag.PrintDefaults()
	}
	// parse the CLI flags
	flag.Parse()
	api := operations.NewTerseURLAPI(swaggerSpec)
	// get server with flag values filled out
	server = restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}