package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"random-stuff-service/rest"

	"github.com/gin-gonic/gin"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Joke struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}

func Get(config rest.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		person, err := getPerson(config.Client)
		if err != nil {
			response := gin.H{"error": err.Error()}
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		joke, err := getJoke(config.Client)
		if err != nil {
			response := gin.H{"error": err.Error()}
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		sentence := fmt.Sprintf("%s %s's %s", person.FirstName, person.LastName, joke.Value.Joke)
		ctx.String(http.StatusOK, sentence)
	}
}

func getPerson(client *http.Client) (Person, error) {
	url := "https://names.mcquay.me/api/v0"
	body, err := getResponseBody(url, client)
	if err != nil {
		return Person{}, err
	}

	var person Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		return Person{}, err
	}

	return person, nil
}

func getJoke(client *http.Client) (Joke, error) {
	//3 QUERY PARAMS: limitTo, firstName, lastName"
	url := "http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=John&lastName=Doe"
	body, err := getResponseBody(url, client)
	if err != nil {
		return Joke{}, err
	}

	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		return Joke{}, err
	}

	return joke, nil
}

func getResponseBody(url string, client *http.Client) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
