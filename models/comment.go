package models

import "time"

type Comment struct {
	User
	time.Time
	QualityFindings
	verified bool
}

type QualityFindings struct {
	quality     uint8
	description string
}
