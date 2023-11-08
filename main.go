package main

import (
	"github.com/pedromartinsb/gosoap/rest"
	"github.com/pedromartinsb/gosoap/soap"
)

func main() {
	// Initialize SOAP request
	soap.InitializeSOAP()

	// Initialize RestAPI request
	rest.InitializeREST()
}
