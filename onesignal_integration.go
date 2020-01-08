package jonesignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	config "github.com/api-eliab/eliab-config-go"
)

func addNewDevice(jsonData OSAddDeviceReq) (OSAddDeviceRes, error) {

	var res OSAddDeviceRes

	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(config.Get.Services["OS_AddDevice"].URL, "application/json", bytes.NewBuffer(jsonValue))

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

func createNotification() {

}
