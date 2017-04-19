package server

type Filter struct {
	Name                string  `json:"name"`
	Active              bool    `json:"active"`
	FilterTopId         string  `json:"filter-top-id"`
	FilterBottomId      string  `json:"filter-bottom-id"`
	TresholdValueTop    float64 `json:"treshold-value-top"`
	TresholdValueButtom float64 `json:"treshhold-value-bottom"`
	Unit                string  `json:"unit"`
}

type Correlator struct {
	Name            string `json:"name"`
	Timestamp       string `json:"timestamp"`
	TimestampFormat string `json:"timestamp_format"`
	MatcherId       string `json:"matcher_id"`
	MatcherValue    string `json:"matcher_value"`
	TimeTreshold    int    `json:"time_treshold"`
	Pitch           int    `json:"pitch"`
}

/************ User Config *******************/
type User struct {
	Toolbar []Items `json:"toolbar"`
	Chain   []Items `json:"chain"`
}

type Items struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}
