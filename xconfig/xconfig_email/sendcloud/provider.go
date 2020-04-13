package sendcloud

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pubgo/g/pkg/timeutil"
)

type Provider struct {
	Name     string
	ApiUser  string
	ApiKey   string
	From     string
	FromName string
}

func ProviderName() string {
	return s.Name
}

func Assembly(to, templateTitle, templateContent string) string {
	params := url.Values{
		"apiUser":  {s.ApiUser},
		"apiKey":   {s.ApiKey},
		"from":     {s.From},
		"fromName": {s.FromName},
		"to":       {to},
		"subject":  {templateTitle},
		"html":     {templateContent},
	}

	return params.Encode()
}

type emailStatus struct {
	EmailId      string `json:"emailId"`
	Status       string `json:"status"`
	ApiUser      string `json:"apiUser"`
	Recipients   string `json:"recipients"`
	RequestTime  string `json:"requestTime"`
	ModifiedTime string `json:"modifiedTime"`
	SendLog      string `json:"sendLog"`
}

type voList struct {
	Total      string        `json:"total"`
	VoListSize int           `json:"voListSize"`
	VoList     []emailStatus `json:"voList"`
}

type statusResult struct {
	Result     bool   `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Info       voList `json:"info"`
}

func GetStatus(thirdEmailId string) (string, error) {

reGetEmailStatus:
	params := url.Values{
		"apiUser":   {s.ApiUser},
		"apiKey":    {s.ApiKey},
		"startDate": {timeutil.GetSystemCurDate()},
		"endDate":   {timeutil.GetSystemCurDate()},
		"emailIds":  {thirdEmailId},
	}

	url := "http://api.sendcloud.net/apiv2/data/emailStatus"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(params.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		goto reGetEmailStatus
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var locStatusResult statusResult
	err = json.Unmarshal([]byte(bodyContent), &locStatusResult)
	if err != nil {
		return "", err
	}

	if len(locStatusResult.Info.VoList) == 0 {
		return "", err
	}

	updateEmailStatus := locStatusResult.Info.VoList[0].Status

	return updateEmailStatus, nil
}
