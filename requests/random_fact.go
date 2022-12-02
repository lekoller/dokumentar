package requests

import (
	"encoding/json"
	"net/http"
)

type randomFact struct {
	Text string `json:"text"`
}

func GetRandomFact(client *http.Client) (fact randomFact, err error) {
	res, err := client.Get("https://uselessfacts.jsph.pl/random.json?language=en")
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&fact)
	return
}
