package cmpln

import( 
    "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
)

var db *sql.DB

func SetupDBConn(user, password, dbname string) (error) {
    if user == "" || password == "" || dbname == "" {
        return fmt.Errorf("user, password, and dbname must not be empty")
    }
    
    dbconfig := mysql.Config{
        User:   user,
        Passwd: password,
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: dbname,
        AllowNativePasswords: true,
    }
   
    // use db gloablly, therefore need to be set by using =. so a err variable has to be used in order to use = 
    var err error
    db, err = sql.Open("mysql", dbconfig.FormatDSN())
    if err != nil {
        return fmt.Errorf("connection in sql.Open could not be established: %v", err)
    }

    if err := db.Ping(); err != nil {
        return fmt.Errorf("DB cannot be reached. Creds are probably wrong: %v", err)
    }

    return nil
}


