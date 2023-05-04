package flow

import (
	"gcreate/model"
	"gcreate/pongo"
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

func duplicateMeta(metaflags ...*Metaflag) *Metaflag {
	metaflag := &Metaflag{}
	metadataMapDuplicate := make(map[int]interface{}, 0)
	metavoMapDuplicate := make(map[int]interface{}, 0)
	for _, flag := range metaflags {
		for k, v := range flag.MetadataMap {
			if metadataMapDuplicate[k] == nil {
				metadataMapDuplicate[k] = v
			}
		}
		for k, v := range flag.MetavoMap {
			if metavoMapDuplicate[k] == nil {
				metavoMapDuplicate[k] = v
			}
		}
	}
	metaflag.MetadataMap = metadataMapDuplicate
	metaflag.MetavoMap = metavoMapDuplicate
	return metaflag
}

func createMetavo(projectRoot string, api *gorm.DB, metaflag *Metaflag) {
	metavoDir := projectRoot + "/metavo"
	m := metaflag.MetavoMap
	for id := range m {
		Metavo := &model.Metavo{}
		api.Table("metavo").Where("id = ?", id).Scan(Metavo)
		meta := Metavo.Meta
		code := Metavo.Code
		if meta != "" && code != "" {
			pongo.ExecOneFromStringTmpl(meta, metavoDir+"/"+strings.ToLower(code)+".go", nil, os.O_CREATE)
		}
	}
}

func createMetadata(projectRoot string, api *gorm.DB, metaflag *Metaflag) {
	metadataDir := projectRoot + "/metadata"
	m := metaflag.MetadataMap
	for id := range m {
		metadata := &model.Metadata{}
		api.Table("metadata").Where("id = ?", id).Scan(metadata)
		meta := metadata.Meta
		code := metadata.Code
		if meta != "" && code != "" {
			pongo.ExecOneFromStringTmpl(meta, metadataDir+"/"+strings.ToLower(code)+".go", nil, os.O_CREATE)
		}
	}
}

func createHandler(projectRoot string, intfs []*model.Interface, metaflag *Metaflag) {
	handlerDir := projectRoot + "/handler"
	pongo.ExecOne("tmpl/project/handler/header.go", projectRoot+"/handler/handler.go", metaflag, os.O_CREATE)
	for _, intf := range intfs {
		pongo.ExecOne("tmpl/project/handler/handler.go", handlerDir+"/handler.go", intf, os.O_APPEND)
	}
}

func createBusiFlow(projectRoot string, busiflows []*model.BusiFlow, metaflag *Metaflag) {
	busiflowDir := projectRoot + "/busiflow"
	pongo.ExecOne("tmpl/project/busiflow/header.go", projectRoot+"/busiflow/busiflow.go", metaflag, os.O_CREATE)
	for _, busiflow := range busiflows {
		pongo.ExecOne("tmpl/project/busiflow/busiflow.go", busiflowDir+"/busiflow.go", busiflow, os.O_APPEND)
	}
}

func createBusiFlowChild(projectRoot string, busiflowchilds []*model.BusiFlowChild, metaflag *Metaflag) {
	busiflowchildDir := projectRoot + "/busiflowchild"
	pongo.ExecOne("tmpl/project/busiflowchild/header.go", projectRoot+"/busiflowchild/busiflowchild.go", metaflag, os.O_CREATE)
	for _, busiflowchild := range busiflowchilds {
		pongo.ExecOne("tmpl/project/busiflowchild/busiflowchild.go", busiflowchildDir+"/busiflowchild.go", busiflowchild, os.O_APPEND)
	}
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

func createMainFile(projectRoot string, app *model.Application) error {
	return pongo.Exec("tmpl/project/main", projectRoot, app, os.O_CREATE)
}

func createProjectDir(app *model.Application) string {
	log.Println("创建项目文件夹")
	//创建项目文件夹
	projectRoot := createProject(app.Code)
	if err := createMainFile(projectRoot, app); err != nil {
		log.Fatal(err)
	}
	//创建所有接口、流程、子流程、元数据、视图元数据文件夹
	createDir(projectRoot + "/metadata")
	createDir(projectRoot + "/metavo")
	createDir(projectRoot + "/handler")
	createDir(projectRoot + "/busiflow")
	createDir(projectRoot + "/busiflowchild")
	return projectRoot
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
