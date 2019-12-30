package jonesignal

// URL ...
var URL string

// OSAddDeviceReq ...
type OSAddDeviceReq struct {
	AppID       string `json:"app_id"`
	DeviceModel string `json:"device_model"`
	DeviceOs    string `json:"device_os"`
	DeviceType  int    `json:"device_type"`
	GameVersion string `json:"game_version"`
	Identifier  string `json:"identifier"`
	Language    string `json:"language"`
	Tags        struct {
		A   string `json:"a"`
		Foo string `json:"foo"`
	} `json:"tags"`
	Timezone int `json:"timezone"`
}

// OSAddDeviceRes ...
type OSAddDeviceRes struct {
	Success bool   `json:"success"`
	ID      string `json:"id"`
}

// Notification ...
type Notification struct {
	AppID    string `json:"app_id"`
	Contents struct {
		En string `json:"en"`
	} `json:"contents"`
	Data struct {
		Foo string `json:"foo"`
	} `json:"data"`
	IncludedSegments []string `json:"included_segments"`
}
