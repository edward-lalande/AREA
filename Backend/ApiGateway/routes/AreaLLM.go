package routes

import (
	models "api-gateway/Models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gage-technologies/mistral-go"
	"github.com/gin-gonic/gin"
)

func getAction(c *gin.Context, request string) *json.RawMessage {

	client := mistral.NewMistralClientDefault("HCZtt06PoHGhiRb1Wjed8b5v4zeQ874H")

	actions, err := http.Get("http://127.0.0.1:8080/actions")

	if err != nil {
		return nil
	}
	defer actions.Body.Close()

	body, err := io.ReadAll(actions.Body)

	if err != nil {
		return nil
	}

	fmt.Println("request: " + request + "\n")

	prompt := "Tu es un agent pour une application de création d'action et de réaction." +
		"Ton objectif va être d'identifier seulement l'action dans la demande de l'utilisateur : \n###\n" + request + "\n###\n" +
		"Tu dois me renvoyer une réponse JSON dans ce format précis : \n###\n { action_id: $, action_type: $, name: $, ... } + \n###\n " +
		"Tu dois donc remplacer les $ par l'id et le type de l'action et les ... par les arguments (le nom et la valeur)." +
		"Afin de trouver les bons ids, types, names et arguments de l'action, tu dois te servir d'une liste de toutes les actions possibles : \n###\n" + string(body) + "\n###\n" +
		"Si dans les arguments il y a un tableau vide ([]) alors ne met rien." +
		"Example: \n###\n{ action_id: 0, action_type: 0, name: 'Every day at', hour: 17, minute: 45 }\n###\n" +
		"Envoie moi SEULEMENT l'objet JSON, je ne veux aucun autre texte."

	chatRes, err := client.Chat("mistral-tiny", []mistral.ChatMessage{{Content: prompt, Role: mistral.RoleUser}}, nil)
	if err != nil {
		fmt.Println("error chat completion action")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting chat completion"})
	}

	chatRes.Choices[0].Message.Content = strings.Replace(chatRes.Choices[0].Message.Content, "`", "", -1)
	chatRes.Choices[0].Message.Content = strings.Replace(chatRes.Choices[0].Message.Content, "\n", "", -1)

	fmt.Println("action: " + chatRes.Choices[0].Message.Content + "\n")

	raw := json.RawMessage(chatRes.Choices[0].Message.Content)

	fmt.Println("action: " + chatRes.Choices[0].Message.Content + "\n")

	return &raw

}

func getReaction(c *gin.Context, request string) *json.RawMessage {

	client := mistral.NewMistralClientDefault("HCZtt06PoHGhiRb1Wjed8b5v4zeQ874H")

	actions, err := http.Get("http://127.0.0.1:8080/reactions")

	if err != nil {
		return nil
	}
	defer actions.Body.Close()

	body, err := io.ReadAll(actions.Body)

	if err != nil {
		return nil
	}

	prompt := "Tu es un agent pour une application de création d'action et de réaction." +
		"Ton objectif va être d'identifier seulement l'reaction dans la demande de l'utilisateur : \n###\n" + request + "\n###\n" +
		"Tu dois me renvoyer une réponse JSON dans ce format précis : \n###\n { reaction_id: $, reaction_type: $, name: $,  ... } + \n###\n " +
		"Tu dois donc remplacer les $ par l'id et le type de la reaction et les ... par les arguments (le nom et la valeur)." +
		"Afin de trouver les bons ids, types, names et arguments de la reaction, tu dois te servir d'une liste de toutes les reactions possibles : \n###\n" + string(body) + "\n###\n" +
		"Example: \n###\n[{ reaction_id: 0, reaction_type: 0, name: 'Send message on channel', hour: 17, minute: 45 }]\n###\n" +
		"Envoie moi SEULEMENT l'objet JSON, je ne veux aucun autre texte."

	chatRes, err := client.Chat("mistral-tiny", []mistral.ChatMessage{{Content: prompt, Role: mistral.RoleUser}}, nil)
	if err != nil {
		fmt.Println("error chat completion reaction")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting chat completion"})
	}

	chatRes.Choices[0].Message.Content = strings.Replace(chatRes.Choices[0].Message.Content, "`", "", -1)
	chatRes.Choices[0].Message.Content = strings.Replace(chatRes.Choices[0].Message.Content, "\n", "", -1)

	fmt.Println("reaction: " + chatRes.Choices[0].Message.Content + "\n")

	raw := json.RawMessage(chatRes.Choices[0].Message.Content)

	fmt.Println("reaction: " + chatRes.Choices[0].Message.Content + "\n")

	return &raw

}

func AreaLLM(c *gin.Context) {
	var payload models.PayloadLLM

	if err := c.ShouldBindJSON(&payload); err != nil {
		fmt.Println("invalid payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	var area []models.PayloadItem = make([]models.PayloadItem, 1)

	area[0] = models.PayloadItem{
		UserToken: payload.UserToken,
		Action:    getAction(c, payload.Request),
		Reactions: make([]*json.RawMessage, 1),
	}

	time.Sleep(8 * time.Second)

	area[0].Reactions[0] = getReaction(c, payload.Request)

	jsonBytes, err := json.Marshal(area)
	if err != nil {
		fmt.Println("invalid area")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid area"})
		return
	}

	_, err = http.Post("http://127.0.0.1:8080/areas", "application/jsons", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("creating area error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Creating area"})
		return
	}

}
