package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seed = rand.NewSource(time.Now().UnixNano())
var random = rand.New(seed)

type experiment struct {
	people []bool
	more   bool
}

func (e *experiment) firstPersonTaste() bool {
	return e.people[0]
}

func NewExperiment(peopleNum int) experiment {
	people := make([]bool, peopleNum)
	numberOfPeopleWhoLike := rand.Intn(peopleNum + 1)
	randomPermutationSlice := rand.Perm(peopleNum)
	for _, num := range randomPermutationSlice[:numberOfPeopleWhoLike] {
		people[num] = true
	}
	more := false
	if numberOfPeopleWhoLike > (peopleNum / 2) {
		more = true
	}
	resp := experiment{
		people: people,
		more:   more,
	}
	return resp
}

func main() {
	numOfPeople := 10000
	numOfExperiments := 10000
	rightGuessByAlwaysTrue := 0
	rightGuessByFirstPerson := 0
	for i := 0; i < numOfExperiments; i++ {
		exp := NewExperiment(numOfPeople)
		guessByAlwaysTrue := true
		guessByFirstPerson := exp.people[0]
		if exp.more == guessByAlwaysTrue {
			rightGuessByAlwaysTrue++
		}
		if exp.more == guessByFirstPerson {
			rightGuessByFirstPerson++
		}
	}
	fmt.Println("Guessting by believing it is always true",
		rightGuessByAlwaysTrue,
		" probability : ",
		float64(rightGuessByAlwaysTrue)/float64(numOfPeople),
	)
	fmt.Println("Guessting by the first person's answer",
		rightGuessByFirstPerson,
		" probability : ",
		float64(rightGuessByFirstPerson)/float64(numOfPeople),
	)
}
