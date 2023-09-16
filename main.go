package main

import (
	"encoding/json"
	"fmt"
	models2 "study/database/models"
	"study/database/queries"
	"study/database/schemas"
)

func main() {

	userData := schemas.UserSchema{
		FirstName: "Ivan",
		LastName:  "Petrov",
		Age:       40,
		Balance:   "3000",
	}

	fmt.Printf("Inserting data: %#v\n", userData)
	newUser := queries.AddUser(&userData)

	fmt.Printf("Getting data: userID : %#v\n", newUser.UserID)
	qData := queries.GetUser(newUser.UserID)

	fmt.Printf("Got data: %s\n", qData)

	js := models2.ToJsonString(qData)
	fmt.Printf("json: %s\n", js)

	var mas models2.UserModel
	_ = json.Unmarshal([]byte(js), &mas)
	fmt.Printf("%T %v", mas, mas)
}
