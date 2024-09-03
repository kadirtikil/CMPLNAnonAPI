package cmpln


import(
    "testing"
)


// testing deletion is a little special. 
// first create the object to delete with CreatePost Function
// then fetch it with Retrieve Post
// then delete it and try to fetch again.
// if fetched then its not deleted. if not then it is. 

// furhter testing cases involve cases like, a post not existing
// or trying to delete a post with a different nickname. This case is a little more work tho, since
// it can be solved with auth as well. but checking it anyway should be good practice


func TestDeletePost(t *testing.T) {
    tests:= []struct{
        name string
        
    }{
        {
            name: "create and delete object",
        },
        {
            name: "delete not existing object",
        },
    }
    

    // these tests dont have the same structure so they cant be looped and therefore have to be tested seperately
    t.Run(tests[0].name, func(t *testing.T){
        // create the post first
        idnum, err := CreatePost("post to delete", "post to delete", "delete") 
        if err != nil {
            t.Errorf("Error trying to create Objecjt in first testcase of TestDeletePost-Function: %#v", err)
        }

        // fetch the post to make sure its created
        post, err := RetrievePost(idnum)
        if err != nil {
            t.Errorf("Error trying to fetch the Post that has just been created in first testcase of TestDeletePost-Function: %#v", err)
        }
       
        // now check if post is ok
        if post.Nickname != "post to delete" {
            t.Errorf("Error, the object that has just been fetched is not the same to initially created object in first testcase of TestDeletePost-Function: %#v", err)
        }

        // now delete the object
        if _, err := DeletePost(idnum); err != nil {
            t.Errorf("Error trying to delete the post in first testcase of TestDeletePost-Function: %#v", err)
        }

        // now try fetching it again
        if fetchingdeletedpost, err := RetrievePost(idnum); fetchingdeletedpost.Nickname == "post to delete" {
            t.Errorf("Error trying to fetch the Post that has just been created in first testcase of TestDeletePost-Function: %#v", err)
        }


    })
    
    t.Run(tests[1].name, func(t *testing.T){
        // id is 6 because 6 once existed. the auto increment passed it tho, so it cannot be created again. this needs changing tho as soon as a new db is used
        // cause the 6 could be the id of a valid post in the db.
        if status, _:= DeletePost(6); status == true {
            t.Error("Error at second testcase in TestDeletePost-Function. This post should not have existed")
        }
        

    })
   
}
