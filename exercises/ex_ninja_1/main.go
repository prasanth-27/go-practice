package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First string `json:"FirstName"`
	Age   int    `json:"Age"`
}

func main() {
	fmt.Println("Json Marshal and unmarshall")

	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "MoneyPenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	userEncoder := json.NewEncoder(os.Stdout)

	err := userEncoder.Encode(users)
	if err != nil {
		fmt.Printf("Error: marshalling []user\n Error:%v", err)
	}

	jsonText := `[{"FirstName":"James","Age":32},{"FirstName":"MoneyPenny","Age":27},{"FirstName":"M","Age":54}]`

	usersBytes := []byte(jsonText)
	unmarshalledUserslice := []user{}

	//unmarshall takes pointer to output var(2nd arg)
	err = json.Unmarshal(usersBytes, &unmarshalledUserslice)

	if err != nil {
		fmt.Printf("Failed to unmarshal\nError:%v", err)
	} else {
		fmt.Printf("%#v", unmarshalledUserslice)
	}

}
