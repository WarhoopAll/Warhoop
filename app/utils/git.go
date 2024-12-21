package utils

import (
	"encoding/json"
	"grimoire/app/log"
	"os"
)

var logger = log.Get()

type GitInfo struct {
	CommitHash string       `json:"commitHash"`
	Branch     string       `json:"branch"`
	CommitDate string       `json:"commitDate"`
	Version    string       `json:"lastTag"`
	Authors    []AuthorInfo `json:"authors"`
	Commits    []CommitDay  `json:"commits"`
}

type AuthorInfo struct {
	Name    string `json:"name"`
	Commits int    `json:"commits"`
}

type CommitDay struct {
	Day     string   `json:"day"`
	Tag     string   `json:"tag"`
	Commits []Commit `json:"commits"`
}

type Commit struct {
	Date    string `json:"date"`
	Author  string `json:"author"`
	Message string `json:"message"`
}

func LoadGitInfo(path string) (*GitInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		logger.Error("Failed to open gitinfo.json",
			log.String("err", err.Error()),
		)
		return nil, err
	}
	defer file.Close()

	var gitInfo GitInfo
	decoder := json.NewDecoder(file)
	decodeErr := decoder.Decode(&gitInfo)
	if decodeErr != nil {
		logger.Error("Failed to decode gitinfo.json",
			log.String("err", err.Error()),
		)
		return nil, err
	}

	return &gitInfo, nil
}
