package main

import "log"

func processSystemRequests() {
	for systemRequest := range systemRequestsChannel {
		log.Printf("processing system request: %v\n", systemRequest)
		processSystemRequest(systemRequest)
		log.Printf("system request processed: %v\n", systemRequest)
	}
}

func processSystemRequest(systemRequest *Request) {
	if systemRequest.Registration.Age == nil {
		go requestCustomerToEnterAge(systemRequest)

	} else if *systemRequest.Registration.Age < 18 {
		go rejectCustomerTooYoung(systemRequest)

	} else if systemRequest.Registration.Docs == nil || len(systemRequest.Registration.Docs) == 0 {
		go requestCustomerDocs(systemRequest)

	} else {
		go requestReviewOfDocs(systemRequest)
	}
}

func requestCustomerToEnterAge(systemRequest *Request) {
	log.Printf("requesting customer to enter age: %v\n", systemRequest)
	customerRequest := systemRequest.WithType(AgeIsMissing)
	customerRequestsChannel <- customerRequest
	log.Printf("customer request sent: %v\n", customerRequest)
}

func rejectCustomerTooYoung(systemRequest *Request) {
	log.Printf("customer is too young, rejecting request: %v\n", systemRequest)
	rejectedRequest := systemRequest.WithType(TooYoung)
	rejectedRequestsChannel <- rejectedRequest
	log.Printf("request rejected: %v\n", rejectedRequest)
}

func requestCustomerDocs(systemRequest *Request) {
	customerRequest := systemRequest.WithType(DocsRequired)
	log.Printf("requesting customer to send docs: %v\n", customerRequest)
	customerRequestsChannel <- customerRequest
	log.Printf("customer request sent: %v\n", customerRequest)
}

func requestReviewOfDocs(systemRequest *Request) {
	reviewerRequest := systemRequest.WithType(ReviewDocs)
	log.Printf("requesting reviewer to review the docs: %v\n", reviewerRequest)
	reviewerRequestsChannel <- reviewerRequest
	log.Printf("reviewer request sent: %v\n", reviewerRequest)
}
