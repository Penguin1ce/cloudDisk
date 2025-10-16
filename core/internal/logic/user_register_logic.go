// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"cloudDisk/core/internal/svc"
	"cloudDisk/core/internal/types"
	"cloudDisk/core/models"
	"cloudDisk/core/utils"
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.CommonResponse, err error) {
	// 1. 判断code是否正确
	key := "cloudDisk:mail:" + req.Email
	code, err := l.svcCtx.Rdb.Get(l.ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		resp = &types.CommonResponse{
			Message: "验证码过期",
		}
		return resp, nil
	}
	if err != nil {
		resp = &types.CommonResponse{
			Message: "服务器错误",
		}
		return resp, nil
	}
	if code != req.Code {
		resp = &types.CommonResponse{
			Message: "验证码错误",
		}
		return resp, nil
	}
	cnt, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		resp = &types.CommonResponse{
			Message: "用户名已存在",
		}
		return resp, nil
	}
	uid := uuid.New()
	user := &models.UserBasic{
		Identity: uid.String(),
		Name:     req.Name,
		Password: utils.Md5(req.Password),
		Email:    req.Email,
	}
	row, err := l.svcCtx.Engine.InsertOne(user)
	if err != nil {
		return nil, err
	}
	log.Println("insert user row:", row)
	resp = &types.CommonResponse{
		Message: "success",
	}
	return resp, nil
}
