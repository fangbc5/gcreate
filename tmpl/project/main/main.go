package main

import (
	"{{Code}}/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//加载路由和处理器
	h := handler.Handler{}
	{% for intf in Interfaces %}r.{{intf.Method|upper}}("{{intf.Url}}",h.{{intf.Code|capfirst}})
	{% endfor %}
	if err := r.Run(":{{Port}}"); err != nil {
		log.Fatalln("服务启动失败。。。。。。")
	}
}
