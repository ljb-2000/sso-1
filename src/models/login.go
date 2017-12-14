package models

import (
	"github.com/pborman/uuid"
	"time"
)

const (
	DEFAULT_EXPIRATION = "86400" // 24*60*60s
)

type Session struct {
	SessionId  string   `json:"sessionId"`
	UserId     int64	`json:"userId"`
	UserName   string   `json:"userName"`
	CreatedAt  string   `json:"createdAt"`
	TimeStamp  string	`json:"timeStamp"`
	Expiration string   `json:"expiration"`
}

func NewSession(userId int64, userName string) *Session {
	return &Session{
		SessionId:  uuid.New(),
		UserId:     userId,
		UserName:   userName,
		CreatedAt:  time.Now().String(),
		TimeStamp:	time.Now().String(),
		Expiration: DEFAULT_EXPIRATION,
	}
}

