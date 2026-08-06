// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysQueueConfig is the golang structure of table xy_sys_queue_config for DAO operations like Where/Data.
type SysQueueConfig struct {
	g.Meta        `orm:"table:xy_sys_queue_config, do:true"`
	Id            any // ID
	Topic         any // Topic 标识（唯一）
	Title         any // 显示名称
	Workers       any // 并行 Worker 数
	MaxRetry      any // 最大重试次数
	RetryDelaySec any // 重试间隔（秒，0=立即重试）
	Status        any // 状态:0禁用,1启用
	Remark        any // 备注
	Sort          any // 排序
	CreatedAt     any // 创建时间
	UpdatedAt     any // 更新时间
}
