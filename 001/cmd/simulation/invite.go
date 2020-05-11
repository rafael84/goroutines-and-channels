package main

import (
	"log"

	"github.com/rafael84/goroutines-and-channels/001/random"
)

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
			Registration: &Registration{
				Name: name,
				Age:  optionalAge,
			},
		}
		log.Println("invited", name)
	}
}
