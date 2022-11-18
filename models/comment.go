package models

import "time"

type Comment struct {
	User
	time.Time
	content string
	QualityFindings
	verified bool
}

type QualityFindings struct {
	quality     uint8
	description string
}

func NewComment(user User, content string) Comment {
	return Comment{
		user,
		time.Now(),
		content,
		QualityFindings{
			10,
			"null",
		},
		false,
	}
}

func (c Comment) GetContent() string {
	return c.content
}

func (c Comment) IsVerified() bool {
	return c.verified
}
