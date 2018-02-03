package main

import (
	"flag"
	"log"
	"os"
)

const DefaultChannel string = "general"

var channel, token string

func init() {
	flag.StringVar(&channel, "channel", DefaultChannel, "")
	flag.StringVar(&token, "token", "", "")
}

func main() {
	flag.Parse()
	encoder := NewPostMessageEncoder(channel, token)
	reader, err := NewBufferedReader(os.Stdin, encoder)
	if err != nil {
		log.Fatal(err)
	}
	if err := reader.Parse(); err != nil {
		log.Fatal(err)
	}
	message := encoder.Message()
	creator, err := NewPostMessageRequestCreator(message)
	if err != nil {
		log.Fatal(err)
	}
	requester := NewRequester(creator)
	if _, err := requester.Request(); err != nil {
		log.Println(err)
	}
}
