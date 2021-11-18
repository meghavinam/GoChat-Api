/*

Descreption : This project is to handile all chat apis
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	chatAct "prod/src/customlibrary/chatActions"
	cg "prod/src/customlibrary/configuration"
	sr "prod/src/customlibrary/services"
	"strconv"
	"strings"
)

func getPostData(r *http.Request) map[string]string {

	decoder := json.NewDecoder(r.Body)
	//	fmt.Println("r.decoder", decoder)
	t := map[string]interface{}{}
	err := decoder.Decode(&t)
	if err != nil {

		fmt.Println(r)
		log.Println(err)
	}
	postdata := map[string]string{}
	for key, val := range t {
		postdata[fmt.Sprint(key)] = fmt.Sprint(val)
	}
	return postdata
}

func SaveChatBoatMessage(w http.ResponseWriter, r *http.Request) {

	postdata := getPostData(r)
	status, response := chatAct.SaveChatBoatMessage(postdata)
	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(400)
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"status": ` + strconv.FormatBool(status) + `, "message":"` + response + `"}`))
}

func UpdateChatMessage(w http.ResponseWriter, r *http.Request) {

	postdata := getPostData(r)
	status, response := chatAct.UpdateChatMessage(postdata)
	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(400)
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"status": ` + strconv.FormatBool(status) + `, "message":"` + response + `"}`))
}

func DeleteChatMessage(w http.ResponseWriter, r *http.Request) {

	postdata := getPostData(r)
	status, response := chatAct.DeleteChatMessage(postdata)
	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(400)
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"status": ` + strconv.FormatBool(status) + `, "message":"` + response + `"}`))
}

func GetAllChatMessage(w http.ResponseWriter, r *http.Request) {

	status, response := chatAct.GetAllChatMessage()
	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(400)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func getCode(r *http.Request, defaultCode int) string {

	p := strings.Split(r.URL.Path, "/")
	if len(p) > defaultCode {
		return p[defaultCode]
	}
	return ""
}

func getUrlData(r *http.Request) map[string]string {

	t := r.URL.Query()
	postdata := map[string]string{}
	for key, val := range t {
		postdata[fmt.Sprint(key)] = val[0]
	}
	return postdata
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
</head>
<body>
<p>Hello world</p>
</body>
</html>`))
}

func main() {

	cg.SetConfigParams()
	sr.SetDbConnection("ClientDatabase", cg.Config.ClientDatabase.MaxOpenConnections, cg.Config.ClientDatabase.MaxIdleConnections)

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/save/chat", SaveChatBoatMessage)
	http.HandleFunc("/update/chat", UpdateChatMessage)
	http.HandleFunc("/delete/chat", DeleteChatMessage)
	http.HandleFunc("/list/chat", GetAllChatMessage)

	fmt.Printf("Starting server at port " + cg.Config.Port1 + "\n")
	if err := http.ListenAndServe(":"+cg.Config.Port1, nil); err != nil {
		log.Fatal(err)
	}
}
