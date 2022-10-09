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
    more bool
}

func(e *experiment) firstPersonTaste() bool{
    return e.people[0]
}

func NewExperiment(peopleNum int) experiment{
    people := make([]bool, peopleNum)
    numberOfPeopleWhoLike := rand.Intn(peopleNum+1)
    randomPermutationSlice := rand.Perm(peopleNum)
    for _, num := range randomPermutationSlice[:numberOfPeopleWhoLike] {
        people[num] = true
    }
    more := false
    if numberOfPeopleWhoLike > (peopleNum / 2){
        more = true
    }
    resp := experiment{
        people: people,
        more : more,
    }
    return resp
}

func main() {
    rightGuess := 0
    for i := 0 ; i < 10000; i++{
        exp := NewExperiment(10000)
        guess := exp.people[0]
        if exp.more == guess{
            rightGuess++
        }
    }
    fmt.Print(rightGuess)
}
