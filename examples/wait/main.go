package main

import (
	"log"

	"github.com/LLKennedy/go-card/smartcard"
)

func main() {
	ctx, err := smartcard.EstablishContext()
	if err != nil {
		log.Fatalln(err)
	}
	closeChan := make(chan struct{}, 1)
	results, errors := ctx.WaitForChanges(closeChan)
	for {
		select {
		case res := <-results:
			log.Printf("%#v\n", res)
		case err := <-errors:
			log.Printf("error: %v\n", err)
		}
	}
}
