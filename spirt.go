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
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	
	var f string
	var n int
	
	if len(os.Args) == 1 {
		fmt.Println("Usage: " + os.Args[0] + " file [n]")
	} else if len(os.Args) == 2 {
		f = os.Args[1]
		n = 1
	} else {
		f = os.Args[1]
		var err error
		n, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}
	for i := 1; i <= n; i++ {
		do_stuff(f)
		if i <= n-1 {
			print("\n")
		}
	}
}

func do_stuff(f string) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i int
	var buf, phrase, spirt string

	// This is where the magic happens.
	for scanner.Scan() {
		buf = scanner.Text()
		if strings.HasPrefix(buf, "#") {
			continue
		}
		i++
		if buf == "␞" {
			spirt += phrase
			i = 0
			phrase = ""
		} else if rand.Intn(i) == 0 {
			phrase = buf
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	spirt += phrase

	width, _, err := terminal.GetSize(0)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	spirt = strings.Replace(spirt, "\\n", "\n", -1)
	fmt.Println(wordwrap.WrapString(spirt, uint(width)))
}
