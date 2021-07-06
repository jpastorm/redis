package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


type WrapperResponse struct {
	Data struct {
		User struct{
			ID    uint   `json:"id"`
			Email string `json:"email"`
		} `json:"user"`
		Token     string `json:"token"`
	} `json:"data"`

}

func Execute(api string) (map[string]interface{}, string, error) {

	url := url.URL{
		Scheme: "https",
		Host:   api,
	}

	httpposturl := url.String() + "/client/v1/login"

	fmt.Println("HTTP JSON POST URL:", httpposturl)

	var jsonData = []byte(`{
		"email": "josepastormendoza32+test7@gmail.com",
		"password": "secret"
	}`)

	request, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	if err != nil {
		return map[string]interface{}{},err.Error() , err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return map[string]interface{}{}, response.Status, err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	//fmt.Println("response Body:", string(body))
	var model WrapperResponse
	err = json.Unmarshal(body, &model)
	if err != nil {
		return map[string]interface{}{}, response.Status, err
	}

	return map[string]interface{}{
		"email": model.Data.User.Email,
		"token": model.Data.Token,
	}, response.Status ,nil
}
