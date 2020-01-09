package jonesignal

import "errors"

// AddDevice ...
func AddDevice(data OSAddDeviceReq) (string, error) {

	response, err := addNewDevice(data)
	if err != nil {
		return "", err
	}

	if !response.Success {
		return "", errors.New("No pudo agregar el dispositivo")
	}

	return response.ID, nil

}

// SendNotification ...
func SendNotification(data NotificationRequest) (string, error) {

	response, err := sendNotification(data)
	if err != nil {
		return "", err
	}

	if response.ID == "" {
		return "", errors.New("No pudo enviar el mensaje")
	}

	return response.ID, nil

}
