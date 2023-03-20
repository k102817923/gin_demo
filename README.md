gin_demo
├── conf    用于存储配置文件
│   └── app.ini
├── main.go
├── middleware  应用中间件
│   └── jwt
│       └── jwt.go
├── models  应用数据库模型
│   ├── article.go
│   ├── auth.go
│   └── models.go
│   └── tag.go
├── pkg 第三方包
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── logging
│   │   ├── file.go
│   │   └── log.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       ├── jwt.go
│       └── pagination.go
├── routers 路由逻辑处理
│   ├── api
│   │   ├── auth.go
│   │   └── v1
│   │       ├── article.go
│   │       └── tag.go
│   └── router.go
└── runtime 应用运行时数据


```shell
# 打开Go modules开关(目前在Go 1.13中默认值为auto)
go env -w GO111MODULE=on
# 设置GOPROXY代理, 这里主要涉及到两个值, 第一个是https://goproxy.cn, 它是由七牛云背书的一个强大稳定的Go模块代理, 可以有效地解决你的外网问题; 第二个是direct, 它是一个特殊的fallback选项, 它的作用是用于指示Go在拉取模块时遇到错误会回源到模块版本的源地址去抓取(比如GitHub等)
go env -w GOPROXY=https://goproxy.cn,direct
# 初始化Go modules, 它将会生成go.mod文件, 需要注意的是MODULE_PATH填写的是模块引入路径, 你可以根据自己的情况修改路径
go mod init gin_demo

go get -u github.com/gin-gonic/gin
go get -u github.com/go-ini/ini
go get -u github.com/unknwon/com
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql
go get -u github.com/astaxie/beego/validation
go get -u github.com/dgrijalva/jwt-go
```


```markdown
- 用`go get`拉取新的依赖
  - 拉取最新的版本(优先择取tag)`go get golang.org/x/text@latest`
  - 拉取master分支的最新commit`go get golang.org/x/text@master`
  - 拉取tag为v0.3.2的commit`go get golang.org/x/text@v0.3.2`
  - 拉取hash为342b231的commit, 最终会被转换为v0.3.2`go get golang.org/x/text@342b2e`
  - 用`go get -u`更新现有的依赖
  - 用`go mod download`下载go.mod文件中指明的所有依赖
  - 用`go mod tidy`整理现有的依赖
  - 用`go mod graph`查看现有的依赖结构
  - 用`go mod init`生成go.mod文件(Go 1.13中唯一一个可以生成go.mod文件的子命令)
- 用`go mod edit`编辑go.mod文件
- 用`go mod vendor`导出现有的所有依赖(事实上Go modules正在淡化Vendor的概念)
- 用`go mod verify`校验一个模块是否被篡改过
```


```markdown
go.mod
- module: 用于定义当前项目的模块路径
- go: 用于设置预期的Go版本
- require: 用于设置一个特定的模块版本
- exclude: 用于从使用中排除一个特定的模块版本
- replace: 用于将一个模块版本替换为另外一个模块版本
```


```sql
"blog 数据库"

"标签表"
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

"文章表"
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` int(11) DEFAULT NULL,
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';

"认证表"
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');
```