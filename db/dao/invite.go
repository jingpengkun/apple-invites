package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const inviteTableName = "Invites"

// InviteInterface 邀请数据模型接口
type InviteInterface interface {
	CreateInvite(invite *model.InviteModel) error
	DeleteInvite(id int32) error
	GetInvite(id int32) (*model.InviteModel, error)
	ListInvitesByActivity(activityId int32) ([]model.InviteModel, error)
	ListInvitesByInviter(inviterId int32) ([]model.InviteModel, error)
}

// InviteInterfaceImp 邀请数据模型实现
type InviteInterfaceImp struct{}

// InviteImp 实现实例
var InviteImp InviteInterface = &InviteInterfaceImp{}

// CreateInvite 创建邀请
func (imp *InviteInterfaceImp) CreateInvite(invite *model.InviteModel) error {
	cli := db.Get()
	return cli.Table(inviteTableName).Create(invite).Error
}

// DeleteInvite 删除邀请
func (imp *InviteInterfaceImp) DeleteInvite(id int32) error {
	cli := db.Get()
	return cli.Table(inviteTableName).Delete(&model.InviteModel{Id: id}).Error
}

// GetInvite 查询单个邀请
func (imp *InviteInterfaceImp) GetInvite(id int32) (*model.InviteModel, error) {
	cli := db.Get()
	invite := new(model.InviteModel)
	err := cli.Table(inviteTableName).Where("id = ?", id).First(invite).Error
	return invite, err
}

// ListInvitesByActivity 根据活动ID查询邀请列表
func (imp *InviteInterfaceImp) ListInvitesByActivity(activityId int32) ([]model.InviteModel, error) {
	cli := db.Get()
	var invites []model.InviteModel
	err := cli.Table(inviteTableName).Where("activityId = ?", activityId).Find(&invites).Error
	return invites, err
}

// ListInvitesByInviter 根据邀请人ID查询邀请列表
func (imp *InviteInterfaceImp) ListInvitesByInviter(inviterId int32) ([]model.InviteModel, error) {
	cli := db.Get()
	var invites []model.InviteModel
	err := cli.Table(inviteTableName).Where("inviterId = ?", inviterId).Find(&invites).Error
	return invites, err
}
