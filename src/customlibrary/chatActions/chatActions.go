/**
 * For handling chat apis
 *
 */
package chatActions

import (
	"fmt"
	rp "prod/src/customlibrary/repository"
)

/**
 * For saving the chat messages in to database
 * @input post message from api
 * @return status ,message
 */

func SaveChatBoatMessage(requestData map[string]string) (bool, string) {
	status := false
	message := "Hii"

	if len(requestData) > 0 {

		addChat := rp.SaveMessage(requestData["message"])
		if addChat > 0 {
			message = "message added"
			status = true
		} else {
			message = "Failed"
		}

	}
	return status, message
}

/**
 * For update the chat messages in to database
 * @input post message and update id from api
 * @return status ,message
 */

func UpdateChatMessage(requestData map[string]string) (bool, string) {
	status := false
	message := "Hii"

	fmt.Println(requestData)

	if len(requestData) > 0 {

		addChat := rp.UpdateMessage(requestData["message"], requestData["id"])
		if addChat > 0 {
			message = "message updated"
			status = true
		} else {
			message = "Failed"
		}

	}
	return status, message
}


/**
 * For delete the chat messages in to database
 * @input delete id from api
 * @return status ,message
 */

func DeleteChatMessage(requestData map[string]string) (bool, string) {
	status := false
	message := "Hii"

	fmt.Println(requestData)

	if len(requestData) > 0 {

		addChat := rp.DeleteChatMessage(requestData["id"])
		if addChat > 0 {
			message = "message deleted"
			status = true
		} else {
			message = "Failed"
		}

	}
	return status, message
}

/**
 * For listing all chat messages from the database
 * 
 * @return status ,chat lists
 */

func GetAllChatMessage() (bool, map[string]string) {
	status := false

	listChats := rp.GetAllChatMessage()
	if len(listChats) > 0 {
		status = true
	}

	return status, listChats
}
