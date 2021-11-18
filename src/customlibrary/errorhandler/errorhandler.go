package errorhandler

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	ticket "prod/src/customlibrary/jira"
	"runtime"
	"time"
)

var DbCon *sql.DB

/*
*Error handing function
*Handle err using log and jira tickte
*@input error
 */
func ErrorCheck(err error) {
	if err != nil {

		orignalError := errors.New(fmt.Sprint(err))
		errorText := fmt.Sprintf(" %+v \n\n", orignalError)
		log.Println(errorText)

		go GenerateCommonTickets(errorText)
		time.Sleep(3 * time.Second)
		runtime.Goexit()
	}
}

/*
*generate jira ticket as error
*
*@input error
 */

func GenerateCommonTickets(errorText string) {
	var name string
	var email string
	descriptionData := map[string]interface{}{}
	name = "Megha"
	email = "meghaa@salesjio.com"
	descriptionData["Reason"] = errorText
	DataConfigs := map[string]interface{}{
		"name":            name,
		"email":           email,
		"subject":         "Chat API Error",
		"priority":        4,
		"descriptionData": descriptionData,
	}
	ticket.GenerateCommonTicketData(DataConfigs)

}
