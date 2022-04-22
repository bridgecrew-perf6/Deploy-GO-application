package youtube

import (
	"backend/pkg/authorization"
	"backend/pkg/dao"
	"backend/pkg/dao/youtube"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// func CreateStreamYoutube(c *fiber.Ctx) (*youtube.Stream, error) {
// 	var data *youtube.Stream
// 	title := c.Params("title")
// 	CreateStream(title)
// 	return data, nil

// }

func CreateStreamYoutube(c *fiber.Ctx) error {
	title := c.Params("title")
	data, err := CreateStream(title)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	fmt.Println(data)
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Stream create successfully!",
	})
}

func CreateStream(title string) (*youtube.Stream, error) {
	url := "https://www.googleapis.com/youtube/v3/liveStreams?part=snippet,cdn,contentDetails,status"
	method := "POST"
	payload := strings.NewReader(`{
    "snippet":{
        "title":"` + title + `"
    },
    "cdn":{
        "frameRate":"variable",
        "ingestionType":"rtmp",
        "resolution":"variable"
    },
    "contentDetails":{
        "isReusable":true
    },
    "status":{
        "streamStatus":"active"
    }
}`)
	log.Println(payload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+youtube.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if res != nil && res.StatusCode != 200 {
		if res.StatusCode == 401 {
			// Refresh token and recall.
			atc, err := authorization.RefreshToken()
			if err != nil {
				fmt.Println("NOT ABLE TO REFRESH:", err)
				return nil, err
			}
			req.Header.Add("Authorization", "Bearer "+atc.AccessToken)
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
			log.Println(string(body))
			var details youtube.Stream
			if err := json.Unmarshal(body, &details); err != nil {
				log.Println("H", err.Error())
				return nil, errors.New(dao.ErrorInvalidUserData)
			}
			return &details, nil
		} else {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			log.Println(res, string(body))
			return nil, err
		}
	}
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
	log.Println(string(body))
	var details youtube.Stream
	if err := json.Unmarshal(body, &details); err != nil {
		log.Println("H", err.Error())
		return nil, errors.New(dao.ErrorInvalidUserData)
	}
	return &details, nil
}
