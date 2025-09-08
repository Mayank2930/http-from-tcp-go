package main

import (
	"os"
	"fmt"
	"log"
	"bytes"
	"io"
)

func getLinesChannel(f io.ReadCloser) <- chan string {
	out := make(chan string, 1)

	go func(){
		defer f.Close()
		defer close(out)
		str := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}

			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				str += string(data[:i])
				data = data[i + 1:]
				out <- str
				str = ""
			}

			str += string(data)
		}

		if len(str) != 0 {
			out <- str
		}
	}()

	return out
}

func  main()  {
	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("Error", err)
	}
	defer f.Close()

	lines := getLinesChannel(f)

	for line := range lines{
		fmt.Printf("read: %s\n", line)
	}
}