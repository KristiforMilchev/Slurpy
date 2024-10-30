package models

import "time"

type Deployment struct {
	Id       int
	Name     string
	Contract string
	Group    string
	Options  []string
	Date     time.Time
}
