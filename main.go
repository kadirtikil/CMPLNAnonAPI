package main


import(
    "fmt"
    "log"
    "net/http"
    "cmpln/cmpln"
)

func main() {
    // const values
    const listeningport = ":8080"
    httpMux := http.NewServeMux()    
    server := http.Server{
        Handler: httpMux,
        Addr: listeningport,
    }
  
    // Setup db connection
    err := cmpln.SetupDBConn("root", "admin", "cmplnDB") 
    
    if err != nil {
        fmt.Println("whoopsie at line 42")
    }

    fmt.Println("reached routes in main.go")
    // Routes 
    // create post
    httpMux.HandleFunc("POST /create", func(w http.ResponseWriter, r *http.Request) {w.Write([]byte("create route works"))}) 

    // retrieve posts randomly
    httpMux.HandleFunc("GET /posts", func(w http.ResponseWriter, r *http.Request) {w.Write([]byte("posts route works"))})
        
    // retrieve certain post
    httpMux.HandleFunc("GET /post/id", func(w http.ResponseWriter, r *http.Request) {w.Write([]byte("fetching one post works"))})

    // update post
    httpMux.HandleFunc("PUT /update/id", func(w http.ResponseWriter, r *http.Request) {w.Write([]byte("update route works"))})

    // delete post
    httpMux.HandleFunc("DELETE /delete/id", func(w http.ResponseWriter, r *http.Request) {w.Write([]byte("delete route works"))}) 


    fmt.Println("Listening on Port: %s", listeningport)
    log.Fatal(server.ListenAndServe())

        


   
    // fetching data limited by num works now
    /*dataToFetch, err := cmpln.RetrievePosts("test", 2)
        
    if err != nil {
        fmt.Println(err)
    }*/


    // Deleting Posts works. return true if went through and false with error if not.
    //cmpln.DeletePost(4)


    // Creating Posts works. 
    //fmt.Println(cmpln.CreatePost("new nick", "new description", "new topic"))


    // updating posts works.
    //fmt.Println(cmpln.UpdatePost(5, "main.go", "something form main.go", "golang is cool")) 


}
