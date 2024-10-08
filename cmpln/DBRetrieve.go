package cmpln

import (
    "fmt"
    "cmpln/models"
)




// Function to return random posts to the client
func RetrievePosts(topic string, limitnum int) ([]models.Post, error) {
    var postsAsArray []models.Post
    

    // Add a test post to the array
    //postsAsArray = append(postsAsArray, Post{Nickname: "something", Description: "to test", Date: time.Now()})

    // Setup DB connection
    err := SetupDBConn("root", "admin", "cmplnDB")
    if err != nil {
        return nil, fmt.Errorf("Error trying to establish a DB connection: %v", err)
    }

    // Use placeholders to prevent SQL injection
    fetchQuery := "SELECT id AS ID, nickname AS Nickname, description AS Description, date AS Date, topic AS Topic FROM Post WHERE topic = ? LIMIT ?"
    rows, err := db.Query(fetchQuery, topic, limitnum)
    if err != nil {
        return nil, fmt.Errorf("Error trying to query the DB: %v", err)
    }
    defer rows.Close()
    // close db connection after fetching posts.
    defer db.Close()

    // Add the rows to the array of post structs
    for rows.Next() {
        var post models.Post
        // Scan the row into the Post struct
        if err := rows.Scan(&post.Id, &post.Nickname, &post.Description, &post.Date, &post.Topic); err != nil {
            return nil, fmt.Errorf("Error scanning row into Post struct: %v", err)
        }
        postsAsArray = append(postsAsArray, post)
    }

    // Check for errors after row iteration
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error iterating over rows: %v", err)
    }
   

    return postsAsArray, nil
}



func RetrievePost(id int64) (models.Post, error) {
    // setup db connection
    if err := SetupDBConn("root", "admin", "cmplnDB"); err != nil {
        return models.Post{}, fmt.Errorf("Error trying to establish connection to DB in RetrievePost-Function: %v", err)
    }
    


    fetchQuery := "SELECT id AS ID, nickname AS Nickname, description AS Description, date AS Date, topic AS Topic FROM Post WHERE id = ? LIMIT 1"
    row, err := db.Query(fetchQuery, id)
    if err != nil {
        return models.Post{}, fmt.Errorf("Error trying to query the DB in RetrievePost-Function:%v", err)
    }
    
    if row == nil {
        return models.Post{}, fmt.Errorf("Row doesnt contain the single post to fetch")
    }

    defer row.Close()
    defer db.Close()

    var post models.Post
    for row.Next(){
        if err := row.Scan(&post.Id ,&post.Nickname, &post.Description, &post.Date, &post.Topic); err != nil {
            return models.Post{}, fmt.Errorf("Error trying to scan for post object in RetrievePost-Function: %v", err)
        }
    }
    

    return post, nil
}

