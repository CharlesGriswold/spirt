package main

import (
	//~ "bytes"
	"bufio"
	"fmt"
	"github.com/mitchellh/go-wordwrap"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	file, err := os.Open("belief.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i int
	var buf, phrase, spurt string

	// This is where the magic happens.
	for scanner.Scan() {
		buf = scanner.Text()
		if strings.HasPrefix(buf, "#") {
			continue
		}
		i++
		if buf == "␞" {
			spurt += phrase
			i = 0
			phrase = ""
		} else if rand.Intn(i) == 0 {
			phrase = buf
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	spurt += phrase

	width, _, err := terminal.GetSize(0)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(wordwrap.WrapString(spurt, uint(width)))

}

