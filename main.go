package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	UserName    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
}

var UserList []User

func init() {

	UserJson := `[
		{
			"username":"aaasss555", 
			"phone_number":"0912323420"
		},
		{
			"username":"ss0987767665", 
			"phone_number":"0987767665"
		},
		{
			"username":"ss0981123456", 
			"phone_number":"0981123456"
		},
		{
			"username":"ss0981112234", 
			"phone_number":"0981112234"
		}
	]`
	err := json.Unmarshal([]byte(UserJson), &UserList)
	if err != nil {
		log.Fatal(err)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

	UserJson, err := json.Marshal(UserList)
	switch r.Method {

	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(UserJson)

	case http.MethodPost:
		var newUser User
		BodyByte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(BodyByte)
			return
		}
		err = json.Unmarshal(BodyByte, newUser)
		if err != nil {
			// log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		UserList = append(UserList, newUser)
		w.WriteHeader(http.StatusCreated)
		// w.Write(newUser)

		return
	}

}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, welcome to %s", r.URL.Host)
	})

	http.HandleFunc("/user", UserHandler)
	http.ListenAndServe(":5050", nil)

	// customer.FindRecords()

	// customer.Adduser("aassdd123456", "0321324343")
	// customer.Getuser("aassdd123456")

	// customer.Adduser("aaasss44555", "0912323422")
	// customer.Getuser("aaasss44555")

}
