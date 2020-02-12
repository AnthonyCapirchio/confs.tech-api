package importer

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"squidward.confs.tech/conference"
)

func TestCategoryFromPath(t *testing.T) {
	assert.Equal(t, CategoryFromPath("/path/to/heaven/java.json"), "java")
	assert.Equal(t, CategoryFromPath("/path/to/heaven/java"), "java")
	assert.Equal(t, CategoryFromPath("/path/to/heaven/java.yaml"), "java.yaml")
	assert.Equal(t, CategoryFromPath("golang.json"), "golang")
	assert.Equal(t, CategoryFromPath(""), "")
}

func TestHandleFileContent(t *testing.T) {

	var wg sync.WaitGroup
	data := `[
		{
		  "name": "HalfStack",
		  "url": "https://halfstackconf.com/phoenix/",
		  "startDate": "2020-01-17T00:00:00Z",
		  "endDate": "2020-01-17T00:00:00Z",
		  "city": "Phoenix",
		  "country": "U.S.A.",
		  "cfpUrl": "https://halfstackconf.com/phoenix/",
		  "cfpEndDate": "2019-09-15T00:00:00Z",
		  "twitter": "@halfstackconf"
		},
		{
		  "name": "Agent Conf",
		  "url": "https://www.agent.sh",
		  "startDate": "2020-01-23T00:00:00Z",
		  "endDate": "2020-01-26T00:00:00Z",
		  "city": "Dornbirn & Lech",
		  "country": "Austria",
		  "cfpUrl": "",
		  "cfpEndDate": "0001-01-01T00:00:00Z",
		  "twitter": "@AgentConf"
		},
		{
		  "name": "Covalence",
		  "url": "http://www.covalenceconf.com",
		  "startDate": "2020-01-24T00:00:00Z",
		  "endDate": "2020-01-24T00:00:00Z",
		  "city": "San Francisco",
		  "country": "U.S.A.",
		  "cfpUrl": "https://www.papercall.io/covalence2020",
		  "cfpEndDate": "2019-10-31T00:00:00Z",
		  "twitter": "@CovalenceConf"
		}
	  ]`

	conferences := make(chan conference.Conference)
	store := &conference.Store{}
	wg.Add(3)

	go HandleFileContent([]byte(data), "/path/to/heaven/javascript.json", conferences)

	go func() {
		wg.Wait()
		close(conferences)
	}()

	for entry := range conferences {
		wg.Done()
		store.Push(entry)
	}

	assert.Len(t, store.Entries, 3)
	assert.Len(t, store.Categories, 1)
	assert.Equal(t, store.Entries[0].Name, "HalfStack")
}
