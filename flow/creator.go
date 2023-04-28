package flow

import (
	"gcreate/model"
	"gcreate/pongo"
	"log"
	"os"
	"strings"
)

func createMetavo(projectRoot string, metavo *model.Metavo) {
	metavoDir := projectRoot + "/metavo"
	meta := metavo.Meta
	code := metavo.Code
	if meta != "" && code != "" {
		pongo.ExecOneFromStringTmpl(meta, metavoDir+"/"+strings.ToLower(code)+".go", nil, os.O_CREATE)
	}
}

func createMetadata(projectRoot string, metadata *model.Metadata) {
	metadataDir := projectRoot + "/metadata"
	meta := metadata.Meta
	code := metadata.Code
	if meta != "" && code != "" {
		pongo.ExecOneFromStringTmpl(meta, metadataDir+"/"+strings.ToLower(code)+".go", nil, os.O_CREATE)
	}
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
	return os.IsExist(err)
}
