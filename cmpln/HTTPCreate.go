package cmpln

import (
	"cmpln/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HTTPCreatePost(w http.ResponseWriter, r *http.Request) {

	// read the body of the request
	requestbody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("The following error occured in HTTPCreatePost 1: %v\n", err)
		http.Error(w, "reading request body didnt work", http.StatusInternalServerError)
		return
	}

	// make a post obj from the data in the request body.

	fmt.Println(string(requestbody))

	var post models.Post
	if err := json.Unmarshal(requestbody, &post); err != nil {
		fmt.Printf("The following error occured in HTTPCreatePost 2: %v\n", err)
		http.Error(w, "Error trying to unmarshal request body", http.StatusInternalServerError)
		return
	}

	// object is arriving here

	id, err := CreatePost(post.Nickname, post.Description, post.Topic)
	if err != nil {
		fmt.Printf("The following error occured in HTTPCreatePost 3: %v\n", err)
		http.Error(w, "Couldnt create new object in HTTPCreatePost-Function", http.StatusInternalServerError)
		return
	}

	// check if post really exists now
	if _, err := RetrievePost(id); err != nil {
		fmt.Printf("The following error occured in HTTPCreatePost 4: %v\n", err)
		http.Error(w, "Post could not be found in HTTPCreatePost-Function", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("New Post created!"))
}
