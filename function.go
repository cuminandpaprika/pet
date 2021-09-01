// Package p contains an HTTP Cloud Function.
package p

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, GetRandomPet())
}

func GetRandomPet() string {
	rand.Seed(int64(time.Now().Nanosecond()))
	answers := []string{
		"pony",
		"dog",
		"cat",
		"gecko",
		"snake",
		"rabbit",
		"hamster",
		"kangaroo",
		"guinea pig",
		"mouse",
		"parrot",
		"goldfish",
		"lion",
	}
	return answers[rand.Intn(len(answers))]
}
