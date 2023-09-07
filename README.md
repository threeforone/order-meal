## 1.描述 description
本项目是一个简单的增删查改的go项目，可以在这里看到gin框架的中间件，jwt，gorm如何使用。初学者练手使用。主要功能进行一个简单的统计，统计公司给员工订加班晚餐的情况。

This is a simple CRUD golang project for practing by the beginner.You can see how to use gin,jwt,gorm in this project.The main goal is to 
record the statistics the overtime dinner the company order for employees.

## 2.环境 environment
MySQL: 8.0.20

Go: 1.20.1

gin:  v1.9.1

gorm: v1.25.4

## 3.运行 run
建表语句在server/sql/meal.sql

运行前需要将 server/conf/config.yaml 的mysql.dsn 替换成自己的地址 `name:password@tcp(ip:port)/db?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s&readTimeout=30s&writeTimeout=20s`
示例
```
root:root123@tcp(127.0.0.1:3306)/meal?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s&readTimeout=30s&writeTimeout=20s
```

The path for the DDL statements is located at server/sql/meal.sql. Before you run ,you should replace the address you owned in the server/conf/config.ymal mysql.dsn.

For example
```
root:root123@tcp(127.0.0.1:3306)/meal?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s&readTimeout=30s&writeTimeout=20s
```
## 4.展示 display
![image](https://github.com/threeforone/order-meal/server/images/home.png)
![image](https://github.com/threeforone/order-meal/server/images/day.png)
![image](https://github.com/threeforone/order-meal/server/images/history.png)
![image](https://github.com/threeforone/order-meal/server/images/users.png)
