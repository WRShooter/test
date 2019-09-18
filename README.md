# 学生证书管理系统
用户管理学生的学生证书,写的一些api
## 技术选型
_________
    1.web:beego<br>
    2.orm:gorm<br>
    3.database:mysql
## 项目结构
_________
    -Student management system
        |-conf 配置文件目录
        |-config 初始化代码
        |-controller 控制器目录
        |-method 复用代码
        |-models 模块目录
        |-routers 路由目录
        |-tls-gen ssl证书
## 安装部署
_________
    git clone https://github.com/WRShooter/test
    cd test
    bee run -gendoc=true -downdoc=true
    
