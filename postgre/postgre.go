package postgreutil

import(
    "database/sql"
    _ "github.com/lib/pq"
    "encoding/json"
    "os"
)

type postgreConf struct {
	Host   string
	User   string
	Passw  string
	Dbname string
    MaxConn int
}

//  Connect to a postgres Postgre
func ConnectDB(host string, user string, passw string, dbname string, maxConn int) (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", "host="+host+" user="+user+" password="+passw+" dbname="+dbname+" sslmode=disable")
	if nil != err {
		return nil, err
	}
	err = db.Ping()
	if nil != err {
		return nil, err
	}
	db.SetMaxOpenConns(maxConn)

	return db, nil
}

func ConnectDBFromFile(file string) (*sql.DB, error) {
	fh, err := os.Open(file)
	if nil != err {
		return nil, err
	}

	dec := json.NewDecoder(fh)
	var dbconf postgreConf
	err = dec.Decode(&dbconf)
	if nil != err {
		return nil, err
	}
	return ConnectDB(dbconf.Host, dbconf.User, dbconf.Passw, dbconf.Dbname, dbconf.MaxConn)
}

func CreateConfFile(file string) error {
	dbconf := postgreConf{"localhost", "dbuser", "dbpass", "yourdatabase", 10}
	fh, err := os.Create(file)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(fh)
	err = enc.Encode(&dbconf)
	return err
}
