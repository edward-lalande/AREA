package area

import (
	"fmt"
	"net/http"
	models "twilio/Models"

	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/rest/api/v2010"
)

func sendMessage(information models.Reactions) (*http.Response, error) {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:+33685982932")
	params.SetFrom("whatsapp:+12186558783")
	params.SetBody("Hello from Golang!")

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Message sent successfully!")
	}
}

func FindReactions(id int, information models.Reactions) (*http.Response, error) {

	var Reactions map[int]func(models.Reactions) (*http.Response, error) = make(map[int]func(models.Reactions) (*http.Response, error))

	Reactions[0] = sendMessage

	return Reactions[id](information)
}
