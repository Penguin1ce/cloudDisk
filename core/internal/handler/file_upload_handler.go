// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"cloudDisk/core/models"
	"cloudDisk/core/utils"
	"crypto/md5"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"path"

	"cloudDisk/core/internal/logic"
	"cloudDisk/core/internal/svc"
	"cloudDisk/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			log.Printf("解析上传请求失败: %v", err)
			return
		}
		// 判断文件是否已存在
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			log.Printf("获取上传文件失败: %v", err)
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			log.Printf("读取上传文件失败: %v", err)
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		log.Printf("开始处理上传文件, 名称: %s, 大小: %d, Hash: %s", fileHeader.Filename, fileHeader.Size, hash)

		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			log.Printf("查询文件是否存在失败, hash: %s, 错误: %v", hash, err)
			return
		}
		if has {
			log.Printf("文件已存在, 直接返回仓库记录, hash: %s", hash)
			httpx.OkJson(w, &types.FileUploadResponse{
				Identity: hash,
				Name:     rp.Name,
				Ext:      rp.Ext,
			})
			return
		}
		// 文件不存在的时候，向腾讯云存储文件
		uploadPath, err := utils.UploadFile(r)
		if err != nil {
			log.Printf("上传文件到对象存储失败, hash: %s, 错误: %v", hash, err)
			return
		}
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = uploadPath
		log.Printf("文件上传至对象存储成功, 路径: %s", uploadPath)

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			log.Printf("写入文件信息到数据库失败, hash: %s, 错误: %v", hash, err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			log.Printf("文件上传流程完成, hash: %s, 记录ID: %s", hash, resp.Identity)
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
