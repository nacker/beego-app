# Beego Application

基于 Beego 框架开发的 Web 应用程序，实现了用户认证和管理功能。

## 项目结构

```
.
├── conf/          # 配置文件目录
├── controllers/   # 控制器目录
├── middlewares/   # 中间件目录
│   └── jwt.go     # JWT认证中间件
├── models/        # 数据模型目录
│   └── user.go    # 用户模型
├── routers/       # 路由配置目录
│   └── router.go  # 路由定义
├── services/      # 业务逻辑目录
├── static/        # 静态资源目录
├── utils/         # 工具函数目录
│   ├── error_code.go  # 错误码定义
│   ├── jwt_utils.go   # JWT工具函数
│   └── response.go    # 响应处理工具
├── views/         # 视图模板目录
└── main.go        # 应用程序入口
```

## SQL

```sql
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID，自增主键',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户名，允许为空',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '电子邮箱，允许为空',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '加密后的密码，允许为空',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号码，允许为空',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`) COMMENT '用户名唯一索引',
  UNIQUE KEY `uk_email` (`email`) COMMENT '电子邮箱唯一索引',
  UNIQUE KEY `uk_phone` (`phone`) COMMENT '手机号唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表，存储用户基本信息';
```

## 功能特性

1. **用户认证**
   - 用户注册
   - 用户登录
   - JWT token 验证

2. **数据库集成**
   - 使用 GORM 作为 ORM 框架
   - 支持自动数据库迁移
   - MySQL 数据库支持

3. **中间件支持**
   - JWT 认证中间件
   - 可扩展的中间件架构

## 技术栈

- [Beego](https://github.com/beego/beego) - Web 框架
- [GORM](https://gorm.io/) - ORM 框架
- [JWT-Go](https://github.com/golang-jwt/jwt) - JWT 实现
- MySQL - 数据库

## 快速开始

1. **创建项目**
```bash
# 创建项目（默认启用 Go Modules）
bee new myproject
cd myproject
```

2. **初始化依赖**
```bash
go mod tidy
```

3. **生成 API 文档**
```bash
# 生成并下载 Swagger 文档
bee run -gendoc=true -downdoc=true

# 生成 swagger.json
swagger generate spec -o ./swagger/swagger.json
```

## API 路由

### 认证相关
- `POST /auth/login` - 用户登录
- `POST /auth/register` - 用户注册

### 用户相关
- `GET /v1/user/:id` - 获取用户信息
- `POST /v1/user/logout` - 用户登出

所有 `/v1` 路径下的接口都需要 JWT 认证。

## 数据模型

### User 模型
```go
type User struct {
    ID         uint      `gorm:"primarykey"`
    Username   string    `gorm:"size:32;not null;unique"`
    Email      string    `gorm:"size:128;not null;unique"`
    Password   string    `gorm:"size:128;not null"`
    Phone      string    `gorm:"size:16"`
    CreateTime time.Time `gorm:"autoCreateTime"`
    UpdateTime time.Time `gorm:"autoUpdateTime"`
}
```

## 最近更新

1. 启用了数据库自动迁移功能
   - 修改了 `models/user.go` 文件
   - 添加了延迟执行逻辑，确保数据库连接初始化完成
   - 增强了错误处理机制

## 开发建议

1. **控制器完善**
   - 实现用户信息的更新功能
   - 添加密码修改接口
   - 完善用户注销功能

2. **数据验证**
   - 添加请求数据验证
   - 实现更严格的密码规则
   - 添加邮箱验证功能

3. **安全性增强**
   - 实现请求频率限制
   - 添加日志记录
   - 实现 CSRF 防护

4. **测试用例**
   - 添加单元测试
   - 编写集成测试
   - 实现 API 测试用例

## 贡献指南

1. Fork 本仓库
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开一个 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件
