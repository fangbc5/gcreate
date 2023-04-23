package pongo

import (
	"os"
	"strconv"
	"strings"

	"github.com/flosch/pongo2/v6"
)

func Exec(data interface{}) {
	pongo2.RegisterFilter("substr", substr)
	pongo2.RegisterFilter("totype", totype)
	// template, _ := pongo2.FromFile(srcFile)
	// err = template.ExecuteWriter(context, file)
	// if err != nil {
	// 	log.Println(err)
	// }
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
