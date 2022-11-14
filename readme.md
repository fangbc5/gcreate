# gcreate - go语言实现的代码生成器

## 目录结构

```
|-- src
    |-- go.mod
    |-- go.sum
    |-- main.go					//主程序
    |-- readme.md				//文档
    |-- conf
    |   |-- app.yaml				//配置文件
    |   |-- viper.go				//配置加载器
    |-- db
    |   |-- info.go				//数据库表、字段类型封装
    |   |-- mysql.go			        //用于查询mysql数据库表字段信息
    |-- handler
    |   |-- panic_handler.go			//使用类似java的Try...Catch...处理异常
    |-- interface
    |   |-- configloader
    |   |   |-- loader.go			//配置加载器接口：viper
    |   |-- datasource
    |   |   |-- loader.go			//数据源接口支持：mysql
    |   |-- tmpl
    |       |-- engine.go			//模版引擎接口
    |-- templates
        |-- .DS_Store
        |-- pongo.go				//模版引擎使用pongo2
        |-- v1					//gin+gorm应用模版
        |   |-- app
        |       |-- router
        |       |-- api
        |       |   |-- api
        |       |-- dao
        |       |   |-- dao
        |       |-- model
        |       |   |-- model
        |       |-- service
        |           |-- service
        |-- v2					//测试生成目录，可在conf/app.yaml中修改生成目录
     
```

## 教程

源代码执行main方法，会在templates/v2中生成代码，将生成的代码copy到gin-gorm项目中

## 配置

文件名：app.yaml

加载路径：

1、环境变量GCREATE_CONFIG_DIR配置的路径

2、相对于可执行文件目录下app.yaml

3、相对于可执行文件目录下conf/app.yaml

4、相对于可执行文件目录下../conf/app.yaml

app.yaml内容如下

```
project:            	      #涉及代码中依赖包的路径{project.name}/app/{project.module}
  name: halo        	      #项目名称
  module: blog      	      #模块名称        
table:
  names: t_user     	      #生成表的名称多个用逗号隔开如t_user,t_role
  prefix: t_        	      #表名称前缀，实体类生成会将表名前缀去掉
mysql:              	      #数据库连接信息
  username: root              #用户名
  password: 123456            #密码
  driver: mysql               #驱动
  url: tcp(127.0.0.1:3306)    #url
  schema: go_test             #数据库
dir:
  tmpl: templates/v1          #模版路径，受环境变量GCREATE_TMPL_PATH控制
  out: templates/v2           #代码生成路径，受环境变量GCREATE_OUT_PATH控制
```

## 发行

1、go build/go install 打包可执行文件

2、创建配置文件

3、配置模版路径和代码生成路径

4、模版路径下创建模版