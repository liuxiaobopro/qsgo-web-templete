/**
这个文件有系统自动生成,请勿修改
*/
package {{.Package}}

import (
	"{{.Project}}/global"
	"{{.Project}}/models"

	respx "github.com/liuxiaobopro/gobox/resp"
)

type {{.Package}}Dao struct {}

var {{.PackageUpper}}Dao *{{.Package}}Dao

func (th *{{.Package}}Dao) Create({{.Package}} *models.{{.PackageUpper}}) *respx.T {
	if _, err := global.DB.Insert({{.Package}}); err != nil {
		global.ZapS.Errorw("{{.PackageUpper}}Dao.Create", "err", err)
		return respx.InternalErrT
	}
	return nil
}
{{range $col := .Cols}}
func (th *{{$.Package}}Dao) DeleteBy{{$col}}({{$.Package}} *models.{{$.PackageUpper}}) *respx.T {
	if _, err := global.DB.Where("{{$col}} = ?", {{$.Package}}.{{$col}}).Delete(&models.{{$.PackageUpper}}{}); err != nil {
		global.ZapS.Errorw("{{$.PackageUpper}}Dao.Delete", "err", err)
		return respx.InternalErrT
	}
	return nil
}
{{end}}
{{range $col := .Cols}}
func (th *{{$.Package}}Dao) UpdateBy{{$col}}({{$.Package}} *models.{{$.PackageUpper}}) *respx.T {
	if _, err := global.DB.Where("{{$col}} = ?", {{$.Package}}.{{$col}}).Update({{$.Package}}); err != nil {
		global.ZapS.Errorw("{{$.PackageUpper}}Dao.Update", "err", err)
		return respx.InternalErrT
	}
	return nil
}
{{end}}
{{range $col := .Cols}}
func (th *{{$.Package}}Dao) DetailBy{{$col}}({{$.Package}} *models.{{$.PackageUpper}}) (*models.{{$.PackageUpper}}, *respx.T) {
	if _, err := global.DB.Where("{{$col}} = ?", {{$.Package}}.{{$col}}).Get({{$.Package}}); err != nil {
		global.ZapS.Errorw("{{$.PackageUpper}}Dao.Detail", "err", err)
		return nil, respx.InternalErrT
	}
	return {{$.Package}}, nil
}
{{end}}
func (th *{{.Package}}Dao) List() (*respx.List, *respx.T) {
	var (
		{{.Package}}s = make([]*models.{{.PackageUpper}}, 0)
		count int64
		err   error
	)

	if count, err = global.DB.FindAndCount(&{{.Package}}s); err != nil {
		global.ZapS.Errorw("{{.PackageUpper}}Dao.List", "err", err)
		return nil, respx.InternalErrT
	}
	return &respx.List{Count: count, List: {{.Package}}s}, nil
}

