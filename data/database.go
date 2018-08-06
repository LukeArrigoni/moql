package data

import "database/sql"

func GetDataConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:invision@/invisionapp")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}