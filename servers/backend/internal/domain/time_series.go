package domain

type TimeSeriesColumn struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	SiUnit      string `json:"si_unit"`
	StatusCode  string `json:"status_code"`
	Priority    int    `json:"priority"`
	Description string `json:"description"`
}

type TimeSeries struct {
	Columns []TimeSeriesColumn `json:"columns"`
	Data    [][]string         `json:"data"`
}
