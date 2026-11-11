package parser

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ContLeftS struct {
	Text string
	Link string
}

type ChangeList struct {
	Date      string // дата
	Pair      string // пара
	Group     string // группа
	Pod_group string // под группа
	Item      string // предмет
	Room      string // кабинет
	Teacher   string // преподователь
}

func ContLeftParse(url string) ([]ContLeftS, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d, %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ContLeft []ContLeftS
	doc.Find("body div#cont div#left a.left_group").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		href, exists := s.Attr("href")
		if exists && text != "" {
			ContLeft = append(ContLeft, ContLeftS{Text: text, Link: href})
		}
	})

	return ContLeft, nil
}

func ChangesParser(url string) ([]ChangeList, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d, %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ChangesList []ChangeList
	doc.Find("body div#cont form div#table table.lux tr.exams").Each(func(i int, tr *goquery.Selection) {
		var cells []string
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			cells = append(cells, strings.TrimSpace(td.Text()))
		})

		if len(cells) < 7 {
			return
		}

		ch := ChangeList{
			Date:      cells[0],
			Pair:      cells[1],
			Group:     cells[2],
			Pod_group: cells[3],
			Item:      cells[4],
			Room:      cells[5],
			Teacher:   cells[6],
		}

		ChangesList = append(ChangesList, ch)
	})

	return ChangesList, nil
}
