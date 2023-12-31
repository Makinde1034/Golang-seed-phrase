package seedphrasegenerator

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Generates seed phrase with sentence length equal to length
*/
func GenerateSeedPhrase(length int) string {
	seedPhrase := ""
	rand.Seed(time.Now().UnixNano())

	for i:= 0; i < length; i++{
		randomInt := rand.Intn(40)
		seedPhrase = fmt.Sprintf("%s %s",seedPhrase,Words[randomInt])
	}

	return seedPhrase

}