# Gin&Gorm for web

### 后端培训项目
---
### 目录结构
<pre><code>
├── config 系统配置文件
├── controller 模拟控制器
├── middelware 中间件模拟鉴权
├── model 数据模型
├── router 路由模块
├── service 业务逻辑服务提供统一入口
├── vendor mod依赖
- consult.sql migrates文件
- main.go 主程序
</code></pre>

### 配置信息
---

- 配置文件: config/config.json
```
{
	"Env": "develop",//环境设置
	"Port": "8080",  //http服务监听端口
	"MySqlDsn": "root:wwww@(127.0.0.1:10110)/consult?charset=utf8&parseTime=True&loc=Local"   //mysql连接dsn
}
```

### 启动方式
---

```
1.执行sql文件
2.配置自己的config.json
3.根目录go run main.go   or  go build 
```
---
### 接口文档



**1\. 前台发布咨询**
###### 接口功能
> 用户发布咨询

###### URL
> [http://localhost:8080/publish](http://localhost:8080/publish)

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
> |参数|必选|类型|说明|
>|:-----  |:-------|:-----|-----                               |
>|content    |true    |string|咨询内容    
                      |
>
###### 返回示例(失败)              

```cassandraql
{
	"code": 0,
	"Msg": "咨询字符数必须大于10"
}
```

###### 返回示例(成功)              
```cassandraql
{
	"code": 1,
	"Msg": "发布成功，咨询id:16"
}
```   
---

**2\. 回复咨询**
###### 接口功能
> 回复咨询，须是律师用户

###### URL
> [http://localhost:8080/lawyer/reply](http://localhost:8080/lawyer/reply)

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
> |参数|必选|类型|说明|
>|:-----  |:-------|:-----|-----                               |
>|consult_id    |true    |int|对应咨询的id                          |
>|reply    |true    |string|回复的内容                          |
>|lawyer_id    |true    |int|律师id（目前仅中间件模拟 2379用户locked）                          |
>
>

###### 返回示例(失败)              
```cassandraql
{
	"code": 0,
	"Msg": "没有该条源咨询"
}
```
###### 返回示例(失败,非律师)  
```cassandraql
header:403 Forbidden
{
	"code": 0,
	"Msg": "you are not a lawyer"
}
```
###### 返回示例(成功)              
```cassandraql
{
	"code": 1,
	"Msg": {
		"Id": 12,
		"reply": "ceshidata",
		"lawyer_id": 1233,
		"created_at": "2020-09-14T17:24:14.260772+08:00"
	}
}
```
---



**3\. 前台获取咨询信息**
###### 接口功能
> 获取咨询，同时返回所有该咨询的所有回复

###### URL
> [http://localhost:8080/consult/ID](http://localhost:8080/consult/ID)

###### 支持格式
> JSON

###### HTTP请求方式
> GET

###### 请求参数
> |参数|必选|类型|说明|
>|:-----  |:-------|:-----|-----                               |
>|ID    |true    |int|咨询id                          |
>
###### 返回示例（失败）


>```cassandraql
>{
>  "code": 0,
>  "Msg": "没有查询到该咨询信息"
>}
>```

###### 返回示例（成功）
```
{
  "code": 1,
  "Msg": {
    "code": 1,
    "Msg": {
      "Id": 10,
      "content": "啊实打实的按时打算",
      "Replys": [
        {
          "Id": 1,
          "reply": "teshreply",
          "lawyer_id": 1233,
          "created_at": "2020-09-14T14:47:10+08:00"
        },
        {
          "Id": 2,
          "reply": "tesh123reply",
          "lawyer_id": 1233,
          "created_at": "2020-09-14T14:47:15+08:00"
        },
        {
          "Id": 3,
          "reply": "tesh123reply",
          "lawyer_id": 1233,
          "created_at": "2020-09-14T14:48:16+08:00"
        },
        {
          "Id": 4,
          "reply": "tesh123reply",
          "lawyer_id": 1233,
          "created_at": "2020-09-14T15:12:37+08:00"
        }
      ],
      "created_at": "2020-09-14T14:34:22+08:00"
    }
  }
}
```