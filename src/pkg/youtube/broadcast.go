package youtube

import (
	"backend/pkg/authorization"
	"backend/pkg/dao/youtube"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func CreateBroadcast(title, description, privacyStatus string, startTime time.Time) (*youtube.Broadcast, error) {
	url := "https://www.googleapis.com/youtube/v3/liveBroadcasts?part=snippet,id,status,contentDetails"
	method := "POST"
	payload := strings.NewReader(`{
    "snippet":{
        "title":"` + title + `",
        "description":"` + description + `",
        "scheduledStartTime":"` + startTime.Format(time.RFC3339) + `"
    },
    "contentDetails":{
        "monitorStream": {
            "enableMonitorStream":true
        },
        "enableAutoStart":false,
        "enableAutoStop":false,
        "latencyPreference":"ultraLow"
    },
    "status":{
        "privacyStatus":"` + privacyStatus + `",
        "selfDeclaredMadeForKids":false,
        "recordingStatus":"recording"
    }
}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("CREATE REQ ERROR:", err.Error())
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
				fmt.Println("NOT ABLE TO REFRESH:", err.Error())
				return nil, err
			}
			req.Header.Add("Authorization", "Bearer "+atc.AccessToken)
			res, err := client.Do(req)
			if err != nil {
				fmt.Println("SEND REQ ERROR:", err.Error())
				return nil, err
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("REQ BODY ERROR:", err.Error())
				return nil, err
			}
			var details youtube.Broadcast
			if err := json.Unmarshal(body, &details); err != nil {
				log.Println("RE-API ATTEMPT:", err.Error())
				return nil, err
			}
			return &details, nil
		} else {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("REQ ERROR BODY:", err.Error())
				return nil, err
			}
			log.Println("LOG REQ ERROR BODY:", res, string(body))
			return nil, err
		}
	}
	if err != nil {
		fmt.Println("REQ SEND ERROR BODY:", err.Error())
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("RESP BODY READ:", err.Error())
		return nil, err
	}
	var details youtube.Broadcast
	if err := json.Unmarshal(body, &details); err != nil {
		log.Println("REQ ERROR BODY:", err.Error())
		return nil, err
	}
	return &details, nil
}
