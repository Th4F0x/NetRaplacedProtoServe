package main

import (
	"fmt"
	"log"

	"github.com/Th4Fox/NetRaplacedProto/internal/config"
	"github.com/Th4Fox/NetRaplacedProto/internal/parser"
)

func main() {
	Urls := config.Config("config.json")
	GroupList, err := parser.ContLeftParse(Urls.Url.MainPage)
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range GroupList {
		fmt.Printf("%s  %s\n", n.Text, n.Link)
	}

	Change, err := parser.ChangesParser(Urls.Url.RaplacedPage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("-----------------------------------------\n")
	for _, i := range Change {
		fmt.Printf("%s | %s | %s | %s | %s | %s | %s\n", i.Date, i.Pair, i.Group, i.Pod_group, i.Item, i.Room, i.Teacher)
	}
}
