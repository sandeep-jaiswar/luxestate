package db

import (
    "database/sql"
    "sync"

    _ "github.com/go-sql-driver/mysql"
)

var (
    dbInstance *sql.DB
    once       sync.Once
)

func GetDBInstance() *sql.DB {
    once.Do(func() {
        var err error
        dbInstance, err = sql.Open("mysql", "root:password@tcp(0.0.0.0:3306)/development")
        if err != nil {
            panic(err)
        }
    })
    return dbInstance
}
