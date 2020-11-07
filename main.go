package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {

	var set, numTag int
	flag.IntVar(&set, "set", 0, "-set=1")
	flag.IntVar(&numTag, "num", 0, "-num=27")
	flag.Parse()

	if flag.NFlag() < 1 {
		log.Printf("run with set property")
		flag.PrintDefaults()
		return
	}

	if numTag <= 0 {
		numTag = 27
	}

	data, err := ioutil.ReadFile("hashtags.txt")

	if err != nil {
		log.Printf("Error occured while reading file : %v", err)
	}

	sets := strings.Split(string(data), "-")

	setToShuffle := strings.TrimPrefix(strings.TrimSuffix(sets[set-1], fmt.Sprintln()), fmt.Sprintln())

	hashtags := strings.Split(setToShuffle, fmt.Sprintln())

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(hashtags), func(i, j int) {
		hashtags[i], hashtags[j] = hashtags[j], hashtags[i]
	})

	fmt.Printf("===============================Set %v=======================================\n", set)
	for i, ht := range hashtags {
		if i >= numTag-1 {
			break
		}
		fmt.Println(strings.TrimSpace(ht))
	}
	fmt.Println("============================================================================")

}
