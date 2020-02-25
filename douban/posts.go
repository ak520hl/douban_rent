package douban

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

var (
	Url = "https://www.douban.com/group/106955/"
)
type Post struct {
	title string
	content string
	postUser string
	postUid int
}

// get all posts
func (c *Client) Posts() []*Post {

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
		switch i {
		case 1:
			s.Find("td").Each(func(ii int, ss *goquery.Selection) {
				switch ii {
				case 1:
					fmt.Println(ss.Text())
				case 2:
					fmt.Println(ss.Text())
				case 3:
					fmt.Println(ss.Text())
				case 4:
					fmt.Println(ss.Text())
				}
			})
		default:
		}
	})
	return []*Post{}
}