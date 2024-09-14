package main

import (
	"cmpln/cmpln"
	"cmpln/handlers"
	"cmpln/view"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	// const values
	const listeningport = ":8080"
	httpMux := http.NewServeMux()
	server := http.Server{
		Handler: httpMux,
		Addr:    listeningport,
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

	// templ<> is used for components in combination with htmx. This application is completely server side.

	/* model_post, err := cmpln.RetrievePost(50)
	if err != nil {
		fmt.Println(err)
	} */

	component := view.MainPage()
	httpMux.Handle("/", templ.Handler(component))

	httpMux.Handle("GET /postBoard", templ.Handler(handlers.HandlePostBoard()))

	httpMux.Handle("GET /modal", templ.Handler(view.PostForm()))

	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

	fmt.Printf("Listening on Port %s\n", listeningport)
	log.Fatal(server.ListenAndServe())

}
