package model

import "time"

// UserModel 用户模型
type UserModel struct {
	Id        int32     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OpenId    string    `gorm:"column:openId;uniqueIndex" json:"openId"` // 微信用户唯一标识
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	AvatarUrl string    `gorm:"column:avatarUrl" json:"avatarUrl"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

// ActivityModel 活动模型
type ActivityModel struct {
	Id          int32     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	StartTime   time.Time `gorm:"column:startTime" json:"startTime"`
	EndTime     time.Time `gorm:"column:endTime" json:"endTime"`
	CreatedAt   time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

// InviteModel 邀请模型
type InviteModel struct {
	Id         int32     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ActivityId int32     `gorm:"column:activityId;index" json:"activityId"`
	InviterId  int32     `gorm:"column:inviterId;index" json:"inviterId"` // 邀请人
	InviteeId  int32     `gorm:"column:inviteeId;index" json:"inviteeId"` // 被邀请人
	Status     string    `gorm:"column:status" json:"status"`             // pending/accepted/rejected
	CreatedAt  time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}
