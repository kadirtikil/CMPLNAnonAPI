package cmpln

import(
    "net/http"
    "strings"
    "strconv"
    "fmt"
    "encoding/json"
)

func HTTPRetrievePosts(w http.ResponseWriter, r *http.Request) {
    // get topic and amount from the url
    splitURL := strings.Split(r.URL.Path, "/")
    topic := splitURL[len(splitURL) - 2]
    limit, err := strconv.Atoi(splitURL[len(splitURL) - 1])
    if err != nil {
        fmt.Errorf("Error trying to convert string num into an int: %#v", err)
    }
    
    // ok now fetch the posts with the topic and limit.

    posts, err := RetrievePosts(topic, limit)
    
    if err != nil{
        fmt.Errorf("Error trying to fetch the posts from db in HTTPRetrievePosts-Function: %#v", err)
    }
    

    // set up the ret value to send to client
    retval, err := json.Marshal(posts)

    if err != nil {
        fmt.Errorf("Error trying to marshal the posts in HTTPRetrievePosts-Function: %#v", err)
    }
   

    w.Write(retval)
}


func HTTPRetrievePost(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL) - 1])
    if err != nil {
        http.Error(w, "no such post available", http.StatusInternalServerError)
    }
        
    post, err := RetrievePost(int64(id))
    
    if err != nil {
        http.Error(w, "no such post available", http.StatusInternalServerError)
    }

    retvalue, err := json.Marshal(post)
    if err != nil {
        http.Error(w, "format of post does not fit", http.StatusNotFound)
    }


    w.Write(retvalue)


}



