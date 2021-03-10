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
	readers, err := ctx.ListReaders()
	if err != nil {
		log.Fatalln(err)
	}
	if len(readers) < 1 {
		log.Fatalln("no readers")
	}
	reader := readers[0]
	closeChan := make(chan struct{}, 1)
	errors := reader.WaitForChanges(closeChan)
	for {
		err = <-errors
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Println("change")
		}
	}
}
