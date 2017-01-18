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

/************ User Config *******************/
type User struct {
	Toolbar []Items `json:"toolbar"`
	Chain   []Items `json:"chain"`
}

type Items struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}
