package model

import "encoding/json"

type Profile struct {
	Name       string `json:"name" xml:"name"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	Height     int    `json:"height"`
	Weight     int    `json:"weight"`
	Income     string `json:"income"`
	Marriage   string `json:"marriage"`
	Education  string `json:"education"`
	Occupation string `json:"occupation"`
	HoKou      string `json:"ho_kou"`
	Xinzuo     string `json:"xinzuo"`
	House      string `json:"house"`
	Car        string `json:"car"`
}

func FromJson(o interface{}) Profile {
	var p Profile
	ret, _ := json.Marshal(o)
	json.Unmarshal(ret, &p)
	return  p
}

type Item struct {
	Url string `json:"url"`
	Id string `json:"id"`
	Type string `json:"type"`
	Payload interface{}
}
