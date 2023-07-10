/*
这个文件有系统自动生成,请勿修改
*/
package {{.Package}}

import (
	"qsgo-web-templete/define"
	"{{.Project}}/global"
	"{{.Project}}/models"

	"github.com/gin-gonic/gin"
	replyx "github.com/liuxiaobopro/gobox/reply"
)

type {{.Package}}Dao struct {}

var {{.PackageUpper}}Dao *{{.Package}}Dao

func (th *{{.Package}}Dao) Create_(c *gin.Context, {{.PackageLower}} *models.{{.PackageUpper}}) *replyx.T {
	if _, err := global.DB.Insert({{.PackageLower}}); err != nil {
		global.Logger.Errorf(c, "{{.PackageUpper}}Dao.Create err %v", err)
		return replyx.InternalErrT
	}
	return nil
}
{{range $col := .Cols}}
func (th *{{$.Package}}Dao) DeleteBy{{index $col 0}}_(c *gin.Context, {{$.PackageLower}} *models.{{$.PackageUpper}}) *replyx.T {
	if _, err := global.DbSession().Where("{{index $col 1}} = ?", {{$.PackageLower}}.{{index $col 0}}).Delete(&models.{{$.PackageUpper}}{}); err != nil {
		global.Logger.Errorf(c, "{{$.PackageUpper}}Dao.Delete err %v", err)
		return replyx.InternalErrT
	}
	return nil
}
{{end}}
{{range $col := .Cols}}
func (th *{{$.Package}}Dao) UpdateBy{{index $col 0}}_(c *gin.Context, {{$.PackageLower}} *models.{{$.PackageUpper}}) *replyx.T {
	if _, err := global.DbSession().Where("{{index $col 1}} = ?", {{$.PackageLower}}.{{index $col 0}}).Update({{$.PackageLower}}); err != nil {
		global.Logger.Errorf(c, "{{$.PackageUpper}}Dao.Update err %v", err)
		return replyx.InternalErrT
	}
	return nil
}
{{end}}
{{range $col := .Cols}}
func (th *{{$.Package}}Dao) DetailBy{{index $col 0}}_(c *gin.Context, {{$.PackageLower}} *models.{{$.PackageUpper}}) (*models.{{$.PackageUpper}}, *replyx.T) {
	var {{$.PackageLower}}Info = &models.{{$.PackageUpper}}{}
	if _, err := global.DbSession().Where("{{index $col 1}} = ?", {{$.PackageLower}}.{{index $col 0}}).Get({{$.PackageLower}}Info); err != nil {
		global.Logger.Errorf(c, "{{$.PackageUpper}}Dao.Detail err %v", err)
		return nil, replyx.InternalErrT
	}
	return {{$.PackageLower}}Info, nil
}
{{end}}
{{range $col := .Cols}}
func (th *{{$.Package}}Dao) ExistBy{{index $col 0}}_(c *gin.Context, {{$.PackageLower}} *models.{{$.PackageUpper}}) (bool, *replyx.T) {
	var (
		has bool
		err error
	)
	if has, err = global.DbSession().Where("{{index $col 1}} = ?", {{$.PackageLower}}.{{index $col 0}}).Exist(&models.{{$.PackageUpper}}{}); err != nil {
		global.Logger.Errorf(c, "{{$.PackageUpper}}Dao.Detail err %v", err)
		return false, replyx.InternalErrT
	}
	return has, nil
}
{{end}}
func (th *{{.Package}}Dao) List_(c *gin.Context, page, size int) (*replyx.List, *replyx.T) {
	var (
		{{.Package}}s = make([]*models.{{.PackageUpper}}, 0)
		count int64
		err   error
	)

	if count, err = global.DbSession().Limit(size, (page-1)*size).OrderBy(define.OrmDefaultOrder).FindAndCount(&{{.Package}}s); err != nil {
		global.Logger.Errorf(c, "{{.PackageUpper}}Dao.List err %v", err)
		return nil, replyx.InternalErrT
	}
	return &replyx.List{Count: count, List: {{.Package}}s}, nil
}

