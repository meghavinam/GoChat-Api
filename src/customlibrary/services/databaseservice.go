package services

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	cg "prod/src/customlibrary/configuration"
	er "prod/src/customlibrary/errorhandler"
	"strings"
	"time"
)

var ClientDb *sql.DB
var SendingDb *sql.DB
var ClientDbSlave *sql.DB

func SetDbConnection(dbType string, maxOpenConnections int, maxIdleConnections int) {
	switch dbType {
	case "ClientDatabase":
		var dbDetails = []string{cg.Config.ClientDatabase.Host, cg.Config.ClientDatabase.Username, cg.Config.ClientDatabase.Password, cg.Config.ClientDatabase.Database, cg.Config.ClientDatabase.Port}
		var connPool = []int{maxOpenConnections, maxIdleConnections}
		ClientDb = setDatabase(dbDetails, connPool)
	default:
		fmt.Printf("Invalid database type")
	}
}

func setDatabase(dbDetails []string, connPool []int) *sql.DB {
	var err error
	if dbDetails[0] == "" || dbDetails[1] == "" || dbDetails[2] == "" || dbDetails[3] == "" || dbDetails[4] == "" || connPool[0] == 0 || connPool[1] == 0 {
		fmt.Println("Client database configuration not set")
		os.Exit(1)
	}
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbDetails[1], dbDetails[2], dbDetails[0], dbDetails[4], dbDetails[3])
	db, err := sql.Open("mysql", connString)
	er.ErrorCheck(err)
	db.SetMaxOpenConns(connPool[0])
	db.SetMaxIdleConns(connPool[1])
	db.SetConnMaxLifetime(time.Hour)
	PingDB(db)
	return db
	//defer ClientDb.Close()
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	er.ErrorCheck(err)
}

func DeletePrepareQuery(db *sql.DB, tableName string, condition map[string]string) int64 {
	var affectedRows int64 = 0
	var e error
	values := make([]interface{}, 0, 0)
	var sqlString = "delete  from " + tableName
	if len(condition) > 0 {
		sqlString += " where "
		var i = 0
		for key, val := range condition {
			if i == 0 {
				sqlString += key + " = ? "
			} else {
				sqlString += " and " + key + " = ? "
			}
			values = append(values, val)
			i++
		}
	}
	stmt, e := db.Prepare(sqlString)
	er.ErrorCheck(e)
	defer stmt.Close()
	res, e := stmt.Exec(values...)
	er.ErrorCheck(e)
	affectedRows, e = res.RowsAffected()
	er.ErrorCheck(e)
	return affectedRows
}

func UpdatePrepareQuery(db *sql.DB, tableName string, fieldValues map[string]string, condition map[string]string) int64 {
	var affectedRows int64 = 0
	var e error
	values := make([]interface{}, 0, 0)
	var sqlString = "update " + tableName + " set "
	if len(fieldValues) > 0 {
		for fk, fv := range fieldValues {
			sqlString += fk + " = ? ,"
			values = append(values, fv)
		}

	}
	sqlString = strings.TrimRight(sqlString, ",")
	if len(condition) > 0 {
		sqlString += " where "
		var i = 0
		for key, val := range condition {
			if i == 0 {
				sqlString += key + " = ? "
			} else {
				sqlString += " and " + key + " = ? "
			}
			i++
			values = append(values, val)
		}
	}
	stmt, e := db.Prepare(sqlString)
	er.ErrorCheck(e)
	defer stmt.Close()
	res, e := stmt.Exec(values...)
	er.ErrorCheck(e)
	affectedRows, e = res.RowsAffected()
	er.ErrorCheck(e)
	return affectedRows
}

func InsertPreparedQueryErr(db *sql.DB, tableName string, fieldValues map[string]string) int64 {
	var id int64 = 0
	var e error
	var fieldString = ""
	var valueString = ""
	values := make([]interface{}, 0, 0)
	if len(fieldValues) > 0 {
		for fk, fv := range fieldValues {
			fieldString += fk + ","
			valueString += "?,"
			values = append(values, fv)
		}
		fieldString = strings.TrimRight(fieldString, ",")
		valueString = strings.TrimRight(valueString, ",")
	}

	var sqlString = "insert into " + tableName + " (" + fieldString + ") values (" + valueString + ") "
	fmt.Println(sqlString)
	stmt, e := db.Prepare(sqlString)
	er.ErrorCheck(e)
	defer stmt.Close()
	res, e := stmt.Exec(values...)
	if e != nil {
		return 0
	}
	er.ErrorCheck(e)
	id, e = res.LastInsertId()
	er.ErrorCheck(e)
	return id

}
func SelectDirectQuery(db *sql.DB, sqlString string) *sql.Rows {
	var rows *sql.Rows
	rows, e := db.Query(sqlString)
	er.ErrorCheck(e)
	return rows
}
