package db

import (
	"database/sql"
	"fmt"
	"gcreate/conf"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"sync"
	"time"
)

var ds *Datasource
var lock sync.Mutex

func MakeMysql() *Datasource {
	if ds == nil {
		lock.Lock()
		if ds == nil {
			ds = &Datasource{}
			c := conf.MakeConfig()
			ds.username = c.Username
			ds.password = c.Password
			ds.driver = c.Driver
			ds.schema = c.Schema
			ds.url = c.Url
			ds.url = ds.username + ":" + ds.password + "@" + ds.url + "/" + ds.schema
		}
		lock.Unlock()
	}
	return ds
}

func (d *Datasource) Conn() {
	//"user:password@tcp(ip:port)/db"
	//dataSourceName := config.Model.Mysql.Username + ":" + config.Model.Mysql.Password + "@" + config.Model.Mysql.Url
	if d.DB == nil {
		lock.Lock()
		if d.DB == nil {
			db, err := sql.Open(d.driver, d.url)
			if err != nil {
				panic(err)
			}
			db.SetMaxOpenConns(10)
			db.SetMaxIdleConns(5)
			db.SetConnMaxLifetime(time.Hour)
			db.SetConnMaxIdleTime(time.Second * 3)
			ds.DB = db
		}
		lock.Unlock()
	}
}

func (d *Datasource) Close() {
	err := d.DB.Close()
	if err != nil {
		log.Printf("datasource close fail: %v", err)
	}
}

func (d *Datasource) GetTables(names ...string) []Data {
	conf := conf.MakeConfig()
	tables := make([]Data, 0)
	for _, name := range names {
		t := Data{}
		t.ProjectName = conf.Project.Name
		t.PackName = conf.Project.PackName
		t.ModuleName = conf.Project.Module
		t.TableName = name
		t.ModelName = strings.ReplaceAll(name, conf.Prefix, "")
		fields := make([]Field, 0)
		sqlstr := "select COLUMN_NAME,COLUMN_COMMENT,DATA_TYPE,IS_NULLABLE,COLUMN_KEY from information_schema.COLUMNS where TABLE_SCHEMA = ? and TABLE_NAME = ?"
		rows, _ := d.DB.Query(sqlstr, conf.Schema, name)
		for rows.Next() {
			var f Field
			err := rows.Scan(&f.FieldName, &f.FieldDesc, &f.DataType, &f.IsNull, &f.Primary)
			if err != nil {
				panic(err)
			}
			ignore := getIgore(f.FieldName)
			if ignore {
				continue
			}
			f.DataType = toType(f.DataType)
			f.FieldName = camelName(f.FieldName)
			if f.DataType == "time.Time" {
				t.HasTime = true
			}
			fields = append(fields, f)
		}
		t.Fields = fields
		tables = append(tables, t)
	}
	return tables
}

func getIgore(name string) bool {
	switch strings.ToLower(name) {
	case "id":
		return true
	case "created_at":
		return true
	case "updated_at":
		return true
	case "deleted_at":
		return true
	default:
		return false
	}
}

// todo 优化算法
func camelName(o string) string {
	for strings.IndexAny(o, "_") != -1 {
		i := strings.IndexAny(o, "_")
		old := o[i : i+2]
		n := strings.ToUpper(o[i+1 : i+2])
		o = strings.ReplaceAll(o, old, n)
	}
	return o
}

func toType(o string) (n string) {
	switch o {
	case "text", "tinytext", "mediumtext", "longtext", "varchar":
		n = "string"
	case "bigint":
		n = "uint"
	case "int", "tinyint", "smallint":
		n = "uint"
	case "float", "double", "decimal":
		n = "float64"
	case "date", "time", "datetime", "timestamp":
		n = "time.Time"
	default:
		n = "string"
	}
	return
}

func (d *Datasource) Execsql(sqlstr string) {
	rows, err := d.DB.Query(sqlstr)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		fmt.Println(rows.Columns())
	}
}
