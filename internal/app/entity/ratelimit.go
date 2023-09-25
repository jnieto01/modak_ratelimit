package entity



type Response struct {
	Status int	`json:"status"`
	Data ResponseData `json:"data"`
}

type ResponseData struct {

	IsAllowed bool `json:"isallowed"`
	Error DataError `json:"error"`
}


type RateLimitRule struct {
	FlowID 	string `json:"flowid"`
	Settings []RuleByType `json:"settings"`
}

type RuleByType struct {
    Key           string `json:"key"`
    MaxRequests   int    `json:"maxrequests"`
    TimeInterval  int	`json:"timeinterval"`
}


