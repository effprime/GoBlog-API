package endpoints

import (
	"database/sql"
	"encoding/json"
	"goblog/internal/database"
	"goblog/internal/utils"
	"net/http"
	"time"
)

type Post struct {
	Title  string    `json:"title"`
	Posted time.Time `json:"posted"`
	Body   string    `json:"body"`
	Id     int       `json:"id"`
}

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	var post Post
	utils.GetPostRequestData(r, &post)
	post.Posted = time.Now()

	db := database.GetDatabase()
	sqlStatement := `
	INSERT INTO public.post (title, posted, body)
	VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, post.Title, post.Posted, post.Body)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func AllPostsHandler(w http.ResponseWriter, r *http.Request) {
	limit, ok := r.URL.Query()["limit"]
	if !ok || len(limit[0]) < 1 {
		limit = nil
	}

	var result *sql.Rows
	var err error
	db := database.GetDatabase()
	if limit != nil {
		result, err = db.Query("SELECT * FROM public.post LIMIT $1", limit[0])
	} else {
		result, err = db.Query("SELECT * FROM public.post")
	}
	if err != nil {
		panic(err)
	}
	defer result.Close()

	posts := []Post{}
	for result.Next() {
		var post Post
		err := result.Scan(&post.Title, &post.Posted, &post.Body, &post.Id)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)

}
