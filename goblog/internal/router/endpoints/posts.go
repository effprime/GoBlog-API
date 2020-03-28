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

	if utils.VerifyMethod(r, "POST") {
		err := utils.GetPostRequestData(r, &post)
		if err != nil {
			code = http.StatusUnprocessableEntity
			R.Status = "failure"
			R.Message = "unable to create post object"
		} else {
			err = database.NewPost(post)
			if err != nil {
				code = http.StatusInternalServerError
				R.Status = "failure"
				R.Message = "unable to save post object"
			} else {
				code = http.StatusCreated
				R.Status = "success"
				R.Message = "new post created"
			}
		}
	} else {
		code = 405
		R.Status = "failure"
		R.Message = "method not allowed"
	}
	json.NewEncoder(w).Encode(R)
	w.WriteHeader(code)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	args, ok := r.URL.Query()["id"]
	var id int

	var R models.HttpResponse
	var code int

	if utils.VerifyMethod(r, "POST") {
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
	} else {
		code = 405
		R.Status = "failure"
		R.Message = "method not allowed"
	}
	json.NewEncoder(w).Encode(R)
	w.WriteHeader(code)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	args, ok := r.URL.Query()["id"]
	var id int

	var R models.HttpResponse
	var code int

	if utils.VerifyMethod(r, "GET") {
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
						R.Message = "unable to create object"
					} else {
						code = http.StatusInternalServerError
						R.Status = "success"
						R.Message = "post retrieved"
						R.Payload = json.RawMessage(post)
					}
				}
			}
		}
	} else {
		code = 405
		R.Status = "failure"
		R.Message = "method not allowed"
	}
	json.NewEncoder(w).Encode(R)
	w.WriteHeader(code)
}

func AllPostsHandler(w http.ResponseWriter, r *http.Request) {
	args, ok := r.URL.Query()["limit"]
	var limit int

	var R models.HttpResponse
	var code int

	if utils.VerifyMethod(r, "GET") {
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
			R.Message = "unable to get posts"
		} else {
			code = http.StatusOK
			R.Status = "success"
			R.Message = "posts retrieved"
			R.Payload = posts
		}
	} else {
		code = 405
		R.Status = "failure"
		R.Message = "method not allowed"
	}
	json.NewEncoder(w).Encode(R)
	w.WriteHeader(code)
}
