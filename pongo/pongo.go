package pongo

import (
	"encoding/json"
	"gcreate/conf"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/flosch/pongo2/v6"
)

func Exec(srcDir string, destDir string, data interface{}, flag int) error {
	pongo2.RegisterFilter("substr", substr)
	pongo2.RegisterFilter("totype", totype)
	pongo2.RegisterFilter("upper", upper)
	marshal, _ := json.Marshal(data)
	context := pongo2.Context{}
	err := json.Unmarshal(marshal, &context)
	if err != nil {
		log.Fatal(err)
	}
	getfiles(srcDir, destDir, func(srcFile, destFile string) {
		template, err := pongo2.FromFile(srcFile)
		if _, err := os.Stat(destFile); err != nil {
			os.Create(destFile)
		}
		file, err := os.OpenFile(destFile, flag, 0644)
		defer file.Close()
		if err != nil {
			log.Fatalf("pongo2：文件写入失败，%v", err)
		}

		err = template.ExecuteWriter(context, file)
		if err != nil {
			log.Fatalf("pongo2：文件写入失败，%v", err)
		}
	})
	return nil
}

func ExecOne(srcFile string, destFile string, data interface{}, flag int) {
	pongo2.RegisterFilter("substr", substr)
	pongo2.RegisterFilter("totype", totype)
	pongo2.RegisterFilter("upper", upper)
	marshal, _ := json.Marshal(data)
	context := pongo2.Context{}
	err := json.Unmarshal(marshal, &context)
	if err != nil {
		log.Fatal(err)
	}
	template, err := pongo2.FromFile(srcFile)
	var file *os.File
	file.Close()
	if flag == os.O_CREATE {
		file, _ = os.Create(destFile)
	} else {
		file, _ = os.OpenFile(destFile, os.O_APPEND, 0644)
	}
	context["ProjectName"] = conf.GetConfig().Dir.ProjectName
	template.ExecuteWriter(context, file)
}

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
	if idot == -1 {
		name = name + ".go"
	} else {
		name = strings.ReplaceAll(name, name[strings.LastIndex(name, "."):], ".go")
	}
	if igang == -1 {
		name = tname + "_" + name
	} else {
		name = name[:igang+1] + tname + "_" + name[igang+1:]
	}
	return name
}

func upper(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	result := strings.ToUpper(in.String())
	return pongo2.AsSafeValue(result), nil
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

func totype(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	var dataType string
	switch in.String() {
	case "text", "tinytext", "mediumtext", "longtext", "varchar":
		dataType = "string"
	case "bigint":
		dataType = "int64"
	case "int", "tinyint", "smallint":
		dataType = "int"
	case "float", "double", "decimal":
		dataType = "float64"
	case "date", "time", "datetime", "timestamp":
		dataType = "time.Time"
	default:
		dataType = "string"
	}
	out = pongo2.AsValue(dataType)
	return out, nil
}
