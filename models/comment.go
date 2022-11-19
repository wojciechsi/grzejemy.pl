package models

import "time"

type Comment struct {
	User
	time.Time
	Content string
	QualityFindings
	Verified bool
}

type QualityFindings struct {
	Quality     uint8
	Description string
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
	return c.Content
}

func (c Comment) IsVerified() bool {
	return c.Verified
}
