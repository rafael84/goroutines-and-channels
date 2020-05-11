package main

import (
	"log"

	"github.com/rafael84/goroutines-and-channels/001/random"
)

func processCustomerRequests() {
	for customerRequest := range customerRequestsChannel {
		log.Printf("processing customer request: %v\n", customerRequest)
		processCustomerRequest(customerRequest)
		log.Printf("customer request processed: %v\n", customerRequest)
	}
}

func processCustomerRequest(customerRequest *Request) {
	switch customerRequest.Type {
	case AgeIsMissing:
		go customerEnteredAge(customerRequest)

	case DocsRequired:
		go customerAttachedDocs(customerRequest)
	}
}

func customerEnteredAge(customerRequest *Request) {
	log.Printf("customer %v is entering their age...\n", customerRequest.Registration.Name)
	age := random.IntBetween(14, 99)
	systemRequest := &Request{
		Type: CustomerResponded,
		Registration: &Registration{
			Name: customerRequest.Registration.Name,
			Age:  &age,
		},
	}
	systemRequestsChannel <- systemRequest
	log.Printf("customer %v entered their age: %v\n", systemRequest.Registration.Name, age)
}

func customerAttachedDocs(customerRequest *Request) {
	log.Printf("customer %v is attaching their docs...\n", customerRequest.Registration.Name)
	docs := []Document{
		{
			Type:   random.Choice([]string{"CPF", "RG", "PASSPORT"}),
			Number: random.IntBetween(100000, 999999999),
		},
	}
	systemRequest := &Request{
		Type: CustomerResponded,
		Registration: &Registration{
			Name: customerRequest.Registration.Name,
			Age:  customerRequest.Registration.Age,
			Docs: docs,
		},
	}
	systemRequestsChannel <- systemRequest
	log.Printf("customer %v attached their docs: %v\n", systemRequest.Registration.Name, docs)
}
