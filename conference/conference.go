package conference

import "time"

type Conference struct {
	Name       string    `json:"name"`
	URL        string    `json:"url"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	City       string    `json:"city"`
	Country    string    `json:"country"`
	CfpUrl     string    `json:"cfpUrl"`
	CfpEndDate time.Time `json:"cfpEndDate"`
	Twitter    string    `json:"twitter"`
	Category   string    `json:"category"`
}
