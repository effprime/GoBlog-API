package endpoints

import (
	"encoding/json"
	"goblog/internal/database"
	"goblog/internal/models"
	"goblog/internal/utils"
	"net/http"
	"strconv"
	"time"
)

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	utils.GetPostRequestData(r, &post)
	post.Posted = time.Now().String()

	database.NewPost(post)

	w.WriteHeader(http.StatusOK)
}

func AllPostsHandler(w http.ResponseWriter, r *http.Request) {
	args, ok := r.URL.Query()["limit"]
	var limit int

	if !ok || len(args[0]) < 1 {
		limit = 0
	} else {
		var err error
		limit, err = strconv.Atoi(args[0])
		if err != nil {
			limit = 0
		}
	}

	json.NewEncoder(w).Encode(database.GetPosts(limit))
}
