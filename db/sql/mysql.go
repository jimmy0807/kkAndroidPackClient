package mysql

import (
	"database/sql"
	"fmt"
	"kkAndroidPackClient/config"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mySQL *sql.DB
)

//MySQLReflect 反射
type MySQLReflect struct {
	module Module
}

//Module 记录类型
type Module interface {
	GenerateReflectValue() interface{}
	Insert()
	Next(interface{})
}

//Type 传实例
func (r *MySQLReflect) Type(module Module) *MySQLReflect {
	r.module = module
	return r
}

//InitMySQL 初始化
func InitMySQL() {
	log.Println("go main")
	db, err := sql.Open("mysql", config.SQLDomain)
	if err != nil {
		fmt.Printf("connect error: %s\n", err.Error())
	}
	err = db.Ping()
	mySQL = db
}

//Insert 插入
func Insert(query string, args ...interface{}) {
	_, err2 := mySQL.Exec(query, args...)
	if err2 != nil {
		fmt.Printf("Insert error: %s\n", err2.Error())
	}
}

//Exec 执行
func Exec(query string, args ...interface{}) {
	_, err2 := mySQL.Exec(query, args...)
	if err2 != nil {
		fmt.Printf("Exec error: %s\n", err2.Error())
	}
}

//InsertPrepare 插入
func InsertPrepare(query string) *sql.Stmt {
	stmt, err := mySQL.Prepare(query)

	if err != nil {
		log.Println(err)
	}

	return stmt
}

//QueryOne 查询一条
func (r *MySQLReflect) QueryOne(query string) []interface{} {
	return r.queryList(query, true)
}

//QueryList 查询多条
func (r *MySQLReflect) QueryList(query string) []interface{} {
	return r.queryList(query, false)
}

func (r *MySQLReflect) queryList(query string, isOne bool) []interface{} {
	rows, err := mySQL.Query(query)
	defer rows.Close()
	if err != nil {
		fmt.Println("sql open error ", err)
	}

	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	result := []interface{}{}

	// p := reflect.ValueOf(r.module).Type()
	// reflectStruct := reflect.New(p)

	// fmt.Println(reflectStruct)

	for rows.Next() {
		value := r.module.GenerateReflectValue()
		elem := reflect.ValueOf(value).Elem()

		for i, v := range columns {
			values[i] = elem.FieldByName(v).Addr().Interface()
		}
		rows.Scan(values...)
		r.module.Next(value)
		result = append(result, value)

		if isOne {
			break
		}
	}

	return result
}
