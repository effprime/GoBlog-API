package database

import (
	"database/sql"
	"goblog/internal/models"

	_ "github.com/lib/pq"
)

func NewPost(post models.Post) {
	db := getDatabase()
	_, err := db.Exec(
		`INSERT INTO public.post 
		(title, posted, body)
		VALUES $1 $2 $3`,
		post.Title,
		post.Posted,
		post.Body,
	)
	if err != nil {
		panic(err)
	}

}

func GetPosts(limit int) []models.Post {
	var result *sql.Rows
	var err error
	db := getDatabase()
	if limit == 0 {
		result, err = db.Query("SELECT * FROM public.post LIMIT $1", limit)
	} else {
		result, err = db.Query("SELECT * FROM public.post")
	}
	if err != nil {
		panic(err)
	}
	defer result.Close()

	posts := []models.Post{}
	for result.Next() {
		var post models.Post
		err := result.Scan(&post.Title, &post.Posted, &post.Body, &post.Id)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	return posts
}

func TestDatabaseConn() int {
	db = getDatabase()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	return 0
}
