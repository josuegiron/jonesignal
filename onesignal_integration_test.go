package jonesignal

import "testing"

func Test_addNewDevice(t *testing.T) {

	URL = "https://onesignal.com/api/v1/players"

	onesignal := OSAddDeviceReq{
		AppID:       "4c2687a0-f3c4-468f-88aa-fc1b34ff61e0",
		DeviceModel: "iPhone 8,3",
		Identifier:  "1244123",
	}

	type args struct {
		jsonData OSAddDeviceReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				jsonData: onesignal,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if v, err := addNewDevice(tt.args.jsonData); (err != nil) != tt.wantErr {
				t.Errorf("addNewDevice() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				t.Log(v)
			}
		})
	}
}
