package service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

// InviteAPIHandler 路由分发
func InviteAPIHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateInviteHandler(w, r)
	case http.MethodGet:
		QueryInviteHandler(w, r)
	case http.MethodPut:
		UpdateInviteStatusHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// CreateInviteHandler 创建邀请
func CreateInviteHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ActivityId int32 `json:"activityId"`
		InviterId  int32 `json:"inviterId"`
		InviteeId  int32 `json:"inviteeId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	invite, err := CreateInvite(db.Get(), req.ActivityId, req.InviterId, req.InviteeId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(invite)
}

// QueryInviteHandler 查询邀请（支持按用户或活动）
func QueryInviteHandler(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("userId")
	activityIdStr := r.URL.Query().Get("activityId")
	var invites []model.InviteModel
	var err error
	if userIdStr != "" {
		userId, _ := strconv.Atoi(userIdStr)
		invites, err = ListInvitesByUser(db.Get(), int32(userId))
	} else if activityIdStr != "" {
		activityId, _ := strconv.Atoi(activityIdStr)
		invites, err = ListInvitesByActivity(db.Get(), int32(activityId))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(invites)
}

// UpdateInviteStatusHandler 更新邀请状态
func UpdateInviteStatusHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		InviteId int32  `json:"inviteId"`
		Status   string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := UpdateInviteStatus(db.Get(), req.InviteId, req.Status); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// CreateActivityHandler 创建活动
func CreateActivityHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		StartTime   string `json:"startTime"`
		EndTime     string `json:"endTime"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	start, _ := time.Parse(time.RFC3339, req.StartTime)
	end, _ := time.Parse(time.RFC3339, req.EndTime)
	activity, err := CreateActivity(db.Get(), req.Title, req.Description, start, end)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(activity)
}

// CreateUserHandler 创建用户
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OpenId    string `json:"openId"`
		Nickname  string `json:"nickname"`
		AvatarUrl string `json:"avatarUrl"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := CreateUser(db.Get(), req.OpenId, req.Nickname, req.AvatarUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
