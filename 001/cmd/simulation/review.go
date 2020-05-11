package main

import "log"

func processReviewerRequests() {
	for reviewerRequest := range reviewerRequestsChannel {
		processReviewerRequest(reviewerRequest)
	}
}

func processReviewerRequest(reviewerRequest *Request) {
	switch reviewerRequest.Type {
	case ReviewDocs:
		go docsReviewed(reviewerRequest)
	}
}

func docsReviewed(reviewerRequest *Request) {
	log.Printf("processing reviewer request: %v\n", reviewerRequest)
	doc := reviewerRequest.Registration.Docs[0]
	if doc.Number%7 == 0 || doc.Number%9 == 0 {
		invalidDocument(reviewerRequest)
	} else {
		validDocument(reviewerRequest)
	}
	log.Printf("reviewer request processed: %v\n", reviewerRequest)
}

func invalidDocument(reviewerRequest *Request) {
	rejectedRequest := reviewerRequest.WithType(InvalidDoc)
	log.Printf("invalid document, rejecting request: %v\n", rejectedRequest)
	rejectedRequestsChannel <- rejectedRequest
	log.Printf("request rejected: %v\n", rejectedRequest)
}

func validDocument(reviewerRequest *Request) {
	approvedRequest := reviewerRequest.WithType(ValidDoc)
	log.Printf("valid document, approving request: %v\n", approvedRequest)
	approvedRequestsChannel <- approvedRequest
	log.Printf("request approved: %v\n", approvedRequest)
}
