package gtools

import (
	"fmt"
	"os/exec"

	"github.com/tocurd/gormt/data/view/genstruct"
	"github.com/xxjwxc/public/mylog"

	"github.com/tocurd/gormt/data/dlg"
	"github.com/tocurd/gormt/data/view/model"

	"github.com/tocurd/gormt/data/config"

	"github.com/tocurd/gormt/data/view/model/genmssql"
	"github.com/tocurd/gormt/data/view/model/genmysql"
	"github.com/tocurd/gormt/data/view/model/gensqlite"
	"github.com/xxjwxc/public/tools"
)

// Execute exe the cmd
func Execute() (pack genstruct.GenPackage) {
	if config.GetIsGUI() {
		dlg.WinMain()
	} else {
		pack = showCmd()
	}
	return
}

func showCmd() (pack genstruct.GenPackage) {
	// var tt oauth_db.UserInfoTbl
	// tt.Nickname = "ticket_001"
	// orm.Where("nickname = ?", "ticket_001").Find(&tt)
	// fmt.Println(tt)

	dbInfo := config.GetDbInfo()
	for key, dbInfoItem := range dbInfo {

		var modeldb model.IModel
		switch dbInfoItem.Type {
		case 0: // mysql
			modeldb = genmysql.GetModel(dbInfoItem) //
		case 1: // sqllite
			modeldb = gensqlite.GetModel(dbInfoItem) //
		case 2: // mssql
			modeldb = genmssql.GetModel(dbInfoItem) // dbInfoItem
		}
		if modeldb == nil {
			mylog.Error(fmt.Errorf("modeldb not fund : please check db_info.type (0:mysql , 1:sqlite , 2:mssql) "))
			return
		}

		pkg := modeldb.GenModel()
		// gencnf.GenOutPut(&pkg)
		// just for test
		// out, _ := json.Marshal(pkg)
		// tools.WriteFile("test.txt", []string{string(out)}, true)

		pkg.PackageName = key
		list, model := model.Generate(pkg)
		pack = model.GetPackage()

		databseList := ""
		for _, v := range list {

			path := config.GetOutDir() + "/" + key + "/" + v.FileName

			// 增加模型列表
			for index := 0; index < len(pack.Structs); index++ {
				databseList += "\"" + pack.Structs[index].Name + "\": " + pack.Structs[index].Name + "{},\r\n"
			}
			tools.WriteFile(path, []string{v.FileCtx, "var DatabseModel = map[string]interface{}{\r\n" + databseList + "}"}, true)

			mylog.Info("formatting differs from goimport's:")
			cmd, _ := exec.Command("goimports", "-l", "-w", path).Output()
			mylog.Info(string(cmd))

			mylog.Info("formatting differs from gofmt's:")
			cmd, _ = exec.Command("gofmt", "-l", "-w", path).Output()
			mylog.Info(string(cmd))
		}
	}

	return
}
