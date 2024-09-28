package dao

import "gorm.io/gorm"

type MsgStatus int

const (
	Generated MsgStatus = iota
	Pushed
	Consumed
)

type Message struct {
	gorm.Model
	Status     MsgStatus `gorm:"status;index"`
	MetaData   string    `gorm:"meta_data;"`
	CampaignID string    `gorm:"campaign_id;index"`
}

func (m *Message) TableName() string {
	return "tb_message"
}
