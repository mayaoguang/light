# Welcome To Light

## 简介

    由于微服务架构需要经常创建新服务，根据多年的golang经验，封装常用数据库组件及应用示例，方便搭建新的服务使用。
参考 [Go程序布局](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)

## 项目结构

```
├── build               // 打包/集成
|  ├── app              // 应用程序名
├── cmd                 // 可执行目录
|  ├── app              // 应用程序名
|  |  ├── main.go       // 入口文件
├── configs             // 配置文件
|  ├── config.json      
├── doc                 // 项目文档
├── example             // 示例目录
├── internal            // 私有程序
|  ├── api              // 接口
|  ├── config           // 配置文件解析
|  ├── cache            // 缓存相关
|  ├── constvar         // 常量
|  ├── domain           // 表结构
|  ├── monitor          // 监控定时服务相关
|  └── rpc              // rpc
├── logs                // 日志存放
├── pkg                 // 安全导入的包(可以被任何项目直接导入使用)
|  ├── clickhouse       // ck组件
|  ├── email            // 邮件组件
|  ├── es               // es组件
|  ├── httpcode         // 请求处理组件
|  ├── jwt              // jwt组件
|  ├── libs             // 封装的公用方法
|  ├── logging          // 日志组件
|  ├── mongo            // mongo组件
|  ├── mq               // mq组件
|  ├── mysql            // mysql组件
|  ├── redis            // redis组件
|  ├── safego           // 安全运行组件
|  └── ws               // socket组件 
├── utils               // 自己封装的通用方法
├── .gitignore          // git忽略文件    
├── go.mod              // 包管理    
├── README.md
```
