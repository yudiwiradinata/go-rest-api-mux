package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/yudiwiradinata/go-rest-api-mux/entity"
	"github.com/yudiwiradinata/go-rest-api-mux/repository"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
)

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	posts, err := postRepository.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error get the posts"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshaling the request"}`))
		return
	}

	post.ID = rand.Int63()
	_, err = postRepository.Save(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error save the posts"}`))
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}
