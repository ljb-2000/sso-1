package models

import (
	"time"

	"github.com/pborman/uuid"
)

// 超时时间
const (
	DEFAULT_EXPIRATION = "86400" // 24*60*60s
)

// session结构体
type Session struct {
	SessionId  string   `json:"sessionId"`
	UserId     int64	`json:"userId"`
	UserName   string   `json:"userName"`
	CreatedAt  string   `json:"createdAt"`
	Expiration string   `json:"expiration"`
}

// 新建session
func NewSession(userId int64, userName string) *Session {
	return &Session{
		SessionId:  uuid.New(),
		UserId:     userId,
		UserName:   userName,
		CreatedAt:  time.Now().String(),
		Expiration: DEFAULT_EXPIRATION,
	}
}

