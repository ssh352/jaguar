package db

import (
	"database/sql"
	"fmt"
	// only use the init function
	_ "github.com/go-sql-driver/mysql"
)

// MysqlConfig the config params for connect to mysql
type MysqlConfig struct {
	MysqlUsernName string
	MysqlPwd       string
	MysqlURL       string
}

// MysqlWorker run sql from sqlQueue
type MysqlWorker struct {
	SQLs chan string
	*MysqlConfig
	DB *sql.DB
}

// Init read configure and connect to mysql
func (worker *MysqlWorker) Init() error {
	return worker.connect()
}

func (worker *MysqlWorker) connect() error {
	var err error
	worker.DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s", worker.MysqlUsernName, worker.MysqlPwd, worker.MysqlURL))
	return err
}

// Release close db and log
func (worker *MysqlWorker) Release() {
	worker.DB.Close()
}

// Run excute sql from sqlQueue
func (worker *MysqlWorker) Run() {
	for {
		sqlstr := <-worker.SQLs
		worker.DB.Exec(sqlstr)
	}
}
