package models

import (
	time "github.com/liuxiaobopro/gobox/time"
)

type Order struct {
	Oid         uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT" json:"oid"`
	TradeNo     string    `xorm:"not null default '' comment('交易ID') VARCHAR(100)" json:"trade_no"`
	PrepayId    string    `xorm:"not null default '' comment('预支付订单ID') VARCHAR(100)" json:"prepay_id"`
	OrderId     string    `xorm:"not null default '' comment('微信订单号') VARCHAR(100)" json:"order_id"`
	RefundId    string    `xorm:"not null default '' comment('退款编号') VARCHAR(100)" json:"refund_id"`
	OpenId      string    `xorm:"not null default '' comment('openid') VARCHAR(100)" json:"open_id"`
	NonceStr    string    `xorm:"not null default '' comment('微信支付noncestr') VARCHAR(500)" json:"nonce_str"`
	PaySign     string    `xorm:"not null default '' comment('微信支付paysign') VARCHAR(500)" json:"pay_sign"`
	Package     string    `xorm:"not null default '' comment('package') VARCHAR(500)" json:"package"`
	Uid         uint64    `xorm:"not null default 0 comment('用户ID') index(IDX_uid) UNSIGNED BIGINT" json:"uid"`
	Description string    `xorm:"not null default '' comment('描述') VARCHAR(500)" json:"description"`
	Attach      string    `xorm:"not null default '' comment('描述') VARCHAR(500)" json:"attach"`
	ExpireTime  time.Time `xorm:"comment('过期时间') TIMESTAMP" json:"expire_time"`
	Price       uint64    `xorm:"not null default 0 comment('价格') UNSIGNED BIGINT" json:"price"`
	OrderType   int       `xorm:"not null default 0 comment('订单类型 1 自动续费，2季度会员，3年卡会员') TINYINT" json:"order_type"`
	PayType     int       `xorm:"not null default 0 comment('订单类型 1 微信，2 支付宝，3苹果') TINYINT" json:"pay_type"`
	State       int       `xorm:"not null default 0 comment('订单状态 1 prepay,2 prepay接口失败,3 成功 ,4失败，退款') TINYINT" json:"state"`
	Openid      string    `xorm:"not null default '' comment('用户平台ID') VARCHAR(100)" json:"openid"`
	DelFlag     int       `xorm:"not null default 0 comment('0新建  1已删除') TINYINT" json:"del_flag"`
	CreateTime  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index(IDX_uid) TIMESTAMP" json:"create_time"`
	UpdateTime  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP" json:"update_time"`
}

func (m *Order) TableComment() string {
	return "order"
}

type Demo struct {
	Id       int       `xorm:"not null pk autoincr INT" json:"id"`
	Name     string    `xorm:"VARCHAR(255)" json:"name"`
	CreateAt time.Time `xorm:"comment('创建时间') DATETIME" json:"create_at"`
	UpdateAt time.Time `xorm:"comment('更新时间') DATETIME" json:"update_at"`
}

func (m *Demo) TableComment() string {
	return "demo"
}

type User struct {
	Uid        uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT" json:"uid"`
	Name       string    `xorm:"not null default '' comment('用户名') index VARCHAR(512)" json:"name"`
	Password   string    `xorm:"not null default '' comment('密码') index VARCHAR(512)" json:"password"`
	Token      string    `xorm:"not null default '' comment('token') index VARCHAR(500)" json:"token"`
	Age        int       `xorm:"not null default 0 comment('年龄') index(IDX_sex_age) INT" json:"age"`
	Sex        int       `xorm:"not null default 0 comment('性别 0为止 1男 2女 3中性') index(IDX_sex_age) TINYINT" json:"sex"`
	Tel        string    `xorm:"not null default '' comment('电话') index VARCHAR(100)" json:"tel"`
	Email      string    `xorm:"not null default '' comment('email') VARCHAR(100)" json:"email"`
	Weixin     string    `xorm:"not null default '' comment('weixin') VARCHAR(100)" json:"weixin"`
	Flag       string    `xorm:"not null default '' comment('个性签名') VARCHAR(500)" json:"flag"`
	Avatar     string    `xorm:"not null default '' comment('头像') VARCHAR(500)" json:"avatar"`
	Cover      string    `xorm:"not null default '' comment('封面图') VARCHAR(500)" json:"cover"`
	Area       string    `xorm:"not null default '' comment('地区') VARCHAR(100)" json:"area"`
	Follows    int       `xorm:"not null default 0 comment('被关注数量') INT" json:"follows"`
	Watches    int       `xorm:"not null default 0 comment('关注数量') INT" json:"watches"`
	Likes      int       `xorm:"not null default 0 comment('获赞') INT" json:"likes"`
	DelFlag    int       `xorm:"not null default 0 comment('0新建  1已删除') TINYINT" json:"del_flag"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP" json:"update_time"`
}

func (m *User) TableComment() string {
	return "user"
}

type Vr struct {
	Vid        uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT" json:"vid"`
	Type       int       `xorm:"not null default 0 index(IDX_create) index(IDX_list) index(IDX_pos) TINYINT" json:"type"`
	XPos       int64     `xorm:"not null default 0 comment('景点坐标x') index(IDX_pos) BIGINT" json:"x_pos"`
	YPos       int64     `xorm:"not null default 0 comment('景点坐标y') index(IDX_pos) BIGINT" json:"y_pos"`
	Content    string    `xorm:"comment('消息内容URL') TEXT" json:"content"`
	DelFlag    int       `xorm:"not null default 0 comment('0新建  1已删除') index(IDX_create) index(IDX_list) TINYINT" json:"del_flag"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index(IDX_create) index(IDX_list) TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP" json:"update_time"`
}

func (m *Vr) TableComment() string {
	return "vr"
}

type Vrtype struct {
	Tid        uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT" json:"tid"`
	Type       int       `xorm:"not null default 0 comment('分类类型') index(IDX_create) index(IDX_list) TINYINT" json:"type"`
	Title      string    `xorm:"comment('分类名称') TEXT" json:"title"`
	DelFlag    int       `xorm:"not null default 0 comment('0新建  1已删除') index(IDX_create) index(IDX_list) TINYINT" json:"del_flag"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index(IDX_create) index(IDX_list) TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP" json:"update_time"`
}

func (m *Vrtype) TableComment() string {
	return "vrtype"
}
