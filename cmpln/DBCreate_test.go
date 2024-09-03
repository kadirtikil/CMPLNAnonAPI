package cmpln


import(
    "testing"
    "reflect"
)

func TestCreatePost(t *testing.T) {
    t.Run("create an object, fetch it, delete it", func(t *testing.T) {
        id, err := CreatePost("test create", "test create", "test")
        if err != nil {
            t.Errorf("Error trying to create a new objectin TestCreatePost-Function: %#v", err)
        }           

        // fetch to check again
        post, err := RetrievePost(id)
        if err != nil {
            t.Errorf("Error trying to fetch newly created post object in TestCreatePost-Function: %#v", err)
        }
        
        if post.Nickname != "test create" {
            t.Errorf("Error, the fetched object is not the initially created object: %#v", err)
        }
       
        // delete the object
        if _, err := DeletePost(id); err != nil {
            t.Errorf("Error trying to delete the created object in TestCreateObject-Function: %#v", err)
        }
        
        
        // Try to fetch the post again. the post should not exist at this point
        initialpost, _:= RetrievePost(id)

        // double check if the post is a default one (empty one)
        if !reflect.DeepEqual(initialpost, Post{}) {
            t.Errorf("Error, the fetched post after deletion should be an empty one in TestCreatPost-Function: %#v", err)
        }
    })


    t.Run("create an object with missing parameters.", func(t *testing.T) {
        if _, err := CreatePost("", "", "test"); err == nil {
            t.Errorf("Error, params missing: %#v", err)
        }
    })

}
