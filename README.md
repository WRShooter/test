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
### Student:<br>
1):GET:curl -i -X GET localhost:8080/student<br>
2):GET:curl -i -X GET localhost:8080/student/id <br>
3):POST:curl -i -X POST localhost:8080/student?id=1&pwd=1&name=1&idcard=1&dept=1&grade=1&class=1&professional=1<br>
4):PUT:curl -i -X PUT localhost:8080/student/id?pwd=1&name=1&idcard=1&dept=1&grade=1&class=1&professional=1<br>
5):DELETE:curl -i -X DELETE localhost:8080/student/id<br>

### Stuorder:<br>
1):GET:curl -i -X GET localhost:8080/stuorder<br>
2):GET:curl -i -X GET localhost:8080/stuorder/wid <br>
3):POST:curl -i -X POST localhost:8080/stuorder -d '{"W_id": 1,"Question": 1, "W_flag": 1, "Stu_id": 1}'
4):PUT:curl -i -X PUT localhost:8080/stuorder/wid?w_flag=1<br>(仅可以修改处理状况)
5):DELETE:curl -i -X DELETE localhost:8080/stuorder/wid<br>

### Staff:<br>
1):GET:curl -i -X GET localhost:8080/staff<br>
2):GET:curl -i -X GET localhost:8080/staff/id <br>
3):POST:curl -i -X POST localhost:8080/staff -d '{"Id": 1,"Pwd": 1, "Dept": 1, "Cap": 1, "Name": 1, "Position": 1, "Phone": 1, "Idcard": 1}'
4):PUT:curl -i -X PUT localhost:8080/staff/id?pwd=1&name=1&idcard=1&dept=1&cap=1&position=1&phone=1<br>
5):DELETE:curl -i -X DELETE localhost:8080/staff/id<br>

### Manager:<br>
1):GET:curl -i -X GET localhost:8080/manager<br>
2):GET:curl -i -X GET localhost:8080/manager/id <br>
3):POST:curl -i -X POST localhost:8080/manager -d '{"Id": 1,"Pwd": 1, "Department": 1, "Cap": 1, "Name": 1, "Phone": 1, "Idcard": 1}'
4):PUT:curl -i -X PUT localhost:8080/manager/id?pwd=1&name=1&idcard=1&dept=1&cap=1&position=1&phone=1=1<br>
5):DELETE:curl -i -X DELETE localhost:8080/manager/wid<br>

### Cert_Cate:<br>
1):GET:curl -i -X GET localhost:8080/cert_cate<br>
2):GET:curl -i -X GET localhost:8080/cert_cate/id <br>
3):POST:curl -i -X POST localhost:8080/cert_cate -d '{"Id": 1,"Question": 1, "Name": 1}'
4):PUT:curl -i -X PUT localhost:8080/cert_cate/id?name=1<br>
5):DELETE:curl -i -X DELETE localhost:8080/cert_cate/id<br>

### Cert_Cate_List:<br>
1):GET:curl -i -X GET localhost:8080/cert_cate_list<br>
2):GET:curl -i -X GET localhost:8080/cert_cate_list/id <br>
3):POST:curl -i -X POST localhost:8080/cert_cate_list -d '{"Id": 1,"Name": 1, "Abbreviation": 1, "Cate_id": 1, "Cate_list_level": 1, "Institution": 1}'
4):PUT:curl -i -X PUT localhost:8080/cert_cate_list/id?name=1&abbreviation=1&cate_id=1&cate_list_level=1&institution=1<br>
5):DELETE:curl -i -X DELETE localhost:8080/cert_cate_list/id<br>

### Cert_Save:<br>
1):GET:curl -i -X GET localhost:8080/cert_save<br>
2):GET:curl -i -X GET localhost:8080/cert_save/id <br>
3):POST:curl -i -X POST localhost:8080/cert_save -d '{"Id": 1,"Cert_id": 1, "Stu_id": 1, "Num": 1, "Time": 1, "Hash": 1}'
4):PUT:curl -i -X PUT localhost:8080/cert_save/id?cert_id=1&num=1&stu_id=1&time=1&hash=1<br>
5):DELETE:curl -i -X DELETE localhost:8080/cert_save/id<br>
    
