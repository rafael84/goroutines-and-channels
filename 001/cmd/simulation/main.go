package main

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/rafael84/goroutines-and-channels/001/random"
)

var (
	systemRequestsChannel   chan *Request
	customerRequestsChannel chan *Request
	reviewerRequestsChannel chan *Request

	approvedRequestsChannel chan *Request
	rejectedRequestsChannel chan *Request

	approvedRequestsCount uint64
	rejectedRequestsCount uint64
)

const (
	numberOfPeopleToInvite = 1000
)

func main() {
	log.Println("simulation started")

	// initialize the channels
	systemRequestsChannel = make(chan *Request, 5)
	customerRequestsChannel = make(chan *Request, 5)
	reviewerRequestsChannel = make(chan *Request, 5)
	approvedRequestsChannel = make(chan *Request, 100)
	rejectedRequestsChannel = make(chan *Request, 100)

	// invite some people
	go invitePeople()

	// process the invitations
	go processSystemRequests()
	go processCustomerRequests()
	go processReviewerRequests()
	go processApprovedRequests()
	go processRejectedRequests()

	log.Println("waiting the simulation to finish")
	for approvedRequestsCount+rejectedRequestsCount < numberOfPeopleToInvite {
		time.Sleep(50 * time.Millisecond)
	}

	log.Printf("simulation finished: approved=%d, rejected=%d\n", approvedRequestsCount, rejectedRequestsCount)
}

func invitePeople() {
	for i := 0; i < numberOfPeopleToInvite; i++ {
		name := random.FullName()

		var optionalAge *int
		if i%2 == 0 {
			age := random.IntBetween(14, 80)
			optionalAge = &age
		}

		log.Println("inviting", name)
		systemRequestsChannel <- &Request{
			Type: CustomerInvited,
			Data: &Registration{
				Name: name,
				Age:  optionalAge,
			},
		}
		log.Println("invited", name)
	}
}

func processApprovedRequests() {
	for approvedRequest := range approvedRequestsChannel {
		atomic.AddUint64(&approvedRequestsCount, 1)
		log.Printf("approved request: %v\n", approvedRequest)
	}
}

func processRejectedRequests() {
	for rejectedRequest := range rejectedRequestsChannel {
		atomic.AddUint64(&rejectedRequestsCount, 1)
		log.Printf("rejected request: %v\n", rejectedRequest)
	}
}
