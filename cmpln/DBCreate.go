package cmpln

import (
	"fmt"
)

func CreatePost(nickname, description, topic string) (int64, error) {

	fmt.Println("create post is active")

	// if params are empty just throw an error
	if nickname == "" || description == "" || topic == "" {
		return 0, fmt.Errorf("Params are empty. Creating a post is not possible i n CreatePost-Function.")
	}

	if err := SetupDBConn("root", "admin", "cmplnDB"); err != nil {
		return 0, fmt.Errorf("Error trying to establich DB conn in CreatePost-Function: %v", err)
	}

	query := "INSERT INTO Post (nickname, description, date, topic) values(?, ?, NOW(), ?)"

	// create Post first
	retvalue, err := db.Exec(query, nickname, description, topic)

	if err != nil {
		return 0, fmt.Errorf("Error trying to create new Post in CreatePost-Function: %v", err)
	}

	// fetch post to check if it really exists. an extra measure to the error check
	// thought its not possible, but the retvalue from exec can give back thelast posts id. which makes it possible
	id, err := retvalue.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("Error trying to fetch last id in CreatePost-Function: %v", err)
	}

	defer db.Close()

	return id, nil
}
