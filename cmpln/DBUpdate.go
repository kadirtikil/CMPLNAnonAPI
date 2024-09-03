package cmpln

import(
    "fmt"
)


func UpdatePost(idnum uint, n, desc, topic string) (error) {
    if err :=  SetupDBConn("root", "admin", "cmplnDB"); err != nil {
        return fmt.Errorf("Error trying to establish DB connection in UpdatePost-Function: %v", err)
    }

    query := "UPDATE Post SET nickname = ?, description = ?, date = NOW(),topic = ? WHERE id = ?"
    
    if _,err := db.Exec(query, n, desc, topic, idnum); err != nil {
        return fmt.Errorf("Error trying to update the Post in UpdatePost-Function: %v", err) 
    }


    return nil
}
