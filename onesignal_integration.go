package jonesignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	config "github.com/api-eliab/eliab-config-go"
	"go.mnc.gt/log"
)

func addNewDevice(jsonData OSAddDeviceReq) (OSAddDeviceRes, error) {

	var res OSAddDeviceRes

	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(config.Get.Services["OS_AddDevice"].URL, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Errorf("The HTTP request failed with error %s\n", err)
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

func sendNotification(jsonData NotificationRequest) (NotificationResponse, error) {

	var res NotificationResponse

	jsonValue, _ := json.Marshal(jsonData)

	req, err := http.NewRequest("POST", config.Get.Services["OS_SendMessage"].URL, bytes.NewBuffer(jsonValue))
	if err != nil {
		return res, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", config.Get.OneSignal.Key))

	var c = &http.Client{}

	response, err := c.Do(req)

	if err != nil {
		log.Errorf("The HTTP request failed with error %s\n", err)
		return res, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err
	}

	log.Info(string(data))

	if err := json.Unmarshal(data, &res); err != nil {
		return res, err
	}

	return res, nil

}
