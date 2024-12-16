package models

type DateTimeResponse struct {
	Routes   string `json:"routes"`
	Datetime string `json:"datetime"`
}

type TimeActionSend struct {
	AreaId    string `json:"area_id"`
	City      string `json:"city"`
	Continent string `json:"continent"`
	Hour      int    `json:"hour"`
	Minute    int    `json:"minute"`
}

type TimeActionDatabase struct {
	TimeActionSend
	Id         int    `json:"id"`
}
