package conference

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var conferences = []Conference{
	Conference{
		Name:       "CapitalGo",
		URL:        "http://capitalgolang.com",
		StartDate:  time.Now().AddDate(2020, 01, 15),
		EndDate:    time.Now().AddDate(2020, 01, 17),
		City:       "Washington, DC",
		Country:    "U.S.A.",
		CfpUrl:     "https://www.papercall.io/capitalgo2019",
		CfpEndDate: time.Now().AddDate(2019, 10, 01),
		Twitter:    "@CapitalGolang",
		Category:   "golang",
	},
	Conference{
		Name:       "HalfStack Conf",
		URL:        "https://halfstackconf.com/phoenix/",
		StartDate:  time.Now().AddDate(2020, 02, 07),
		EndDate:    time.Now().AddDate(2020, 02, 07),
		City:       "Phoenix",
		Country:    "U.S.A.",
		CfpUrl:     "https://halfstackconf.com/phoenix/",
		CfpEndDate: time.Now().AddDate(2020, 01, 01),
		Twitter:    "@halfstackconf",
		Category:   "javascript",
	},
	Conference{
		Name:       "Agent Conf",
		URL:        "https://www.agent.sh",
		StartDate:  time.Now().AddDate(2020, 06, 01),
		EndDate:    time.Now().AddDate(2020, 06, 05),
		City:       "Dornbirn & Lech",
		Country:    "Austria",
		CfpUrl:     "",
		CfpEndDate: time.Now().AddDate(2020, 01, 01),
		Twitter:    "@AgentConf",
		Category:   "javascript",
	},
}

func TestStringExists(t *testing.T) {
	var result bool

	slice := []string{"lorem", "ipsum", "dolor"}
	emptySlice := []string{}

	result = stringExists(slice, "ipsum")
	assert.True(t, result)

	result = stringExists(slice, "amet")
	assert.False(t, result)

	result = stringExists(emptySlice, "amet")
	assert.False(t, result)

	result = stringExists(slice, "")
	assert.False(t, result)
}

func TestPush(t *testing.T) {
	s := &Store{}

	assert.Len(t, s.Entries, 0)
	assert.Len(t, s.Cities, 0)
	assert.Len(t, s.Countries, 0)
	assert.Len(t, s.Categories, 0)

	for _, conference := range conferences {
		s.Push(conference)
	}

	assert.Len(t, s.Entries, 3)
	assert.Len(t, s.Cities, 3)
	assert.Len(t, s.Countries, 2)
	assert.Len(t, s.Categories, 2)
}

func TestFilterByCity(t *testing.T) {
	s := &Store{}

	for _, conference := range conferences {
		s.Push(conference)
	}

	paris := s.FilterByCity("paris")
	phoenix := s.FilterByCity("phoenix")

	assert.Len(t, paris.Entries, 0)
	assert.Len(t, phoenix.Entries, 1)
}

func TestFilterByCountry(t *testing.T) {
	s := &Store{}

	for _, conference := range conferences {
		s.Push(conference)
	}

	france := s.FilterByCountry("france")
	usa := s.FilterByCountry("U.S.A.")

	assert.Len(t, france.Entries, 0)
	assert.Len(t, usa.Entries, 2)
}

func TestFilterByCategory(t *testing.T) {
	s := &Store{}

	for _, conference := range conferences {
		s.Push(conference)
	}

	javascript := s.FilterByCategory("JavaSCRipt")
	golang := s.FilterByCategory("GolAnG")
	java := s.FilterByCategory("java")

	assert.Len(t, javascript.Entries, 2)
	assert.Len(t, golang.Entries, 1)
	assert.Len(t, java.Entries, 0)
}

func TestFilterByName(t *testing.T) {
	s := &Store{}

	for _, conference := range conferences {
		s.Push(conference)
	}

	conf := s.FilterByName("conf", false)
	agent := s.FilterByName("agent", false)
	lorem := s.FilterByName("lorem", false)
	halfExact := s.FilterByName("HalfStack conf", true)

	assert.Len(t, conf.Entries, 2)
	assert.Len(t, agent.Entries, 1)
	assert.Len(t, lorem.Entries, 0)
	assert.Len(t, halfExact.Entries, 1)
}

func TestFilterByDateRange(t *testing.T) {
	s := &Store{}

	for _, conference := range conferences {
		s.Push(conference)
	}

	case_1 := s.FilterByDateRange(time.Now().AddDate(2020, 6, 1), time.Now().AddDate(2020, 6, 5), 0)
	case_1_b := s.FilterByDateRange(time.Now().AddDate(2020, 6, 6), time.Now().AddDate(2020, 6, 9), 0)
	case_1_c := s.FilterByDateRange(time.Now().AddDate(2020, 6, 6), time.Now().AddDate(2020, 6, 9), 10)

	case_2 := s.FilterByDateRange(time.Now().AddDate(2020, 10, 10), time.Now().AddDate(2020, 10, 12), 0)

	assert.Len(t, case_1.Entries, 1)
	assert.Len(t, case_1_b.Entries, 0)
	assert.Len(t, case_1_c.Entries, 1)

	assert.Len(t, case_2.Entries, 0)
}
