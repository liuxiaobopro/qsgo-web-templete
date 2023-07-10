package define

import (
	"encoding/json"

	jsonx "github.com/liuxiaobopro/gobox/json"
)

// DefaultResStyle 默认返回样式
// 1. 处理返回样式(小驼峰, 下划线...)
// 2. 删除不需要返回的字段
// 3. args 预留, 用于扩展
func DefaultResStyle(data interface{}, args ...interface{}) (map[string]interface{}, error) {
	var (
		b   []byte
		err error
	)
	if b, err = json.Marshal(jsonx.JsonCamelCase{Value: data}); err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	//#region 字段处理
	delete(m, "updateAt")
	delete(m, "deleteAt")

	if v, ok := m["list"]; ok {
		// list字段处理
		if list, ok1 := v.([]interface{}); ok1 {
			for _, item := range list {
				if itemMap, ok2 := item.(map[string]interface{}); ok2 {
					delete(m, "updateAt")
					delete(m, "deleteAt")

					//#region 处理useTime
					if v, ok3 := itemMap["useTime"]; ok3 {
						if useTime, ok4 := v.(string); ok4 {
							if useTime == "0001-01-01 00:00:00" {
								itemMap["useTime"] = ""
							}
						}
					}
					//#endregion
				}
			}
		}
	}
	//#endregion
	return m, nil
}
