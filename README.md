# gin-bbs
as a bbs program of golang

##### [示例地址](http://localhost:8080)

### 初衷
....

### 技术选型
1. web:gin
2. orm:gorm
3. database:sqlite3
4. ....

### 项目结构
```shell
-gin-bbs
|-controllers 控制器目录
|-models 数据库访问目录
|-static 静态资源目录
|-go.mod 模块文件
|-main.go 程序执行入口
|-...
```

### TODO
- [ ] 系统日志
- [ ] 网站统计 
- [ ] 文章、页面访问统计
- [ ] github登录发表评论 
- [ ] rss 
- [ ] 定时备份系统数据 
- [ ] 邮箱订阅功能

### 安装部署
本项目使用go mod管理依赖包，go mod安装方法
```shell
go mod download
```

```shell
go get -u github.com/daryeon/gin-bbs
```

```shell
git clone https://github.com/daryeon/gin-bbs.git
cd gin-bbs
go mod download
go run main.go
```

### 使用方法
#### 使用说明
1. 修改....
2. 访问http://xxx.xxx/
3. 修改....