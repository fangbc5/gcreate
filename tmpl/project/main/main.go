package main

import (
	"{{Code}}/handler"
	"log"

	database "github.com/fangbc5/gogo/core/database/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据库连接
	database.Init(
		database.WithAddress("43.143.136.7"),
		database.WithPort("3306"),
		database.WithUsername("root"),
		database.WithPassword("1qaz!QAZ"),
		database.WithDatabase("genapi"),
	)
	r := gin.Default()
	//加载路由和处理器
	h := handler.Handler{}
	{% for intf in Interfaces %}r.{{intf.Method|upper}}("{{intf.Url}}",h.{{intf.Code|capfirst}})
	{% endfor %}
	if err := r.Run(":{{Port}}"); err != nil {
		log.Fatalln("服务启动失败。。。。。。")
	}
}
