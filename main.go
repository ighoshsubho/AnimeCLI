package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/xml"
	"os"
	"strconv"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Category    string `xml:"category"`
}

func main(){
	q := 5
	if len(os.Args) >= 2 {
		arg, _ := strconv.Atoi(os.Args[1])
		q = arg
	}

	res, err := http.Get("https://www.animenewsnetwork.com/all/rss.xml?ann-edition=in")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200{
		panic("Anime news not available!")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	
	xmlData := string(body)
	var rss Rss
	err1 := xml.Unmarshal([]byte(xmlData), &rss)
	if err1 != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, item := range rss.Channel.Items {
		if i >= q {
			break
		}
		fmt.Println("Title:", item.Title)
		fmt.Println("Link:", item.Link)
		fmt.Println("Description:", item.Description)
		fmt.Println("PubDate:", item.PubDate)
		fmt.Println("Category:", item.Category)
		fmt.Println()
	}
}