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
        http.Error(w, fmt.Sprintf("Error trying to convert string num into an int: %#v", err), http.StatusInternalServerError)
        return
    }
    
    // ok now fetch the posts with the topic and limit.

    posts, err := RetrievePosts(topic, limit)
    
    if err != nil{
        http.Error(w, fmt.Sprintf("Error trying to fetch the posts from db in HTTPRetrievePosts-Function: %#v",err ), http.StatusInternalServerError)
        return
    }



    // set up the ret value to send to client
    retval, err := json.Marshal(posts)

    if err != nil {
        http.Error(w, fmt.Sprintf("Error trying to marshal the posts in HTTPRetrievePosts-Function: %#v", err), http.StatusInternalServerError)
        return
    }
    

    w.Write(retval)
}


func HTTPRetrievePost(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(splitURL[len(splitURL) - 1])
    if err != nil {
        http.Error(w, fmt.Sprintf("no such post available see: %#v", err), http.StatusInternalServerError)
        return
    }
        
    post, err := RetrievePost(int64(id))
    
    if err != nil {
        http.Error(w, fmt.Sprintf("no such post available see: %#v", err), http.StatusInternalServerError)
        return
    }

    retvalue, err := json.Marshal(post)
    if err != nil {
        http.Error(w, fmt.Sprintf("format of post does not fit see: %#v", err), http.StatusNotFound)
        return
    }


    w.Write(retvalue)


}



