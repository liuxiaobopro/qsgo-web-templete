package define

const (
	OrmDefaultOrder = "create_at Desc" // orm默认排序
)

const (
	SexUnknown int = iota - 1 // 未知
	SexMan     int = iota     // 男
	SexWoman                  // 女
)

var SexList = []int{SexUnknown, SexMan, SexWoman}
