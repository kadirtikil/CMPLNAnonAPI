package cmpln


import(
    "testing"
    "fmt"
    "reflect"
)


func TestRetrievePosts(t *testing.T) {
    var posts []Post
    posts = append(posts, Post{Nickname: "test name", Description: "test description", Date: "2024-09-01 23:51:21", Topic: "test"})
    posts = append(posts, Post{Nickname: "second name", Description: "second description", Date : "2024-09-01 23:54:59", Topic: "test"})
    

    tests := []struct{
        name string
        topic string
        limit int 
        postsArray []Post
    }{
        {
            name: "testing the fetching of some posts",
            topic: "test",
            limit: 2,
            postsArray: posts,
        },
    }
   
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T){
            actuals, err := RetrievePosts(test.topic, test.limit)

            if err != nil {
                fmt.Println(err)
                t.Errorf("Error, trying to fetch data in TestRetrievePosts-Function")
                return
            }
            
            if !reflect.DeepEqual(test.postsArray, actuals) {
                t.Errorf("actual does not equal to the expected values of postsArray in tests Array.")
            }

        })

    }


}


func TestRetrievePost(t *testing.T) {


    expectedPostObj := Post{
        Nickname: "test name",
        Description: "test description",
        Date: "2024-09-01 23:51:21",
        Topic: "test",
    }

    tests := []struct{
        name string
        id int64
        expected Post
    }{
        {
            name: "Test fetching single Post by id",
            id: 1,
            expected: expectedPostObj, 
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T){
            actual, err := RetrievePost(test.id)
            if err != nil {
                t.Errorf("Error trying to fetch Post in TestRetrievePost-Function: %v", err)
            }
            
            if !reflect.DeepEqual(actual, test.expected) {
                t.Errorf("The expected return value doesnt equal to the actual value in TestRetrievePost-Function")
            }
        })
    }
}
