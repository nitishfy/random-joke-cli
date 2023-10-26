package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type Joke struct {
	ID string `json: "id"`
	Joke string `json: "joke"`
	Status int `json: "status"`

}

// randomJokeCmd represents the randomJoke command
var randomJokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "generate random joke",
	Long: `This command is used to generate random funny jokes`,
	Run: func(cmd *cobra.Command, args []string) {
		generateRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomJokeCmd)
}

func RandomJoke(baseURL string) []byte {
	req, err := http.NewRequest(http.MethodGet,baseURL,nil)
	if err != nil  {
		log.Println("Error generating a request - %v:", err)
	}

	req.Header.Add("Accept","application/json")
	req.Header.Add("User-Agent","random joke cli (https://github.com/nitishfy/random-joke-cli)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error hitting the request to the server: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading the content of the response: %v", err)
	}

	return body
}

func generateRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := RandomJoke(url)
	joke := Joke{}

	err := json.Unmarshal(responseBytes,&joke)
	if err != nil {
		log.Println("error unmarshalling json: %v", err)
	}

	fmt.Printf(string(joke.Joke))
}
