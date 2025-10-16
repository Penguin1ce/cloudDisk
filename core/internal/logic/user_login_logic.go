// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"cloudDisk/core/models"
	"cloudDisk/core/utils"
	"context"
	"errors"

	"cloudDisk/core/internal/svc"
	"cloudDisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 1. 从数据库中查询当前用户
	user := new(models.UserBasic)
	get, err := l.svcCtx.Engine.Where("name = ? AND password = ?", req.Name, utils.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("用户名或者密码错误")
	}
	// 2. 生成token
	token, err := utils.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	resp = &types.LoginResponse{}
	resp.Token = token
	return resp, nil
}
