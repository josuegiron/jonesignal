package jonesignal

import (
	"fmt"
	"testing"

	config "github.com/api-eliab/eliab-config-go"
)

func TestSendNotification(t *testing.T) {

	notification := NotificationRequest{
		AppID:            config.Get.OneSignal.AppID,
		IncludePlayerIds: []string{"6646ca8d-ca15-4c25-a685-6ddb3bf977a0"},
		Headings:         Headings{En: "Hola", Es: "Hola"},
		Contents:         Contents{En: "como estas", Es: "test"},
	}

	messageID, err := SendNotification(notification)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf(messageID)

	// type args struct {
	// 	data NotificationRequest
	// }
	// tests := []struct {
	// 	name    string
	// 	args    args
	// 	want    string
	// 	wantErr bool
	// }{
	// 	{
	// 		name: "test1",
	// 		args: args{
	// 			data: notification,
	// 		},
	// 		wantErr: false,
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got, err := SendNotification(tt.args.data)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("SendNotification() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if got != tt.want {
	// 			t.Errorf("SendNotification() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
