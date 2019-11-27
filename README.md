# beemongo
基于beego框架，mongodb数据库的简单golang服务器框架
#### 已集成特性
1. beego 一个使用 Go 的思维来帮助您构建并开发 Go 应用程序的开源框架<br>
https://beego.me/
2. mongo-go-driver mongodb官方go驱动<br>
https://github.com/mongodb/mongo-go-driver
3. casbin 访问控制<br>
https://casbin.org/en/
4. beego-pongo2 go服务端渲染用模板语言
https://github.com/oal/beego-pongo2
#### demo中有哪些实现？
1. 通过扫描domains目录自动创建表和索引
2. mongodb数据库的crud
3. mongodb事务
4. mongodb视图
5. mongodb复杂关联分页查询
6. beego统一异常处理
7. beego拦截器，访问权限控制
8. 使用pongo2做页面渲染
#### 待处理的问题
1. golang的版本管理工具，实在是难用，老出问题，导致git上没有使用包的当前版本信息
2. 考虑怎么接入微服务架构
#### 解决部分问题
1. 依赖管理直接使用 ```go mod```
