package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {

	data, err := ioutil.ReadFile("hashtags.txt")

	if err != nil {
		log.Printf("Error occured while reading file : %v", err)
	}

	hashtags := strings.Split(string(data), fmt.Sprintln())

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(hashtags), func(i, j int) {
		hashtags[i], hashtags[j] = hashtags[j], hashtags[i]
	})

	for _, ht := range hashtags {
		fmt.Println(ht)
	}

}
