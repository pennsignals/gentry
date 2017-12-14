package main

import (
	"flag"
	"log"
	"os"
)

const DefaultChannel string = "general"

var address, channel, token string

func init() {
	flag.StringVar(&address, "address", "", "")
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
	factory, err := NewPostMessageRequestFactory(address, message)
	if err != nil {
		log.Fatal(err)
	}
	requester := NewRequester(factory)
	if _, err := requester.Request(); err != nil {
		log.Println(err)
	}
}
