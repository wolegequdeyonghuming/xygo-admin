// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysAttachment is the golang structure of table xy_sys_attachment for DAO operations like Where/Data.
type SysAttachment struct {
	g.Meta     `orm:"table:xy_sys_attachment, do:true"`
	Id         any // ID
	Topic      any // 分组/主题标识
	UserId     any // 上传用户ID（预留）
	Url        any // 物理路径（相对路径）
	Width      any // 宽度
	Height     any // 高度
	Name       any // 原始名称
	Size       any // 大小（字节）
	Mimetype   any // MIME类型
	Quote      any // 上传(引用)次数
	Storage    any // 存储方式：local=本地
	Sha1       any // sha1摘要
	CreateTime any // 创建时间
	UpdateTime any // 更新时间
	TenantId   any // 所属租户ID（0=平台）
}
