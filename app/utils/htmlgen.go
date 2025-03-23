package utils

import (
	"fmt"
	"html/template"
	"os"
	"warhoop/app/config"
	"warhoop/app/log"
)

func LoadAndGenerateHTML(gitInfoPath string) error {
	gitInfo, err := LoadGitInfo(gitInfoPath)
	if err != nil {
		return err
	}

	data := GitInfo{
		CommitHash: gitInfo.CommitHash,
		Branch:     gitInfo.Branch,
		CommitDate: gitInfo.CommitDate,
		Authors:    gitInfo.Authors,
		Version:    gitInfo.Version,
		Commits:    gitInfo.Commits,
	}

	err = GenerateHTML(data)
	if err != nil {
		return err
	}

	return nil
}

func GenerateHTML(data GitInfo) error {
	tmplContent, err := os.ReadFile(config.Get().TemplateWelcome)
	if err != nil {
		log.Get().Error("utils.GenerateHTML",
			log.String("err", err.Error()),
		)
		return err
	}

	tmplParsed, err := template.New("welcome").Parse(string(tmplContent))
	if err != nil {
		log.Get().Error("utils.GenerateHTML",
			log.String("err", err.Error()),
		)
		return err
	}

	shortCommitHash := ""
	if len(data.CommitHash) >= 20 {
		shortCommitHash = data.CommitHash[:20]
	} else {
		shortCommitHash = data.CommitHash
	}

	var authorsList string
	for _, author := range data.Authors {
		authorsList += fmt.Sprintf("%d commits - %s\n", author.Commits, author.Name)
	}

	templateData := struct {
		ShortCommitHash string
		Branch          string
		CommitDate      string
		AuthorsList     string
		Version         string
		Commits         []CommitDay
	}{
		ShortCommitHash: shortCommitHash,
		Branch:          data.Branch,
		CommitDate:      data.CommitDate,
		AuthorsList:     authorsList,
		Version:         data.Version,
		Commits:         data.Commits,
	}

	err = os.MkdirAll("./static", os.ModePerm)
	if err != nil {
		log.Get().Error("utils.GenerateHTML",
			log.String("err", err.Error()),
		)
		return err
	}

	outputFile := config.Get().TemplateStatic
	file, err := os.Create(outputFile)
	if err != nil {
		log.Get().Error("utils.GenerateHTML",
			log.String("err", err.Error()),
		)
		return err
	}
	defer file.Close()

	err = tmplParsed.Execute(file, templateData)
	if err != nil {
		log.Get().Error("utils.GenerateHTML",
			log.String("err", err.Error()),
		)
		return err
	}
	log.Get().Debug("utils.GenerateHTML",
		log.String("directory", outputFile),
	)
	return nil
}
