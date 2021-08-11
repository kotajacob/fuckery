package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"git.sr.ht/~kota/fuckery"
)

var Version string

func usage() {
	log.Println("fuckery v" + Version)
	log.Fatal(`usage: fuckery [style]

styles:
	Strike
	Underline
	BoldSans
	BoldSerif
	ItalicSans
	ItalicSerif
	BoldItalicSans
	BoldItalicSerif
	Double
	Cursive
	Fraktur
	Zalgo NUMBER`)
}

func main() {
	var style string
	log.SetPrefix("")
	log.SetFlags(0)
	flag.Parse()
	args := flag.Args()
	if len(args) >= 1 {
		style = args[0]
	} else {
		usage()
	}
	switch style {
	case "Strike":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(fuckery.Strike(scanner.Text()))
		}
	case "Underline":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(fuckery.Underline(scanner.Text()))
		}
	case "BoldSans":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(fuckery.BoldSans(scanner.Text()))
		}
	case "BoldSerif":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(fuckery.BoldSerif(scanner.Text()))
		}
	case "ItalicSans":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(fuckery.ItalicSans(scanner.Text()))
		}
	case "ItalicSerif":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(fuckery.ItalicSerif(scanner.Text()))
		}
	case "BoldItalicSans":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(fuckery.BoldItalicSans(scanner.Text()))
		}
	case "BoldItalicSerif":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(fuckery.BoldItalicSerif(scanner.Text()))
		}
	case "Double":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(fuckery.Double(scanner.Text()))
		}
	case "Cursive":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(fuckery.Cursive(scanner.Text()))
		}
	case "Fraktur":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(fuckery.Fraktur(scanner.Text()))
		}
	case "Zalgo":
		var zal int
		var err error
		if len(args) >= 2 {
			zal, err = strconv.Atoi(args[1])
			if err != nil {
				usage()
			}
		} else {
			usage()
		}
		scanner := bufio.NewScanner(os.Stdin)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for scanner.Scan() {
			fmt.Println(fuckery.Zalgo(scanner.Text(), r, zal))
		}
	default:
		log.Println("unknown style")
		usage()
	}
}
