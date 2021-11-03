

# gin_devstructure









**基于go语言Gin框架的web开发脚手架**

## 架构分层

-   前端(vue/react)
-   Nginx/Apache
-   Controller(服务入口，路由处理，参数校验，请求转发)
-   Logic/Service(逻辑层(服务层)， 负责业务处理逻辑)
-   DAO.Reposttory(数据库存储相关实现)










## 项目总体模块分布
-   setting模块
-   controller模块
-   dao/mysql模块
-   dao/redis模块
-   logger模块
-   models模块
-   pkg模块
-   routes模块
-   service模块
-   main.go





## 基于登录注册的脚手架








### Git推送配置
`git config --global http.postBuffer 524288000`