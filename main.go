package main

import (
	"os"
	"fmt"
	"log"
)

func  main()  {
	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("Error", err)
	}

	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}

		fmt.Printf("read: %s\n", string(data[:n]))
	}
}