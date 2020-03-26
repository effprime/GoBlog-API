package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"goblog/internal/models"
	"time"

	_ "github.com/lib/pq"
)

func NewPost(post models.Post) error {
	db := getDatabase()
	_, err := db.Exec(
		"INSERT INTO public.post (title, posted, body) VALUES ($1, $2, $3)",
		post.Title,
		time.Now(),
		post.Body,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(id int) error {
	post, err := GetPost(id)
	if err != nil {
		return err
	} else if post.Id == 0 {
		return errors.New("no post found")
	}
	db := getDatabase()
	_, err = db.Exec("DELETE FROM public.post WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetPost(id int) (models.Post, error) {
	var post models.Post
	db := getDatabase()
	result, err := db.Query("SELECT * FROM public.post where id=$1", id)
	if err != nil {
		return post, err
	}
	for result.Next() {
		err = result.Scan(&post.Title, &post.Posted, &post.Body, &post.Id)
		if err != nil {
			return post, err
		}
		break
	}
	return post, nil
}

func GetPosts(limit int) ([]models.Post, json.RawMessage, error) {
	var result *sql.Rows
	var err error

	db := getDatabase()
	if limit == 0 {
		result, err = db.Query("SELECT * FROM public.post")
	} else {
		result, err = db.Query("SELECT * FROM public.post LIMIT $1", limit)
	}
	if err != nil {
		return nil, nil, err
	}
	defer result.Close()

	posts := []models.Post{}
	for result.Next() {
		var post models.Post
		err := result.Scan(&post.Title, &post.Posted, &post.Body, &post.Id)
		if err != nil {
			return nil, nil, err
		}
		posts = append(posts, post)
	}

	serialized_posts, err := json.Marshal(posts)
	if err != nil {
		return nil, nil, err
	}
	x := string(serialized_posts)

	return posts, json.RawMessage(x), nil
}

func TestDatabaseConn() {
	db = getDatabase()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
}
