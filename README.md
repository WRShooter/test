# 学生证书管理系统
用户管理学生的学生证书,写的一些api
## 技术选型
_________
    1.gin:轻量级Web框架<br>
    2.gorm:orm,配合mysql使用<br>
    3.mysql:数据库
## 项目结构
_________
    -Student management system
        |-conf 配置文件目录
        |-config 初始化代码
        |-controller 控制器目录
        |-method 复用代码
        |-models 模块目录
        |-tls-gen ssl证书
## 安装部署
_________
    git clone https://github.com/WRShooter/test
    cd test
    go run main.go

## 功能
_________
    1.创建了用户模型<br>
    2.实现了基本的curd接口

## API
_________
### 1.Student:<br>
1):GET:curl -i -X GET localhost:8080/student<br>
2):GET:curl -i -X GET localhost:8080/student/id <br>
3):POST:curl -i -X POST localhost:8080/student?id=1&pwd=1&name=1&idcard=1&dept=1&grade=1&class=1&professional=1<br>
4):PUT:curl -i -X PUT localhost:8080/student/id?pwd=1&name=1&idcard=1&dept=1&grade=1&class=1&professional=1<br>
5):DELETE:curl -i -X DELETE localhost:8080/student/id<br>
    
