package main

import (
	"fmt"

	"github.com/simpleforce/simpleforce"
)

var (
	sfURL      = "https://techilaglobalservices47-dev-ed.my.salesforce.com" //instance url
	sfUser     = "ajith@techilaservices.com"                                //sf username
	sfPassword = "benny96@"                                                 // sf password
	sfToken    = "hsu5h2lQWfOby8Q5qOpCCeTH2"                                //security token in sf
)

func createClient() *simpleforce.Client { //function for authenticating sf
	client := simpleforce.NewClient(sfURL, simpleforce.DefaultClientID, simpleforce.DefaultAPIVersion)
	if client == nil {

		return nil
	}

	err := client.LoginPassword(sfUser, sfPassword, sfToken)
	if err != nil {

		return nil
	}

	q := "SELECT Name, BillingCity, BillingPostalCode, BillingCountry FROM Account WHERE BillingPostalCode='75251'"
	result, err := client.Query(q) // Note: for Tooling API, use client.Tooling().Query(q)
	if err != nil {
		// handle the error
		return client
	}

	for _, record := range result.Records {
		// access the record as SObjects.
		fmt.Println(record)
	}

	return client
}

func main() {

	createClient()
}
