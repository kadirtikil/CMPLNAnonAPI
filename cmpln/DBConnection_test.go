package cmpln

import (
    "testing"
)

func TestSetupDBConnection(t *testing.T) {
    test := struct{
        nameTest string
        nameDB string
        userDB string
        passwordDB string
        expectedMsg string
    }{
        nameTest: "DB Conncetion worked!",
        nameDB: "cmplnDB",
        userDB: "root",
        passwordDB: "admin",
        expectedMsg: "Not a nil value",
    }
    
    t.Run(test.nameTest, func (t *testing.T){
        err := SetupDBConn(test.userDB, test.passwordDB, test. nameDB)
        if err != nil {
            t.Error("Error occured while trying to establish the DB-Connection")
            return
        }

    })

}
