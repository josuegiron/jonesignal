package jonesignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func addNewDevice(jsonData OSAddDeviceReq) (OSAddDeviceRes, error) {

	var res OSAddDeviceRes

	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(PlayersService, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return res, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return res, err
	}

	return res, nil
}

func sendNotification(jsonData NotificationRequest, oneSignalKey string) (NotificationResponse, error) {

	var res NotificationResponse

	jsonValue, _ := json.Marshal(jsonData)

	req, err := http.NewRequest("POST", NotificationsService, bytes.NewBuffer(jsonValue))
	if err != nil {
		return res, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", oneSignalKey))

	var c = &http.Client{}

	response, err := c.Do(req)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return res, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err
	}

	fmt.Print(string(data))

	if err := json.Unmarshal(data, &res); err != nil {
		return res, err
	}

	return res, nil

}
