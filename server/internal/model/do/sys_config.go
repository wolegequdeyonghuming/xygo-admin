// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

// SysConfig is the golang structure of table xy_sys_config for DAO operations like Where/Data.
type SysConfig struct {
	g.Meta     `orm:"table:xy_sys_config, do:true"`
	Id         any         // 主键
	Group      any         // 分组标识，如 basics/mail/oss
	GroupName  any         // 分组名称
	Name       any         // 配置项显示名
	Key        any         // 配置键（唯一）
	Value      any         // 配置值（字符串/JSON）
	Type       any         // 控件类型:text/textarea/number/switch/select/radio/checkbox/json/object/array/color/upload
	Options    *gjson.Json // 组件参数/选项 JSON
	Rules      *gjson.Json // 校验规则 JSON
	Sort       any         // 排序(大靠后)
	Remark     any         // 备注
	AllowDel   any         // 允许删除:0=否,1=是
	CreatedBy  any         // 创建人
	UpdatedBy  any         // 更新人
	CreateTime any         // 创建时间
	UpdateTime any         // 更新时间
	TestField  any         // 测试字段（v1.2.7）
}
