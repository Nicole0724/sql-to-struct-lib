package service

import (
	"database/sql"
	"fmt"
	"gitee.com/nicole-go-libs/sql-to-struct-lib/internal/model"
	"gitee.com/nicole-go-libs/sql-to-struct-lib/internal/templates"
	"gitee.com/nicole-go-libs/sql-to-struct-lib/libs/words"
	"log"
)

func SqlTwoStructByTable(dbEngine *sql.DB, dataBaseName string, packageName string, saveFilePwd string, tableName string) {
	tableList, _ := model.GetColumnsByDatabaseName(dbEngine, dataBaseName, tableName)
	for _, column := range tableList {
		fmt.Println(column.TableName)
		sqlTwoStructByTableName(dbEngine, dataBaseName, column.TableName, packageName, saveFilePwd)
		//time.Sleep(500 * time.Millisecond) //模拟执行的耗时任务
	}
}

func sqlTwoStructByTableName(dbEngine *sql.DB, dataBaseName string, tableName string, packageName string, saveFilePwd string) {
	// 获取表信息
	tableCol, _ := model.GetColumnsByTableName(dbEngine, dataBaseName, tableName)
	saveFile(tableCol, tableName, packageName, saveFilePwd)
}

func saveFile(tableCol []*model.TableColumn, tableName string, packageName string, saveFilePwd string) {
	// 使用 templete
	tlp := templates.NewStructTemplate()
	structCol := tlp.AssemblyColumns(tableCol)
	err := tlp.Generate(packageName, tableName, structCol, saveFilePwd, fmt.Sprintf("%sModel", words.ToCamelCase(tableName)))
	if err != nil {
		log.Fatalf("template.Generate err: %v", err)
	}

	//println(fmt.Sprintf("%s", structCol))
}
