package hackernewsapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var apiEndpoint = "https://hacker-news.firebaseio.com/v0/"

type Ids []int

type MaxItem int

type User struct {
	ID        int    `json:"id"`
	Delay     int    `json:"delay"`
	Created   int    `json:"created"`
	Karma     int    `json:"karna"`
	About     string `json:"about"`
	Submitted []int  `json:"submitted"`
}

type Story struct {
	By          string `json:"by"`
	Dead        bool   `json:"dead"`
	Deleted     bool   `json:"deleted"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Parent      int    `json:"parent"`
	Parts       []int  `json:"parts"`
	Score       int    `json:"score"`
	Text        string `json:"text"`
	Time        int64  `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

type Updates struct {
	Items    []int    `json:"items"`
	Profiles []string `json:"profiles"`
}

// GetUser returns User
func GetUser(name string) (User, error) {

	var data User

	body, _ := getJSON(apiEndpoint + "user/" + name + ".json")
	json.NewDecoder(body).Decode(&data)
	defer body.Close()

	return data, nil
}

// GetStories returns []int ids
func GetStories(target string) ([]int, error) {

	var stories Ids
	switch strings.ToLower(target) {
	case "new", "top", "job", "ask":
		body, _ := getJSON(apiEndpoint + target + "stories.json")
		err := json.NewDecoder(body).Decode(&stories)
		defer body.Close()

		if err != nil {
			return nil, err
		}

	}
	return stories, nil
}

// GetItem returns item (story, comment, ask, job, poll, pollopt)
func GetItem(id int) (Story, error) {

	body, _ := getJSON(apiEndpoint + "item/" + strconv.Itoa(id) + ".json")
	defer body.Close()

	var data Story
	json.NewDecoder(body).Decode(&data)
	return data, nil
}

// GetMaxItemID returns hackernews max item id (LIVE)
func GetMaxItemID() (int, error) {
	body, _ := getJSON(apiEndpoint + "maxitem.json")
	defer body.Close()
	var num MaxItem
	json.NewDecoder(body).Decode(&num)
	return (int)(num), nil
}

// GetUpdates returns item and profile changes
func GetUpdates() (Updates, error) {

	var data Updates

	body, _ := getJSON(apiEndpoint + "updates.json")
	json.NewDecoder(body).Decode(&data)
	defer body.Close()

	return data, nil
}

func getJSON(url string) (io.ReadCloser, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Close = true // connection reset

	client := new(http.Client)
	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return response.Body, nil
}
