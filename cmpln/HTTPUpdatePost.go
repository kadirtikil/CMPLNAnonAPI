package cmpln

import(
    "net/http"
    "encoding/json"
)

func HTTPUpdatePost(w http.ResponseWriter, r *http.Request) {
    
    var post Post

    err := json.NewDecoder(r.Body).Decode(&post)
    if err != nil {
        http.Error(w, "Error trying to read the request body", http.StatusInternalServerError)
        return
    }

    defer r.Body.Close()

    if len(post.Nickname) > 49 || len(post.Description) > 399 || len(post.Topic) > 49 {
        http.Error(w, "Format of attributes is wrong", http.StatusInternalServerError)
        return
    }
    
    if err := UpdatePost(post.Id, post.Nickname, post.Description, post.Topic); err != nil {
        http.Error(w, "couldnt update post in HTTPUpdatePost-Function", http.StatusInternalServerError)
        return
    }
    

    w.Write([]byte("Post updated!"))

}
