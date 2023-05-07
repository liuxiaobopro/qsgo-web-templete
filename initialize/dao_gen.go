package initialize

import (
	"fmt"
	"os"
	"strings"

	"xorm.io/xorm"
)

type genXormDao struct {
	Mysql   *xorm.Engine // xorm Engine
	DaoPath string       // 生成的dao文件绝对路径
	Prefix  string       // 前缀(用于dao文件的生成是否有前缀和package名, 建议传)

	tableInfo []tableInfo // 表信息
}

type tableInfo struct {
	tableName          string   // 表名
	cols               []string // 列名
	programDaoFileName string   // 程序生成的dao文件名
	userDaoFileName    string   // 用户可修改dao文件名
}

var (
	programDaoFileName = "%s_program.go" // 程序生成的dao文件名
	userDaoFileName    = "%s_default.go" // 用户可修改dao文件名
)

func NewGenXormDao(mysql *xorm.Engine, daoPath, prefix string) *genXormDao {
	if mysql == nil {
		panic("mysql is nil")
	}
	if daoPath == "" {
		daoPath = "./dao"
	}
	return &genXormDao{
		Mysql:   mysql,
		DaoPath: daoPath,
		Prefix:  prefix,
	}
}

func (g *genXormDao) InitData() error {
	// 获取所有表
	tables, err := g.Mysql.DBMetas()
	if err != nil {
		return err
	}

	// 获取所有表列名
	for _, table := range tables {
		// fmt.Printf("table name: %s\n", table.Name)
		ti := tableInfo{
			tableName: table.Name,
		}

		for _, col := range table.Columns() {
			ti.cols = append(ti.cols, col.Name)
			// 判断前缀
			if g.Prefix != "" {
				// 判断表名前缀是否正确(防止乱传, 不匹配就不截取了)
				if strings.HasPrefix(table.Name, g.Prefix) {
					ti.programDaoFileName = fmt.Sprintf(programDaoFileName, table.Name[len(g.Prefix):])
					ti.userDaoFileName = fmt.Sprintf(userDaoFileName, table.Name[len(g.Prefix):])
				} else {
					ti.programDaoFileName = fmt.Sprintf(programDaoFileName, table.Name)
					ti.userDaoFileName = fmt.Sprintf(userDaoFileName, table.Name)
				}
			} else {
				ti.programDaoFileName = fmt.Sprintf(programDaoFileName, table.Name)
				ti.userDaoFileName = fmt.Sprintf(userDaoFileName, table.Name)
			}
		}
		g.tableInfo = append(g.tableInfo, ti)
	}
	return nil
}

func (g *genXormDao) Gen() error {
	if err := g.InitData(); err != nil {
		return err
	}

	for _, v := range g.tableInfo {
		// fmt.Println("table name: ", v.tableName)
		programDaoFilePath := g.DaoPath + "/" + v.programDaoFileName
		userDaoFilePath := g.DaoPath + "/" + v.userDaoFileName

		if _, err := os.Stat(programDaoFilePath); err != nil {
			if !os.IsNotExist(err) {
				return err
			} else {
				// 文件不存在，创建
				if _, err := os.Create(programDaoFilePath); err != nil {
					return err
				}
			}
		} else {
			// 文件存在，删除
			if err := os.Remove(programDaoFilePath); err != nil {
				return err
			}

			// 删除成功，创建
			if _, err := os.Create(programDaoFilePath); err != nil {
				return err
			}
		}

		if _, err := os.Stat(userDaoFilePath); err != nil {
			if os.IsNotExist(err) {
				// 文件不存在，创建
				if _, err := os.Create(userDaoFilePath); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
