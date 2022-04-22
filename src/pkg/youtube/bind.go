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

	"github.com/gofiber/fiber/v2"
)

type Bind struct {
	BroadCastId string `json:"broadCastId"`
	StreamId    string `json:"streamId"`
}

func BindStreamYoutube(c *fiber.Ctx) error {
	p := Bind{}
	if err := json.Unmarshal(c.Body(), &p); err != nil {
		log.Println(err)
	}

	data, err := BindStream(p.BroadCastId, p.StreamId)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": data,
	})
}

func BindStream(broadCastId, streamId string) (*youtube.Broadcast, error) {
	url := "https://www.googleapis.com/youtube/v3/liveBroadcasts/bind?part=snippet,id,status,contentDetails&id=" + broadCastId + "&streamId=" + streamId
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(dao.ErrorInvalidUserData)
	}
	req.Header.Add("Authorization", "Bearer "+youtube.AccessToken)
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
				return nil, errors.New(dao.ErrorInvalidUserData)
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return nil, errors.New(dao.ErrorInvalidUserData)
			}
			var details youtube.Broadcast
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
		return nil, errors.New(dao.ErrorInvalidUserData)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(dao.ErrorInvalidUserData)
	}
	var details youtube.Broadcast
	if err := json.Unmarshal(body, &details); err != nil {
		log.Println("H", err.Error())
		return nil, errors.New(dao.ErrorInvalidUserData)
	}
	return &details, nil
}
