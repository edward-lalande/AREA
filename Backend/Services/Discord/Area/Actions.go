package area

import (
	models "discord-service/Models"
	"net/http"
)

func ReceiveMessage(informations models.Actions) (*http.Response, error) {
	return nil, nil
}

func FindActions(id int, informations models.Actions) (*http.Response, error) {
	var Actions map[int]func(models.Actions) (*http.Response, error) = make(map[int]func(models.Actions) (*http.Response, error))
	Actions[0] = ReceiveMessage

	return Actions[id](informations)
}
