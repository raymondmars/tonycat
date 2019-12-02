package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"tonycat/config"
)

type InputMessage struct {
	ReqType     int `json:"reqType"`
	*Perception `json:"perception"`
	*UserInfo   `json:"userInfo"`
}

type Perception struct {
	*InputText `json:"inputText"`
}

type InputText struct {
	Text string `json:"text"`
}

type UserInfo struct {
	ApiKey string `json:"apiKey"`
	UserId string `json:"userId"`
}

type OutPutMessage struct {
	Intent  *OutIntent `json:"intent"`
	Results []Results  `json:"results"`
}

type OutIntent struct {
	Code int `json:"code"`
}

type Results struct {
	ResultType string            `json:"resultType"`
	Values     map[string]string `json:"values"`
	GroupType  int               `json:"groupType"`
}

type TuringBot struct{}

func (tb *TuringBot) Chat(message string) string {
	input := InputMessage{
		ReqType: 0,
		Perception: &Perception{
			InputText: &InputText{Text: message},
		},
		UserInfo: &UserInfo{ApiKey: config.TURING_APIKEY, UserId: config.TURING_USERID},
	}

	b, _ := json.Marshal(input)
	// log.Println("send:", string(b))

	client := &http.Client{}
	request, _ := http.NewRequest("POST", config.TURING_OPEN_API, bytes.NewBuffer(b))
	request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		log.Println("bot return:", string(body))
		outMessage := OutPutMessage{}
		outMessage.Results = make([]Results, 0)
		err := json.Unmarshal(body, &outMessage)
		if err != nil {
			log.Printf("err: %v", err)
		}
		// fmt.Println(outMessage.Results[0].Values["text"])
		if len(outMessage.Results) > 0 {
			txt := outMessage.Results[0].Values["text"]
			if txt != "" {
				return txt
			} else {
				return outMessage.Results[0].Values["url"]
			}
		} else {
			return ""
		}

	} else {
		return ""
	}

}
