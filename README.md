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
![Alt text](.readmestuff/DBD.png)
(DBConnection.go)

## Create
Ist eine simple Funktion, die ein POST Request behandelt.


## Retrieve
Hierzu gibt es zwei Funktionen. Die eine fetched einen bestimmten Post. 
![Alt text](.readmestuff/DBD.png)

Die andere fetched eine Bestimmte Anzahl an Posts. ACHTUNG: Diese Funktion fetched Posts unter einem gewissen topic mit einem Limit (Anzahl an Posts). Um zu gewährleisten, dass alle Posts gleich oft (also linear oft) gefetched werden, muss noch ein Algorithmus hierzu entwickelt werden (folgt).
![Alt text](.readmestuff/DBD.png)

## Update
Ist eine simple Funktion, die ein PUT Request behandelt.
![Alt text](.readmestuff/DBD.png)

## Delete
Ist eine simple Funktion, die eine DELETE Request behandelt.
![Alt text](.readmestuff/DBD.png)
