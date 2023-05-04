package flow

import (
	"gcreate/conf"
	"gcreate/model"
	"log"

	db "github.com/fangbc5/gogo/core/database/mysql"
	"gorm.io/gorm"
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
	intfs := loadApp(api, app)
	c.Dir.ProjectName = app.Code
	//创建项目文件夹
	projectRoot := createProjectDir(app)
	bfs, metaflagintf := loadBusiFlow(api, intfs)
	bfcs, metaflagbf := loadBusiFlowChild(api, bfs)
	metaflagbfc := loadBusiFlowChildOper(api, bfcs)
	
	//并行创建handler
	go createHandler(projectRoot, intfs, metaflagintf)
	//并行创建busiflow
	go createBusiFlow(projectRoot, bfs, metaflagbf)
	//并行创建busiflowchild
	go createBusiFlowChild(projectRoot, bfcs, metaflagbfc)
	//并行创建metadata和metavo
	metaflag := duplicateMeta(metaflagintf,metaflagbf,metaflagbfc)
	go createMetadata(projectRoot, api, metaflag)
	go createMetavo(projectRoot, api, metaflag)
	log.Println("应用全部信息加载完毕！！！")

	execCommand(app.Code, "go", "mod", "init")
	execCommand(app.Code, "go", "mod", "tidy")
}

func loadApp(api *gorm.DB, app *model.Application) []*model.Interface {
	api.Preload("Interfaces").Find(app)
	return app.Interfaces
}

func loadBusiFlow(api *gorm.DB, intfs []*model.Interface) ([]*model.BusiFlow, *Metaflag) {
	bf := make([]*model.BusiFlow, 0)
	metaflag := new(Metaflag)
	metaflag.MetadataMap = make(map[int]interface{}, 0)
	metaflag.MetavoMap = make(map[int]interface{}, 0)
	for _, intf := range intfs {
		api.Preload("BusiFlows").Find(intf)
		bf = append(bf, intf.BusiFlows...)
		if intf.InputParamsType == "metadata" || intf.InputParamsType == "metadata" {
			metaflag.Metadata = true
			if intf.InputParamsId != 0 {
				metaflag.MetadataMap[intf.InputParamsId] = 1
			}
			if intf.OutputParamsId != 0 {
				metaflag.MetadataMap[intf.OutputParamsId] = 1
			}
		}
		if intf.InputParamsType == "metavo" || intf.InputParamsType == "metavo" {
			metaflag.Metavo = true
			if intf.InputParamsId != 0 {
				metaflag.MetavoMap[intf.InputParamsId] = 1
			}
			if intf.OutputParamsId != 0 {
				metaflag.MetavoMap[intf.OutputParamsId] = 1
			}
		}
	}
	return bf, metaflag
}

func loadBusiFlowChild(api *gorm.DB, bfs []*model.BusiFlow) ([]*model.BusiFlowChild, *Metaflag) {
	bfc := make([]*model.BusiFlowChild, 0)
	metaflag := new(Metaflag)
	metaflag.MetadataMap = make(map[int]interface{}, 0)
	metaflag.MetavoMap = make(map[int]interface{}, 0)
	for _, bf := range bfs {
		api.Preload("Childs").Find(bf)
		bfc = append(bfc, bf.Childs...)
		if bf.InputParamsType == "metadata" || bf.InputParamsType == "metadata" {
			metaflag.Metadata = true
			if bf.InputParamsId != 0 {
				metaflag.MetadataMap[bf.InputParamsId] = 1
			}
			if bf.OutputParamsId != 0 {
				metaflag.MetadataMap[bf.OutputParamsId] = 1
			}
		}
		if bf.InputParamsType == "metavo" || bf.InputParamsType == "metavo" {
			metaflag.Metavo = true
			if bf.InputParamsId != 0 {
				metaflag.MetavoMap[bf.InputParamsId] = 1
			}
			if bf.OutputParamsId != 0 {
				metaflag.MetavoMap[bf.OutputParamsId] = 1
			}
		}
	}
	return bfc, metaflag
}

func loadBusiFlowChildOper(api *gorm.DB, bfcs []*model.BusiFlowChild) *Metaflag {
	metaflag := new(Metaflag)
	metaflag.MetadataMap = make(map[int]interface{}, 0)
	metaflag.MetavoMap = make(map[int]interface{}, 0)
	for _, bfc := range bfcs {
		api.Preload("ActionDao").Preload("ActionDco").Preload("ActionRpc").Preload("ActionMsg").Find(bfc)
		if bfc.InputParamsType == "metadata" || bfc.InputParamsType == "metadata" {
			metaflag.Metadata = true
			if bfc.InputParamsId != 0 {
				metaflag.MetadataMap[bfc.InputParamsId] = 1
			}
			if bfc.OutputParamsId != 0 {
				metaflag.MetadataMap[bfc.OutputParamsId] = 1
			}
		}
		if bfc.InputParamsType == "metavo" || bfc.InputParamsType == "metavo" {
			metaflag.Metavo = true
			if bfc.InputParamsId != 0 {
				metaflag.MetavoMap[bfc.InputParamsId] = 1
			}
			if bfc.OutputParamsId != 0 {
				metaflag.MetavoMap[bfc.OutputParamsId] = 1
			}
		}
	}
	return metaflag
}
