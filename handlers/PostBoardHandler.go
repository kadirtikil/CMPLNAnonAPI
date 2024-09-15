package handlers

import (
	"cmpln/cmpln"
	"cmpln/view"
	"fmt"

	"github.com/a-h/templ"
)

func HandlePostBoard() templ.Component {
	fmt.Println("handler is active")

	posts, err := cmpln.RetrievePosts("test", 16)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// get the component now
	component := view.Post(posts)

	// now just return the rendered component with its new posts.
	// need to find an alogirthm, which will then distribute posts evenly on call of client
	// such that the posts do not repeat

	// to now send the new component back with updated elements i have to write a render function, that
	// turns it into a byte slice
	return component

	// Or just make it return a component and use templ.handler at routes.
}
