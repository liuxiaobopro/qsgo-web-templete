/**
这个文件有系统自动生成,请勿修改
*/
package {{.Package}}

import (
	"time"

	"{{.Project}}/global"
	"{{.Project}}/models"

	replyx "github.com/liuxiaobopro/gobox/reply"
	timex "github.com/liuxiaobopro/gobox/time"
)

type {{.Package}}Dao struct {}

var {{.PackageUpper}}Dao *{{.Package}}Dao

func (th *{{.Package}}Dao) Create({{.Package}} *models.{{.PackageUpper}}) *replyx.T {
	{{.Package}}.CreateAt = timex.Time(time.Now())
	if _, err := global.DB.Insert({{.Package}}); err != nil {
		global.ZapS.Errorw("{{.PackageUpper}}Dao.Create", "err", err)
		return replyx.InternalErrT
	}
	return nil
}
{{range $col := .Cols}}
func (th *{{$.Package}}Dao) DeleteBy{{$col}}({{$.Package}} *models.{{$.PackageUpper}}) *replyx.T {
	if _, err := global.DB.Where("{{$col}} = ?", {{$.Package}}.{{$col}}).Delete(&models.{{$.PackageUpper}}{}); err != nil {
		global.ZapS.Errorw("{{$.PackageUpper}}Dao.Delete", "err", err)
		return replyx.InternalErrT
	}
	return nil
}
{{end}}
{{range $col := .Cols}}
func (th *{{$.Package}}Dao) UpdateBy{{$col}}({{$.Package}} *models.{{$.PackageUpper}}) *replyx.T {
	{{$.Package}}.UpdateAt = timex.Time(time.Now())
	if _, err := global.DB.Where("{{$col}} = ?", {{$.Package}}.{{$col}}).Update({{$.Package}}); err != nil {
		global.ZapS.Errorw("{{$.PackageUpper}}Dao.Update", "err", err)
		return replyx.InternalErrT
	}
	return nil
}
{{end}}
{{range $col := .Cols}}
func (th *{{$.Package}}Dao) DetailBy{{$col}}({{$.Package}} *models.{{$.PackageUpper}}) (*models.{{$.PackageUpper}}, *replyx.T) {
	if _, err := global.DB.Where("{{$col}} = ?", {{$.Package}}.{{$col}}).Get({{$.Package}}); err != nil {
		global.ZapS.Errorw("{{$.PackageUpper}}Dao.Detail", "err", err)
		return nil, replyx.InternalErrT
	}
	return {{$.Package}}, nil
}
{{end}}
{{range $col := .Cols}}
func (th *{{$.Package}}Dao) ExistBy{{$col}}({{$.Package}} *models.{{$.PackageUpper}}) (bool, *replyx.T) {
	var (
		has bool
		err error
	)
	if has, err = global.DB.Where("{{$col}} = ?", {{$.Package}}.{{$col}}).Exist({{$.Package}}); err != nil {
		global.ZapS.Errorw("{{$.PackageUpper}}Dao.Detail", "err", err)
		return false, replyx.InternalErrT
	}
	return has, nil
}
{{end}}
func (th *{{.Package}}Dao) List(page, size int) (*replyx.List, *replyx.T) {
	var (
		{{.Package}}s = make([]*models.{{.PackageUpper}}, 0)
		count int64
		err   error
	)

	if count, err = global.DB.Limit(size, (page-1)*size).FindAndCount(&{{.Package}}s); err != nil {
		global.ZapS.Errorw("{{.PackageUpper}}Dao.List", "err", err)
		return nil, replyx.InternalErrT
	}
	return &replyx.List{Count: count, List: {{.Package}}s}, nil
}

