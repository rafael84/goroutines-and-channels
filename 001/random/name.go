package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var (
	names    = []string{"Anne", "Paul", "Bob", "Carl", "Mike", "Hanna", "Jessie", "James", "David", "Daniel"}
	surnames = []string{"Smith", "Johnson", "Willians", "Brown", "Jones", "Miller", "Garcia", "Mendes"}
)

func FullName() string {
	name := names[rand.Intn(len(names))]
	surname := surnames[rand.Intn(len(surnames))]
	return name + " " + surname
}
