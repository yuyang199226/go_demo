package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	content, _ := ioutil.ReadFile("dumb.csv")

	reader := csv.NewReader(bytes.NewBuffer(content))
	_, err := reader.Read() // skip first line
	if err != nil {
		if err != io.EOF {
			log.Fatalln(err)
		}
	}
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
				break
			}
		}
		fmt.Println(line)
	}

}
