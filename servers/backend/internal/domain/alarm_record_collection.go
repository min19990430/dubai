package domain

import "time"

type AlarmRecordCollection struct {
	Name      string    `json:"name"`
	FullName  string    `json:"full_name"`
	OccurTime time.Time `json:"occur_time"`
	Content   string    `json:"content"`
}
