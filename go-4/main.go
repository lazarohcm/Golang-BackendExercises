package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Repository structure
type Repository struct {
	Name        string `json:"full_name"`
	Description string `json:"description"`
	URL         string `json:"html_url"`
	Starts      int    `json:"stargazers_count"`
}

type JsonRepo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Starts      int    `json:"stars"`
}

func main() {
	_ = githubStars("go")
}

type Response struct {
	Name         string       `json:"name"`
	Repositories []Repository `json:"items"`
}

func githubStars(lang string) error {
	url := fmt.Sprintf("https://api.github.com/search/repositories?q=language:golang&sort=stars&order=desc&page=1&per_page=10")
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}

	defer resp.Body.Close()

	var responseObject Response

	responseData, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(responseData, &responseObject)

	if err != nil {
		log.Fatal(err)
	}
	var jsonRepos []JsonRepo
	for _, rep := range responseObject.Repositories {
		jsonRepos = append(jsonRepos, JsonRepo{Name: rep.Name, Description: rep.Description, URL: rep.URL, Starts: rep.Starts})
	}

	b, err := json.Marshal(jsonRepos)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	d1 := []byte(b)
	err = ioutil.WriteFile("./stars.json", d1, 0644)

	return nil
}
