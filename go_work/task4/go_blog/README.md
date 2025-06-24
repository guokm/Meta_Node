本项目采用 gin+gorm框架实现简单的用户注册、用户认证、文章发布、修改、查询、评论创建、查询等功能
## 安装

go get github.com/gin-gonic/gin

go get -u gorm.io/gorm

go get -u gorm.io/driver/mysql

go get github.com/golang-jwt/jwt/v5

go get github.com/spf13/viper

go get github.com/go-playground/validator/v10@v10.20.0

go get github.com/gin-gonic/gin/binding

go get github.com/zeromicro/go-zero

go get github.com/zeromicro/go-zero/core/logx


## 项目包
入口函数为main.go

config目录下加载配置以及mysql初始化
controller包下放的路由处理方法

middleware目录中放了jwt解析、权限校验、全局异常捕获的处理逻辑

model目录下存在user、post、comment的数据库实体以及出入参

router包下配置的请求地址与路由映射


