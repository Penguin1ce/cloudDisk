// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"cloudDisk/core/define"
	"cloudDisk/core/internal/svc"
	"cloudDisk/core/internal/types"
	"cloudDisk/core/models"
	"cloudDisk/core/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendResponse, err error) {
	// 1. 该邮箱已注册
	resp = &types.MailCodeSendResponse{}
	cnt, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		resp.Message = "服务器错误"
		return resp, err
	}
	if cnt > 0 {
		resp.Message = "该邮箱已被注册"
		return resp, nil
	}
	// 2. 如果该邮箱未注册
	code := utils.RandomCode()
	err = utils.MailSendCode(req.Email, code)
	key := "cloudDisk:mail:" + req.Email
	_, err = l.svcCtx.Rdb.Set(l.ctx, key, code, define.CodeExpire).Result()
	if err != nil {
		resp.Message = "服务器错误"
		return resp, err
	}
	resp.Message = "发送成功"
	return resp, nil
}
