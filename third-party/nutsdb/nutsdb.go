package nutsdb

import (
	"log"

	"github.com/nutsdb/nutsdb"
)

var db *nutsdb.DB

func Initial() {
	// Open the database located in the /tmp/nutsdb directory.
	// It will be created if it doesn't exist.
	db, err := nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir("/tmp/nutsdb"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func GetInstance() *nutsdb.DB {
	return db
}
