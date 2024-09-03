package cmpln

import (
    "fmt"
)


func DeletePost(id uint) (bool, error) {
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
