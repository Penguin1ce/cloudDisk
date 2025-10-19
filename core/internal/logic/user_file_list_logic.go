// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"cloudDisk/core/define"
	"cloudDisk/core/models"
	"context"

	"cloudDisk/core/internal/svc"
	"cloudDisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListResponse, err error) {
	var uf []*types.UserFile
	// 分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	// 分页查询用户文件数据
	err = l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Limit(size, offset).Find(&uf)
	if err != nil {
		return nil, err
	}
	// 查询用户文件总数
	count, err := l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Count(&models.UserRepository{})
	if err != nil {
		return nil, err
	}
	logx.Info("查询成功")
	resp = &types.UserFileListResponse{
		List:  uf,
		Count: count,
	}
	return
}
