package test

import (
	"database/sql"
	"gitee.com/nicole-go-libs/sql-to-struct-lib/internal/service"
	"gitee.com/nicole-go-libs/sql-to-struct-lib/libs/sqlTools"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

var (
	DBEngine    *sql.DB
	SaveFilePwd = "D:\\work_space\\open_source\\go-libs\\sql-to-struct-lib\\entity\\" //绝对路径方式,需要保证当前路径确实存在
)

func TestWorking(t *testing.T) {

	// 连接数据库
	var sqlConfig sqlTools.SqlConnectConfig

	sqlConfig.Type = "mysql"
	sqlConfig.Host = "127.0.0.1"
	sqlConfig.Port = "3306"
	sqlConfig.Account = "root"
	sqlConfig.Password = "root"
	sqlConfig.DataBaseName = "go_frame"

	DBEngine := sqlTools.ConnectSql(sqlConfig)

	// 查询要转换的表
	service.SqlTwoStructByTable(DBEngine, sqlConfig.DataBaseName, "entityModel", SaveFilePwd, "dict")

	if DBEngine != nil {
		sqlTools.CloseMysql(DBEngine)
	}

}
