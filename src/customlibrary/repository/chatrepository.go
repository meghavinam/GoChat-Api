package repository

import (
	"fmt"
	er "prod/src/customlibrary/errorhandler"
	sr "prod/src/customlibrary/services"
	"time"
)

/*
*save chat message into database
 */
func SaveMessage(message string) int64 {
	insertData := map[string]string{}
	insertData["message"] = message
	lastId := sr.InsertPreparedQueryErr(sr.ClientDb, "chat_messages", insertData)
	return lastId
}

/*
*update chat message into database
 */
func UpdateMessage(message string, id string) int64 {
	insertData := map[string]string{}

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	insertData["updated_date"] = now.Format("2006-01-02 15:04:05")

	insertData["message"] = message
	condition := make(map[string]string)
	condition["id"] = fmt.Sprint(id)

	retId := sr.UpdatePrepareQuery(sr.ClientDb, "chat_messages", insertData, condition)
	return retId
}

/*
*delete chat message from database
 */

func DeleteChatMessage(id string) int64 {

	condition := make(map[string]string)
	condition["id"] = fmt.Sprint(id)

	retId := sr.DeletePrepareQuery(sr.ClientDb, "chat_messages", condition)
	return retId
}

/*
*list all chat message from database
 */

func GetAllChatMessage() map[string]string {

	var sqlString = fmt.Sprint("select id,message   FROM chat_messages")
	resultQuery := sr.SelectDirectQuery(sr.ClientDb, sqlString)

	customData := map[string]string{}

	for resultQuery.Next() {
		var id, message string
		err := resultQuery.Scan(&id, &message)
		er.ErrorCheck(err)
		customData[id] = message

	}
	defer resultQuery.Close()
	return customData

}
