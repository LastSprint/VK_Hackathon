package notifications

import (
	"fmt"
	"log"
	"suncity/commod"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
)

func SendNotification(fromUser *commod.ServiceUser, toUser *commod.ServiceUser) {

	fmt.Println("SEND START")

	cert, err := certificate.FromP12File("cert.p12", "")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	fmt.Println("CERT CREATED")

	notification := &apns2.Notification{}
	notification.DeviceToken = toUser.Apns
	notification.Topic = "surf.vk.hackathon.sun.city"

	str := `{
		"aps": {
		  "alert": "Breaking News! Антон пиздабол",
		  "sound": "default",
		  "link_url": "https://raywenderlich.com"
		}
	  }`
	notification.Payload = []byte(str) // See Payload section below

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
