package endpoints

import (
	"encoding/json"
	"goblog/internal/database"
	"goblog/internal/models"
	"goblog/internal/utils"
	"net/http"
	"strconv"
)

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	var R models.HttpResponse
	var code int

	err := utils.GetPostRequestData(r, &post)
	if err != nil {
		code = http.StatusUnprocessableEntity
		R.Status = "failure"
		R.Message = err.Error()
	} else {
		err = database.NewPost(post)
		if err != nil {
			code = http.StatusInternalServerError
			R.Status = "failure"
			R.Message = err.Error()
		} else {
			code = http.StatusCreated
			R.Status = "success"
			R.Message = "new post created"
		}
	}
	utils.MakeResponse(w, code, R)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	args, ok := r.URL.Query()["id"]
	var id int

	var R models.HttpResponse
	var code int

	if !ok || len(args[0]) < 1 {
		code = http.StatusNotAcceptable
		R.Status = "failure"
		R.Message = "no ID value provided"
	} else {
		var err error
		id, err = strconv.Atoi(args[0])
		if err != nil {
			code = http.StatusUnprocessableEntity
			R.Status = "failure"
			R.Message = err.Error()
		} else {
			err = database.DeletePost(id)
			if err != nil {
				code = http.StatusInternalServerError
				R.Status = "failure"
				R.Message = err.Error()
			} else {
				code = http.StatusOK
				R.Status = "success"
				R.Message = "post deleted"
			}
		}
	}
	utils.MakeResponse(w, code, R)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	args, ok := r.URL.Query()["id"]
	var id int

	var R models.HttpResponse
	var code int

	if !ok || len(args[0]) < 1 {
		code = http.StatusNotAcceptable
		R.Status = "failure"
		R.Message = "no ID value provided"
	} else {
		var err error
		id, err = strconv.Atoi(args[0])
		if err != nil {
			code = http.StatusUnprocessableEntity
			R.Status = "failure"
			R.Message = err.Error()
		} else {
			post, err := database.GetPost(id)
			if err != nil {
				code = http.StatusInternalServerError
				R.Status = "failure"
				R.Message = err.Error()
			} else if post.Id == 0 {
				code = http.StatusNotFound
				R.Status = "failure"
				R.Message = "post not found"
			} else {
				post, err := json.Marshal(post)
				if err != nil {
					code = http.StatusInternalServerError
					R.Status = "failure"
					R.Message = err.Error()
				} else {
					code = http.StatusInternalServerError
					R.Status = "success"
					R.Message = "post retrieved"
					R.Payload = json.RawMessage(post)
				}
			}
		}
	}
	utils.MakeResponse(w, code, R)
}

func AllPostsHandler(w http.ResponseWriter, r *http.Request) {
	args, ok := r.URL.Query()["limit"]
	var limit int

	var R models.HttpResponse
	var code int

	if !ok || len(args[0]) < 1 {
		limit = 0
	} else {
		var err error
		limit, err = strconv.Atoi(args[0])
		if err != nil {
			limit = 0
		}
	}
	_, posts, err := database.GetPosts(limit)
	if err != nil {
		code = http.StatusInternalServerError
		R.Status = "failure"
		R.Message = err.Error()
	} else {
		code = http.StatusOK
		R.Status = "success"
		R.Message = "posts retrieved"
		R.Payload = posts
	}
	utils.MakeResponse(w, code, R)
}
