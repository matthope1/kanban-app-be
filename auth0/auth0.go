package auth0

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kanban-app-be/types"
	"net/http"
	"os"
)

func GetUserInfo(accessToken string) types.UserInfo {
	// fmt.Println("get user info called")
	url := "https://" + os.Getenv("AUTH0_DOMAIN") + "/userinfo"

	// fmt.Println("getting user data from this url", url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-type", "application/json")

	// fmt.Println("request information", req)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	// fmt.Println("Response from getting auth0 user data: ")
	// fmt.Println(string(body))

	var result types.UserInfo
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		fmt.Println("Error unmarshaling data from request.")
	}

	// fmt.Println("testing print user email", result.Email)
	return result
}
