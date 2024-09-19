package domain

type Station struct {
	UUID     string `json:"uuid"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	Address  string `json:"address"`
	IsEnable bool   `json:"is_enable"`

	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Description string  `json:"description"`
}
