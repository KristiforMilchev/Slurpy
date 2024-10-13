package models

import "time"

type Deployment struct {
	Id       int
	Contract string
	Options  []string
	Date     time.Time
}
