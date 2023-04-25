package flow

import (
	"fmt"
	"gcreate/conf"
	"gcreate/model"
	"gcreate/pongo"
	"log"
	"os"
	"os/exec"

	db "github.com/fangbc5/gogo/core/database/mysql"
)

func Start(c *conf.Configuration) {
	//获取数据库连接
	db.Init(db.WithAddress(c.Address),
		db.WithPort(c.Port),
		db.WithUsername(c.Username),
		db.WithPassword(c.Password),
		db.WithDatabase(c.Database),
	)
	//加载application
	api := db.GetGormApi()
	app := &model.Application{Id: 1}
	api.Preload("Interfaces").Find(app)
	c.Dir.ProjectName = app.Code
	log.Println("创建项目文件夹")
	//创建项目文件夹
	projectRoot := createProject(app.Code)
	if err := createMainFile(projectRoot, app); err != nil {
		log.Fatal(err)
	}
	var i, j, k int = 0, 0, 0
	for _, intf := range app.Interfaces {
		if i == 0 {
			createDir(projectRoot + "/handler")
			pongo.ExecOne("tmpl/project/handler/header.go", projectRoot+"/handler/handler.go", app, os.O_CREATE)
			i++
		}
		api.Preload("BusiFlows").Find(intf)
		log.Println("生成处理器: " + intf.Name)
		createHandler(projectRoot, intf)
		for _, flowmain := range intf.BusiFlows {
			if j == 0 {
				createDir(projectRoot + "/flowmain")
				pongo.ExecOne("tmpl/project/flowmain/header.go", projectRoot+"/flowmain/flowmain.go", intf, os.O_CREATE)
				j++
			}
			api.Preload("Childs").Find(flowmain)
			log.Println("生成flowmain: " + flowmain.Name)
			createFlowMain(projectRoot, flowmain)
			for _, flowchild := range flowmain.Childs {
				if k == 0 {
					createDir(projectRoot + "/flowchild")
					pongo.ExecOne("tmpl/project/flowchild/header.go", projectRoot+"/flowchild/flowchild.go", flowmain, os.O_CREATE)
					k++
				}
				api.Preload("ActionDao").Preload("ActionDco").Preload("ActionRpc").Preload("ActionMsg").Find(flowchild)
				log.Println("生成flowchild: " + flowchild.Name)
				createFlowChild(projectRoot, flowchild)
				//生成dao
				//生成dco
				//生成rpc
				//生成msg
			}
		}
	}
	execCommand(app.Code,"go","mod","init")
	execCommand(app.Code,"go","mod","tidy")
}

func createFlowMain(projectRoot string, flowmain *model.BusiFlow) {
	flowmainDir := projectRoot + "/flowmain"
	pongo.ExecOne("tmpl/project/flowmain/flowmain.go", flowmainDir+"/flowmain.go", flowmain, os.O_APPEND)
}

func createFlowChild(projectRoot string, flowchild *model.BusiFlowChild) {
	flowmainDir := projectRoot + "/flowchild"
	pongo.ExecOne("tmpl/project/flowchild/flowchild.go", flowmainDir+"/flowchild.go", flowchild, os.O_APPEND)
}

func createMainFile(projectRoot string, app *model.Application) error {
	return pongo.Exec("tmpl/project/main", projectRoot, app, os.O_CREATE)
}

func createHandler(projectRoot string, intf *model.Interface) {
	handlerDir := projectRoot + "/handler"
	pongo.ExecOne("tmpl/project/handler/handler.go", handlerDir+"/handler.go", intf, os.O_APPEND)
}

func createProject(projectName string) string {
	if projectName == "" {
		log.Fatalln("配置文件中dir.projectName值不能为空")
	}
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatalln("GOPATH未定义，请设置GOPATH后重试")
	}
	dir := gopath + "/src/" + projectName
	createDir(dir)
	return dir
}

func createDir(dir string) {
	if !isExist(dir) {
		// 目录不存在，创建
		os.Mkdir(dir, 0755)
	}
}

func isExist(dir string) bool {
	_, err := os.Stat(dir)
	if os.IsExist(err) {
		return true
	}
	return false
}

func execCommand(projectName string, command string , args ...string) error {
	gopath := os.Getenv("GOPATH") // GOPATH环境变量指定的目录
	cmd := exec.Command(command, args...)
	cmd.Dir = gopath + "/src/" + projectName
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return err
	}
	log.Println(string(output))
	return nil
}
