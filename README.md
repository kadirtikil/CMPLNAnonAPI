# CMPLN

CMPLN steht für das englische Wort **complain**, ohne die Selbstlaute (ja ich weiss sehr kreativ). 
Die Anwendung kann Beschwerden von allen Nutzern hier anzeigen. Eine bestimmte Authentifizierung ist nicht gegeben. Es werden nur Nicknames auf Email-Adressen aufgelöst, um eventuelle Beiträge zurück zu verfolgen. Der Post muss dann per Email bestätigt werden. <br>


## Test Driven Development
Alle Komponenten der Anwendung werden Testbegleitend implementiert. Ein Workflow wird diese später ausführen und nur bei 100% Erfolg der Testfälle diesen Deployen (Sei es in einem Docker Container auf irgendeiner cloud (oder auch nicht weil es kostet Geld) oder auf einem Server per FTP oder SFTP, oder auf meiner IP per Portforwarding per FTP oder SFTP (sehr ungerne)).<br>


## Client Side
Das Frontend wird in HTMX mit der Nutzung von TEMPL zur Komponentisierung (dieses Wort existiert eventuell nicht) gerendered. HTMX holt sich per GET-Request die Komponenten aus dem Server, und rendered diese durch Adressierung eines Objekts in der DOM per **hx-target** oder **hx-swap** (je nach Use-Case).
(siehe [HTMX Docs](https://htmx.org/docs/#targets))



## Datenbank
MariaDB. Keine Ahnung warum unbedingt. Ist halt eine Relationale Datenbank die gut funktioniert und kostenlos ist (zumindest lokal). Werde wahrscheinlich aber auf PostgreSQL wechseln.
![Alt text](.readmestuff/DBD.png)


# Struktur
Aufbau des gesamten Projekts hier erklärt.

## cmpln
Enthält:
- Datenbanklogik
- HTTPlogik
- Tests der Datenbanklogik

## view
Enthält:
- TEMPL Komponenten mit HTMX features.
- (Go Dateien dort sind Produkt aus **templ generate** command)

## handlers
Enthält:
- Handler für das fetchen der TEMPL Komponenten.


## Weiteres
Makefile zur automatisierung gängiger commands. Wie das pushen auf git, oder das builden des Servers (jede Änderung in TEMPL erfordert ein weiteres **templ generate**, nervt auf dauer). Weiter noch Go Module Metadata und Ordner für Readmezeugs.

# CRUD Operationen
Um meine go Kenntnisse zu erweitern und aufzufrischen, habe ich CRUD Operationen implementiert. Diese sind keine API. Alles wird Serverseitig gerendered. Da HTMX in kombination mit TEMPL dafür sorgt, dass die ganze DOM nicht neu geladen werden muss, bei jeder Operation, ist die Belastung nicht zu hoch. Zumal es an sich eine leichtgewichtigte Anwendung ist.

## Datenbankverbindung
Globale Variable die Verbindung der Datenbank enthält. Jede Datenbankoperation öffnet eine Verbindung, und schließt diese danach. D.h. die Verbindung bleibt nicht offen außerhalb einer Operation.

Zutrittsdaten sind öffentlich da es sich um den localhost handelt. Sonst in einer .env die von git ignoriert wird. 
```go
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
```

(DBConnection.go)

## Create
Ist eine simple Funktion, die ein POST Request behandelt.
```go
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
```

## Retrieve
Hierzu gibt es zwei Funktionen. Die eine fetched einen bestimmten Post. 
```go
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
```

Die andere fetched eine Bestimmte Anzahl an Posts. ACHTUNG: Diese Funktion fetched Posts unter einem gewissen topic mit einem Limit (Anzahl an Posts). Um zu gewährleisten, dass alle Posts gleich oft (also linear oft) gefetched werden, muss noch ein Algorithmus hierzu entwickelt werden (folgt).
```go
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

```
## Update
Ist eine simple Funktion, die ein PUT Request behandelt.
```go
func UpdatePost(idnum int64, n, desc, topic string) (error) {
    if err :=  SetupDBConn("root", "admin", "cmplnDB"); err != nil {
        return fmt.Errorf("Error trying to establish DB connection in UpdatePost-Function: %v", err)
    }

    query := "UPDATE Post SET nickname = ?, description = ?, date = NOW(),topic = ? WHERE id = ?"
    
    if _,err := db.Exec(query, n, desc, topic, idnum); err != nil {
        return fmt.Errorf("Error trying to update the Post in UpdatePost-Function: %v", err) 
    }


    return nil
}
```


## Delete
Ist eine simple Funktion, die eine DELETE Request behandelt.
```go
func HTTPDeletePost(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    
    id, err := strconv.ParseInt( splitURL[len(splitURL) - 1], 10, 0) 
    
    if err != nil {
        http.Error(w, "Id not in URL", http.StatusInternalServerError)
        return        
    }

    if _, err := DeletePost(id); err != nil {
        http.Error(w, "Deletion not possible", http.StatusInternalServerError)
        return
    }

    // apparently used if DELETE method is used. dont really know why but seems to be best practise.
    w.WriteHeader(http.StatusNoContent)
}
```
