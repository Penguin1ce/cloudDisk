// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"cloudDisk/core/internal/svc"
	"cloudDisk/core/internal/types"
	"cloudDisk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileRenameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileRenameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileRenameLogic {
	return &UserFileRenameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileRenameLogic) UserFileRename(req *types.UserFileRenameRequest, userIdentity string) (resp *types.CommonResponse, err error) {
	data := &models.UserRepository{
		Name: req.Name,
	}
	cnt, err := l.svcCtx.Engine.
		Where("name = ? AND parent_id = (select parent_id from user_repository ur where ur.identity = ?)", req.Name, req.Identity).
		Count(&models.UserRepository{})
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		resp = &types.CommonResponse{
			Message: "已经存在相同的文件名",
		}
		return resp, errors.New("存在相同的文件名")
	}
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).
		Update(data)
	if err != nil {
		resp = &types.CommonResponse{
			Message: "fail",
		}
		return resp, err
	}
	resp = &types.CommonResponse{
		Message: "success" + req.Name,
	}
	return
}
