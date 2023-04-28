package flow

import (
	"gcreate/conf"
	"gcreate/model"
	"gcreate/pongo"
	"log"
	"os"

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
	metaAllTypeMap := make(map[string]interface{}, 10)
	for _, intf := range app.Interfaces {
		metaAllTypeMap[intf.InputParamsType] = nil
		metaAllTypeMap[intf.OutputParamsType] = nil
	}
	//创建所有接口、流程、子流程需要用到的metadata和metavo的map
	createDir(projectRoot + "/metadata")
	createDir(projectRoot + "/metavo")
	metadatamap := make(map[int]interface{}, 32)
	metavomap := make(map[int]interface{}, 32)
	var i, j, k int = 0, 0, 0
	for _, intf := range app.Interfaces {
		if i == 0 {
			createDir(projectRoot + "/handler")
			pongo.ExecOne("tmpl/project/handler/header.go", projectRoot+"/handler/handler.go", app, os.O_CREATE)
			i++
		}
		if intf.InputParamsId != 0 && intf.InputParamsCode != "" {
			if intf.InputParamsType == "metadata" {
				metadatamap[intf.InputParamsId] = nil
			} else if intf.InputParamsType == "metavo" {
				metavomap[intf.InputParamsId] = nil
			}
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
			if flowmain.InputParamsId != 0 && flowmain.InputParamsCode != "" {
				if flowmain.InputParamsType == "metadata" {
					metadatamap[intf.InputParamsId] = nil
				} else if flowmain.InputParamsType == "metavo" {
					metavomap[intf.InputParamsId] = nil
				}
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
				if flowchild.InputParamsId != 0 && flowchild.InputParamsCode != "" {
					if flowchild.InputParamsType == "metadata" {
						metadatamap[intf.InputParamsId] = nil
					} else if flowchild.InputParamsType == "metavo" {
						metavomap[intf.InputParamsId] = nil
					}
				}
				api.Preload("ActionDao").Preload("ActionDco").Preload("ActionRpc").Preload("ActionMsg").Find(flowchild)
				log.Println("生成flowchild: " + flowchild.Name)
				createFlowChild(projectRoot, flowchild)
			}
		}
	}
	for key := range metadatamap {
		metadata := &model.Metadata{}
		api.Table("metadata").Where("id = ?", key).First(metadata)
		createMetadata(projectRoot, metadata)
	}
	for key := range metavomap {
		metavo := &model.Metavo{}
		api.Table("metavo").Where("id = ?", key).First(metavo)
		createMetavo(projectRoot, metavo)
	}
	execCommand(app.Code, "go", "mod", "init")
	execCommand(app.Code, "go", "mod", "tidy")
}
