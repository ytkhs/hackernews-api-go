package hackernewsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var apiEndpoint = "https://hacker-news.firebaseio.com/v0/"

var httpClient = &http.Client{Timeout: 10 * time.Second}

type Ids []int

type MaxItem int

type User struct {
	ID        string `json:"id"`
	Delay     int    `json:"delay"`
	Created   int64  `json:"created"`
	Karma     int    `json:"karma"`
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
	err := getJSON(apiEndpoint+"user/"+name+".json", &data)
	return data, err
}

// GetStories returns []int ids
func GetStories(target string) ([]int, error) {
	t := strings.ToLower(target)
	switch t {
	case "new", "top", "job", "ask":
	default:
		return nil, fmt.Errorf("hackernewsapi: unknown story target %q", target)
	}

	var stories Ids
	if err := getJSON(apiEndpoint+t+"stories.json", &stories); err != nil {
		return nil, err
	}
	return stories, nil
}

// GetItem returns item (story, comment, ask, job, poll, pollopt)
func GetItem(id int) (Story, error) {
	var data Story
	err := getJSON(apiEndpoint+"item/"+strconv.Itoa(id)+".json", &data)
	return data, err
}

// GetMaxItemID returns hackernews max item id (LIVE)
func GetMaxItemID() (int, error) {
	var num MaxItem
	err := getJSON(apiEndpoint+"maxitem.json", &num)
	return int(num), err
}

// GetUpdates returns item and profile changes
func GetUpdates() (Updates, error) {
	var data Updates
	err := getJSON(apiEndpoint+"updates.json", &data)
	return data, err
}

func getJSON(url string, v interface{}) error {
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("hackernewsapi: GET %s: %s", url, resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("hackernewsapi: decode %s: %w", url, err)
	}
	return nil
}
