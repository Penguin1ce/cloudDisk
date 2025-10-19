// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"cloudDisk/core/internal/svc"
	"cloudDisk/core/internal/types"
	"cloudDisk/core/models"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateResponse, err error) {
	cnt, err := l.svcCtx.Engine.
		Where("name = ? AND parent_id = ?", req.Name, userIdentity).
		Count(&models.UserRepository{})
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("存在相同的文件名")
	}
	data := &models.UserRepository{
		Identity:     uuid.New().String(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return nil, err
	}
	resp = &types.UserFolderCreateResponse{
		Identity: data.Identity,
	}
	return
}
