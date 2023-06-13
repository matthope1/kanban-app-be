package auth0

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func GetUserInfo(accessToken string) {
	url := "https://" + os.Getenv("AUTH0_DOMAIN") + "/userinfo"
	
	fmt.Println("getting user data from this url", url)

	req, _ := http.NewRequest("GET", url, nil)

	fmt.Println("using this access token", accessToken)
	req.Header.Add("authorization", "Bearer " + " " + accessToken)
	req.Header.Add("Content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Response from getting auth0 user data: ")
	fmt.Println("I expect to see some fucking emails...")
	fmt.Println(res)
	fmt.Println(string(body))
}