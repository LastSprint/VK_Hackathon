package notifications

import (
	"encoding/json"
	"fmt"
	"log"
	"suncity/auth"
	"suncity/commod"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sirupsen/logrus"
)

func SendNotification(fromUser *commod.ServiceUser, toUser *commod.ServiceUser, payload string) {

	fmt.Println("SEND START")

	cert, err := certificate.FromP12File("cert.p12", "")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	fmt.Println("CERT CREATED")

	notification := &apns2.Notification{}
	notification.DeviceToken = toUser.Apns
	notification.Topic = "surf.vk.hackathon.sun.city"

	jsr := map[string]interface{}{
		"aps": map[string]interface{}{
			"alert": map[string]string{
				"title": fromUser.Name,
				"body":  payload,
			},
		},
	}

	data, _ := json.Marshal(jsr)

	notification.Payload = []byte(data) // See Payload section below

	fmt.Println("NOTIF START")

	client := apns2.NewClient(cert).Development()

	fmt.Println("CLIENT START")

	res, err := client.Push(notification)

	fmt.Println("PUSH START")

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}

func SendNotificationByUserId(sender *commod.ServiceUser, recivier string, data string) {
	usr, err := auth.AuthRepo.GetUserById(recivier)

	if err != nil {
		logrus.Errorln(err)
	}

	SendNotification(sender, usr, data)
}
