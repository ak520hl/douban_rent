package douban

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

var (
	Url = "https://www.douban.com/group/637628/"
)
type Post struct {
	Title string
	Author string
	LastPostTime string
	LastRePost string
	Url string
}

// get all posts
func (c *Client) Posts() (posts []*Post) {

	req, err := c.NewRequest(http.MethodGet, Url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := c.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	name := doc.Find("div.group-desc h1").Text()
	fmt.Println(name)
	doc.Find("table.olt tbody tr").Each(func(i int, s *goquery.Selection) {
		post := &Post{}
		s.Find("td").Each(func(ii int, ss *goquery.Selection) {
			switch ii {
			case 0: // title
				post.Title = strings.TrimSpace(ss.Text())
				url, exist := ss.Find("a").Attr("href")
				if exist {
					// Url
					detail := c.PostDetail(url)
					Insert(detail)
					post.Url = url
				}
			case 1: // author
				post.Author = strings.TrimSpace(ss.Text())
			case 2: // lastRePost
				post.LastRePost = strings.TrimSpace(ss.Text())
			case 3: // lastPostTime
				post.LastPostTime = strings.TrimSpace(ss.Text())
			}
		})

		posts = append(posts, post)
	})
	return posts
}