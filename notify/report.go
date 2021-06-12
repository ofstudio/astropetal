package notify

import "time"

type StatusType string

const (
	StatusInfo StatusType = ""
	StatusOk   StatusType = "🌻"
	StatusErr  StatusType = "🍄"
)

type ReportItem struct {
	At      time.Time
	Status  StatusType
	Message string
}

type Report struct {
	Items []ReportItem
}

func NewReport() *Report {
	return &Report{}
}

func (r *Report) Push(status StatusType, message string) {
	r.Items = append(r.Items, ReportItem{time.Now(), status, message})
}
