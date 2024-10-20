package models

import "time"

type Deployment struct {
	Id       int
	Contract string
	Group    string
	Options  []string
	Date     time.Time
}
