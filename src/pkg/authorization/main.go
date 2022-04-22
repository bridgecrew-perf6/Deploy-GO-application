package authorization

import (
	"backend/pkg/dao"
	"backend/pkg/dao/youtube"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func RefreshToken() (*youtube.AccessAfterRefreshToken, error) {
	url := "https://accounts.google.com/o/oauth2/token"
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf("refresh_token=%s&client_id=%s&client_secret=%s&grant_type=refresh_token", youtube.RefreshToken, youtube.ClientId, youtube.ClientSecret))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var details youtube.AccessAfterRefreshToken
	if err := json.Unmarshal(body, &details); err != nil {
		log.Println("H", err.Error())
		return nil, errors.New(dao.ErrorInvalidUserData)
	}
	youtube.SetAccessToken(details.AccessToken)
	return &details, nil
}
