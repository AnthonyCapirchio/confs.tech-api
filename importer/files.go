package importer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"squidward.confs.tech/conference"
)

func CategoryFromPath(path string) string {
	splitted := strings.Split(path, "/")
	fileName := splitted[len(splitted)-1]
	return strings.TrimSuffix(fileName, ".json")
}

type ConferenceFileLine struct {
	Name       string `json:"name"`
	URL        string `json:"url"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	City       string `json:"city"`
	Country    string `json:"country"`
	CfpUrl     string `json:"cfpUrl"`
	CfpEndDate string `json:"cfpEndDate"`
	Twitter    string `json:"twitter"`
	Category   string `json:"category"`
}

type LocalFileImporter struct {
	BasePath string
	Years    []string
}

func loadFile(path string, conferences chan conference.Conference, wg *sync.WaitGroup) {
	defer wg.Done()

	if !strings.HasSuffix(path, ".json") {
		return
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	HandleFileContent(f, path, conferences)
}

func HandleFileContent(f []byte, path string, conferences chan conference.Conference) {
	data := []ConferenceFileLine{}

	_ = json.Unmarshal([]byte(f), &data)
	category := CategoryFromPath(path)

	layout := "2006-01-02"

	for _, entry := range data {
		startDate, _ := time.Parse(layout, entry.StartDate)
		endDate, _ := time.Parse(layout, entry.EndDate)
		cfpEndDate, _ := time.Parse(layout, entry.CfpEndDate)

		line := conference.Conference{
			Name:       entry.Name,
			URL:        entry.URL,
			StartDate:  startDate,
			EndDate:    endDate,
			City:       entry.City,
			Country:    entry.Country,
			CfpUrl:     entry.CfpUrl,
			CfpEndDate: cfpEndDate,
			Twitter:    entry.Twitter,
			Category:   category,
		}

		conferences <- line
	}
}

func (l *LocalFileImporter) LoadFiles() chan conference.Conference {
	var wg sync.WaitGroup
	conferences := make(chan conference.Conference)

	for _, year := range l.Years {
		folder := l.BasePath + "/" + year

		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			wg.Add(1)
			go loadFile(path, conferences, &wg)
			return nil
		})
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	go func() {
		wg.Wait()
		close(conferences)
	}()

	return conferences
}

func (l *LocalFileImporter) Import() (*conference.Store, error) {
	store := &conference.Store{}
	conferences := l.LoadFiles()
	for entry := range conferences {
		store.Push(entry)
	}

	return store, nil
}
