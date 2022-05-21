package database

import (
	"fmt"
	"github.com/test_server/internal/config"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

func OpenConnection() db.Session {
	sess, err := postgresql.Open(config.GetDatabaseSetting())
	if err != nil {
		fmt.Println("Open: ", err)
	}

	return sess
}

func CloseConnection(sess db.Session) {
	err := sess.Close()
	if err != nil {
		fmt.Println("Close Database error: ", err)
	}
}
