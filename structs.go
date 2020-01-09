package jonesignal

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

// NotificationRequest ...
type NotificationRequest struct {
	AppID            string   `json:"app_id"`
	IncludedSegments []string `json:"included_segments"`
	IncludePlayerIds []string `json:"include_player_ids"`
	Data             struct {
		Foo string `json:"foo"`
	} `json:"data"`
	Headings  Headings `json:"headings"`
	Contents  Contents `json:"contents"`
	LargeIcon string   `json:"large_icon"`
}

// Headings doc ...
type Headings struct {
	En string `json:"en"`
	Es string `json:"es"`
}

// Contents doc ...
type Contents struct {
	En string `json:"en"`
	Es string `json:"es"`
}

// NotificationResponse doc ...
type NotificationResponse struct {
	ID         string      `json:"id"`
	Recipients int         `json:"recipients"`
	ExternalID interface{} `json:"external_id"`
}
