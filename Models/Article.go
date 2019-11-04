package Models

import "encoding/json"

type Article struct {
	ID       int `json:"ID"`
	Parent   *Article
	Owner    User   `json:"Owner"`
	Title    string `json:"Title"`
	Comments []Article
}

func (a *Article) Json() (string, error) {
	data, err := json.Marshal(a)
	return string(data), err
}
