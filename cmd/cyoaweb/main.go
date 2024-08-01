package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"github.com/nandotech/cyoa"
)

func main() {
	filename := flag.String("file", "../../gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story %s\n", *filename)

	f, err := os.Open(*filename)
	if err!= nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err!= nil{
		fmt.Println("not working")
	}

	fmt.Printf("%+v\n", story)
}