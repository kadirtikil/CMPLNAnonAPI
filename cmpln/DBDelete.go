package cmpln

import (
    "fmt"
    "reflect"
    "cmpln/models"
)


func DeletePost(id int64) (bool, error) {
    

    // check if the post with this id exists
    post, err := RetrievePost(id)
    if err != nil {
        return false, fmt.Errorf("Error trying to fetch not existing Post in DeletePost-Function: %v", err)
    }
   

    if reflect.DeepEqual(post, models.Post{}) {
        return false, fmt.Errorf("Error, empty set returned. The Post does not exist: %v", err)
    }

    // setup db connection
    if err := SetupDBConn("root", "admin", "cmplnDB"); err != nil {
        return false, fmt.Errorf("Error trying to establish DB Connection in DeletePost-Function: %v", err)
    }

    query:= "DELETE FROM Post where id = ?" 
    if _, err := db.Exec(query, id); err != nil {
        return false, fmt.Errorf("Error trying to delete a post in DeletePost function: %v",err)
    }
    
    defer db.Close()

    return true, nil
}
