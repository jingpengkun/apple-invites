package service

import (
	"time"
	"wxcloudrun-golang/db/model"

	"gorm.io/gorm"
)

// CreateUser 创建用户
func CreateUser(db *gorm.DB, openId, nickname, avatarUrl string) (*model.UserModel, error) {
	user := &model.UserModel{
		OpenId:    openId,
		Nickname:  nickname,
		AvatarUrl: avatarUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByOpenId 根据 openId 查询用户
func GetUserByOpenId(db *gorm.DB, openId string) (*model.UserModel, error) {
	var user model.UserModel
	if err := db.Where("openId = ?", openId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateActivity 创建活动
func CreateActivity(db *gorm.DB, title, description string, startTime, endTime time.Time) (*model.ActivityModel, error) {
	activity := &model.ActivityModel{
		Title:       title,
		Description: description,
		StartTime:   startTime,
		EndTime:     endTime,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := db.Create(activity).Error; err != nil {
		return nil, err
	}
	return activity, nil
}

// GetActivityById 查询活动
func GetActivityById(db *gorm.DB, id int32) (*model.ActivityModel, error) {
	var activity model.ActivityModel
	if err := db.Where("id = ?", id).First(&activity).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

// CreateInvite 创建邀请
func CreateInvite(db *gorm.DB, activityId, inviterId, inviteeId int32) (*model.InviteModel, error) {
	invite := &model.InviteModel{
		ActivityId: activityId,
		InviterId:  inviterId,
		InviteeId:  inviteeId,
		Status:     "pending",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := db.Create(invite).Error; err != nil {
		return nil, err
	}
	return invite, nil
}

// UpdateInviteStatus 更新邀请状态
func UpdateInviteStatus(db *gorm.DB, inviteId int32, status string) error {
	return db.Model(&model.InviteModel{}).Where("id = ?", inviteId).Update("status", status).Error
}

// ListInvitesByUser 查询用户收到的邀请
func ListInvitesByUser(db *gorm.DB, userId int32) ([]model.InviteModel, error) {
	var invites []model.InviteModel
	if err := db.Where("inviteeId = ?", userId).Find(&invites).Error; err != nil {
		return nil, err
	}
	return invites, nil
}

// ListInvitesByActivity 查询活动下所有邀请
func ListInvitesByActivity(db *gorm.DB, activityId int32) ([]model.InviteModel, error) {
	var invites []model.InviteModel
	if err := db.Where("activityId = ?", activityId).Find(&invites).Error; err != nil {
		return nil, err
	}
	return invites, nil
}
