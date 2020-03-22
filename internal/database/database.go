package database

import (
	"database/sql"
	"fmt"
	"goblog/internal/models"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "goblog"
)

var db *sql.DB

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
	if limit == -1 {
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

func getDatabase() *sql.DB {
	if db == nil {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		return db
	}
	return db
}

func TestDatabaseConn() int {
	db = getDatabase()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	return 0
}
