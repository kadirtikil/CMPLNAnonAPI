package cmpln

import (
    "net/http"
    "io"
    "encoding/json"
)


func HTTPCreatePost(w http.ResponseWriter, r *http.Request) {
    
    // read the body of the request 
    requestbody, err := io.ReadAll(r.Body)
    
    if err != nil {
        http.Error(w, "reading request body didnt work", http.StatusInternalServerError)
        return
    }
  

    // make a post obj from the data in the request body.
    
    var post Post
    if err := json.Unmarshal(requestbody, &post); err != nil {
        http.Error(w, "Error trying to unmarshal request body" , http.StatusInternalServerError)
        return
    }
    
    // object is arriving here

    id, err:= CreatePost(post.Nickname, post.Description, post.Topic)
    if err != nil {
        http.Error(w, "Couldnt create new object in HTTPCreatePost-Function", http.StatusInternalServerError)
        return
    }
    

    // check if post really exists now
    if _, err:= RetrievePost(id); err != nil {
        http.Error(w, "Post could not be found in HTTPCreatePost-Function", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("New Post created!"))
}
