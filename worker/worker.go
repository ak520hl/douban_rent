package worker

import (
	"ak520hl.cn/ak520hl/douban_rent/douban"
	"log"
)

type Worker struct {
	client *douban.Client
}

func New(client *douban.Client) *Worker {
	return &Worker{client: client}
}

func (w *Worker) Run() {
	//get all posts
	//get each post detail
	w.fetchPosts()
}

func (w *Worker) fetchPosts() []*douban.Post {
	log.Printf("Beggin fetch douban rent posts")
	defer log.Printf("End fetch douban rent posts")
	posts := w.client.Posts()
	for _, post := range posts {
		log.Printf("%s %s %s %s %s", post.Title, post.Author, post.LastRePost, post.LastPostTime, post.Url)
		// INSERT into elasticsearch
	}
	return posts
}
