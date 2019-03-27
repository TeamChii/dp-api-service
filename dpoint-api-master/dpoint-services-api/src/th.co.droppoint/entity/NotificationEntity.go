package entity

import "time"

type NotificationEntity struct {
	Noti_id            int        `gorm:"column:noti_id; primary_key" json:"noti_id"`
	Mc_id              int        `gorm:"column:mc_id" json:"mc_id"`
	Noti_category      string     `gorm:"column:noti_category" json:"noti_category"`
	Noti_subject       string     `gorm:"column:noti_subject" json:"noti_subject"`
	Noti_message       string     `gorm:"column:noti_message" json:"noti_message"`
	Noti_send_time     *time.Time `gorm:"column:noti_send_time" json:"noti_send_time"`
	Noti_send_time_str *string    `gorm:"-" json:"noti_send_time_str"`
	Noti_flag          string     `gorm:"column:noti_flag" json:"noti_flag"`
	Create_by          string     `gorm:"column:create_by" json:"create_by"`
	Create_date        *time.Time `gorm:"column:create_date" json:"create_date"`
}

func (NotificationEntity) TableName() string {
	return "dp_tb_notification"
}
