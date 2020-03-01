package douban

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic"
	"log"
	"net/http"
	"os"
	"strings"
)

var client *elastic.Client
var eshost = "http://127.0.0.1:9200/"

type PostDetail struct {
	Title        string 	`json:"title"`
	PosterName   string		`json:"poster_name"`
	PostTime     string		`json:"post_time"`
	PosterAvatar string		`json:"post_avatar"`
	Content 	 string		`json:"content"`
}

func init() {
	errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(eshost))
	if err != nil {
		panic(err)
	}
}

func (c *Client) PostDetail(url string) (pd PostDetail) {
	request, err := c.NewRequest(http.MethodGet, url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := c.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	pd = PostDetail{}
	pd.Title = strings.TrimSpace(doc.Find("div#content h1").Text())
	pd.PosterName = doc.Find("div.topic-doc h3 span.from a").Text()
	pd.PosterAvatar, _ = doc.Find("div.user-face a img").Attr("src")
	pd.PostTime = doc.Find("div.topic-doc h3 span.color-green").Text()
	pd.Content, _ = doc.Find("div.topic-richtext").Html()
	return pd
}

func Insert(de PostDetail) {
	put, err := client.Index().Index("douban").
		Type("post_detail").
		BodyJson(de).
		Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("data id: %s", put.Id)
}
