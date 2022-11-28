package sqlTools

import (
	"database/sql"
	"fmt"
	"gitee.com/nicole-go-libs/print-colors/color_print"
	"gitee.com/nicole-go-libs/sql-to-struct-lib/internal/service"
)

var (
	Conf *SqlToStructParams
)

type SqlConnectConfig struct {
	Type         string
	Host         string
	Port         string
	Account      string
	Password     string
	DataBaseName string
}

type SqlToStructParams struct {
	SqlConfig     SqlConnectConfig
	PackageName   string
	IsAllDataBase bool
	TableName     string
	SaveFileName  string
	SaveFilePwd   string
}

func ConnectSql(sqlConfig SqlConnectConfig) *sql.DB {
	db, err := sql.Open(sqlConfig.Type, fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		sqlConfig.Account,
		sqlConfig.Password,
		sqlConfig.Host,
		sqlConfig.Port,
		sqlConfig.DataBaseName,
	))

	if err != nil {
		fmt.Println(color_print.Red(fmt.Sprintf("初始化数据连接失败，失败原因：\n%s", err)))
		panic(err)
	}

	fmt.Println(color_print.Green("数据库连接成功!"))
	return db
}

func RunSqlToStruct(dbEngine *sql.DB, dataBaseName string, packageName string, saveFilePwd string, tableName string) {
	service.SqlTwoStructByTable(dbEngine, dataBaseName, packageName, saveFilePwd, tableName)
}

func CloseMysql(dbEngine *sql.DB) {
	defer func() {
		_ = dbEngine.Close()
		fmt.Println(color_print.Yellow("数据库断开成功!"))
	}()
}
