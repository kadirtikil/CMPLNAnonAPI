package cmpln


import(
    "testing"
)


func TestUpdatePost(t *testing.T) {
    t.Run("check if newly created object has been updated", func (t *testing.T) {
        id, err := CreatePost("change this", "change desc", "change")

        if err != nil {
            t.Errorf("Error trying to create new Post in TestUpdatePost-Function: %#v", err)
        }
        
        if err := UpdatePost(id, "changed nickname", "changed description", "changed topic"); err != nil {
            t.Errorf("Error trying to Update post in TestUpdatePost-Function: %#v", err)
        }
        
        if _ , err := DeletePost(id); err != nil {
            t.Errorf("Error trying to delete updated Post in TestUpdatePost-Function: %#v", err)
        }

   }) 
}
