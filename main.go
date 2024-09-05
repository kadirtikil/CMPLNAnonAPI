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
 

    // The following is the CRUD
    // All this will be packed in some middleware that checks a jwt or maybe not even that. The application is light weight anyway. 
    // Maybe just leaving your email with a nickname will suffice. 
    // so ill just check, if the email and the nickname match and then let the post through. 
    // checks for profanity, spam and so on will folow
    // |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
    // Setup db connection
    err := cmpln.SetupDBConn("root", "admin", "cmplnDB") 
    
    if err != nil {
        fmt.Println("whoopsie at line 42")
    }

    // Might start some go routines internally. here not needet, http takes care of the queue of requests.
    // internally there might be some cases where go routines can amplify the speed of the API. will search for a way to benchmark it.

    // Routes 
    // create post
    httpMux.HandleFunc("POST /create", cmpln.HTTPCreatePost) 

    // retrieve posts randomly
    httpMux.HandleFunc("GET /posts/{topic}/{limit}", cmpln.HTTPRetrievePosts)
    
    // retrieve certain post
    httpMux.HandleFunc("GET /post/{id}", cmpln.HTTPRetrievePost)

    // update post
    httpMux.HandleFunc("PUT /update", cmpln.HTTPUpdatePost)

    // delete post
    httpMux.HandleFunc("DELETE /delete/{id}", cmpln.HTTPDeletePost) 



    // |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||





    // The following is for serving the html
    // |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
    
    // Set the directory for file serving. index.html is / by default

    fs := http.FileServer(http.Dir("./HTMX/static"))

    httpMux.Handle("/", fs) 
 

    // |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

    fmt.Printf("Listening on Port %s", listeningport)
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
