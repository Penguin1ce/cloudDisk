CloudDisk Core 服务
==================

项目介绍
--------
CloudDisk Core 是云盘系统的核心后端服务，基于 Go 语言与 go-zero 框架构建，负责用户认证、文件上传以及个人文件仓库的管理。服务通过 RESTful API 对外提供能力，并依赖 MySQL 与 Redis 作为数据持久化与缓存组件。

功能特性
--------
- 用户注册、登录与 Token 鉴权
- 邮箱验证码发送与校验
- 文件秒传检测与上传记录管理
- 用户文件仓库目录管理

技术栈
------
- 语言：Go 1.20+
- 框架：go-zero (rest)
- 数据库：MySQL
- 缓存：Redis

快速开始
--------
1. 准备数据库与缓存：确保 MySQL 与 Redis 可访问，并在 `etc/core-api.yaml` 中配置连接信息。
2. 启动服务：
   ```bash
   go run core.go -f etc/core-api.yaml
   ```
3. 服务默认监听 `0.0.0.0:8080`，可通过 cURL/Postman 调用下述接口。

接口说明
--------

### 公共接口

#### 用户登录 `/user/login`
- 方法：`POST`
- 描述：校验用户账号密码并返回访问 Token。
- 请求体：

  | 字段 | 类型 | 必填 | 说明 |
  | --- | --- | --- | --- |
  | `name` | string | 是 | 用户名 |
  | `password` | string | 是 | 用户密码 |

- 响应体：

  | 字段 | 类型 | 说明 |
  | --- | --- | --- |
  | `token` | string | 鉴权 Token，后续受保护接口在 Header `Authorization: Bearer <token>` 中携带 |

#### 用户详情 `/user/detail`
- 方法：`GET`
- 描述：根据用户标识查询账号详情。
- 请求参数（Query）：

  | 字段 | 类型 | 必填 | 说明 |
  | --- | --- | --- | --- |
  | `identity` | string | 是 | 用户唯一标识 |

- 响应体：

  | 字段 | 类型 | 说明 |
  | --- | --- | --- |
  | `name` | string | 用户昵称 |
  | `email` | string | 用户邮箱 |

#### 邮箱验证码发送 `/user/code/send/register`
- 方法：`POST`
- 描述：发送注册验证码至邮箱。
- 请求体：

  | 字段 | 类型 | 必填 | 说明 |
  | --- | --- | --- | --- |
  | `email` | string | 是 | 收件邮箱地址 |

- 响应体：

  | 字段 | 类型 | 说明 |
  | --- | --- | --- |
  | `message` | string | 操作结果描述 |

#### 用户注册 `/user/register`
- 方法：`POST`
- 描述：创建新用户账号。
- 请求体：

  | 字段 | 类型 | 必填 | 说明 |
  | --- | --- | --- | --- |
  | `name` | string | 是 | 用户名 |
  | `password` | string | 是 | 登录密码 |
  | `email` | string | 是 | 邮箱地址 |
  | `code` | string | 是 | 邮箱验证码 |

- 响应体：

  | 字段 | 类型 | 说明 |
  | --- | --- | --- |
  | `message` | string | 操作结果描述 |

### 需鉴权接口
以下接口需在请求 Header 中添加 `Authorization: Bearer <token>`。

#### 文件上传 `/file/upload`
- 方法：`POST`
- 描述：上传文件或触发秒传校验，返回文件标识。
- 请求体：

  | 字段 | 类型 | 必填 | 说明 |
  | --- | --- | --- | --- |
  | `hash` | string | 否 | 文件哈希，用于秒传校验 |
  | `name` | string | 否 | 原始文件名 |
  | `ext` | string | 否 | 文件扩展名 |
  | `size` | int64 | 否 | 文件大小（字节） |
  | `path` | string | 否 | 文件存储路径或临时地址 |

- 响应体：

  | 字段 | 类型 | 说明 |
  | --- | --- | --- |
  | `identity` | string | 文件唯一标识 |
  | `name` | string | 文件名 |
  | `ext` | string | 文件扩展名 |

#### 保存文件到用户仓库 `/file/repository/save`
- 方法：`POST`
- 描述：将文件记录写入用户文件仓库，实现个人目录管理。
- 请求体：

  | 字段 | 类型 | 必填 | 说明 |
  | --- | --- | --- | --- |
  | `parentId` | int64 | 是 | 目标目录的父节点 ID，根目录填 0 |
  | `repositoryIdentity` | string | 是 | 文件库标识 |
  | `name` | string | 是 | 文件或目录名称 |
  | `ext` | string | 否 | 文件扩展名（目录为空） |

- 响应体：空对象，若成功返回 HTTP 200。

维护建议
--------
- 更新接口后同步调整本文档，确保请求参数与响应字段保持一致。
- 结合 goctl 生成的 `core.api` 文件进行比对，可减少遗漏。
