package templates

import (
	"encoding/json"
	"gcreate/conf"
	"gcreate/db"
	"gcreate/handler"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/flosch/pongo2/v6"
)

type Pongo struct {
}

func MakePongo() *Pongo {
	return &Pongo{}
}

func (p *Pongo) Exec(data interface{}) {
	handler.Try(func() {
		pongo2.RegisterFilter("substr", substr)
		// 获取目录下的所有模版文件
		getfiles(conf.MakeConfig().Tmpl, conf.MakeConfig().Out, func(srcFile string, destFile string) {
			template, _ := pongo2.FromFile(srcFile)
			//pongo2.ApplyFilter("firstWord", , pongo2.Value{})
			//准备数据
			tables := data.([]db.Data)
			for _, table := range tables {
				marshal, _ := json.Marshal(table)
				context := pongo2.Context{}
				err := json.Unmarshal(marshal, &context)
				if err != nil {
					log.Println(err)
				}
				//destFile名称处理
				destFile = getFileName(destFile, table.ModelName)
				file, _ := os.Create(destFile)
				err = template.ExecuteWriter(context, file)
				if err != nil {
					log.Println(err)
				}
			}
		})
	}, func(err interface{}) {
		log.Panicln(err)
	})

}

//func getFiles(src string, dest string) map[string]string{
//	os.Stat(src)
//}

type copycall func(srcFile string, destFile string)

func getfiles(src string, dest string, f copycall) {
	dir, _ := os.ReadDir(src)
	for _, file := range dir {
		ndir := src + "/" + file.Name()
		ndest := dest + "/" + file.Name()
		if file.IsDir() {
			err := os.MkdirAll(ndest, 0777)
			if err != nil {
				panic(err)
			}
			getfiles(ndir, ndest, f)
		} else {
			f(ndir, ndest)
		}
	}
}

func getFileName(name string, tname string) string {
	idot := strings.LastIndex(name, ".")
	igang := strings.LastIndex(name, "/")
	suffix := ".go"
	lang := conf.MakeConfig().Lang
	split := ""
	if lang == "go" {
		suffix = ".go"
		split = "_"
	} else if lang == "java" {
		suffix = ".java"
		if igang != -1 {
			if len(name) == igang + 2 {
				name = name[:igang+1] + strings.ToUpper(name[igang+1:igang+2])
			} else {
				name = name[:igang+1] + strings.ToUpper(name[igang+1:igang+2]) + name[igang+2:]
			}
		}
	} else {
		panic("暂不支持该编程语言：" + lang)
	}
	if idot == -1 {
		name = name + suffix
	} else {
		name = strings.ReplaceAll(name, name[strings.LastIndex(name, "."):], suffix)
	}
	if igang == -1 {
		name = tname + split + name
	} else {
		name = name[:igang+1] + tname + split + name[igang+1:]
	}
	return name
}

func substr(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	arr := strings.Split(param.String(), ",")
	begin, err2 := strconv.Atoi(arr[0])
	if err2 != nil {
		return nil, nil
	}
	end, err3 := strconv.Atoi(arr[1])
	if err3 != nil {
		return nil, nil
	}
	out = pongo2.AsValue(in.String()[begin:end])
	return out, nil
}
