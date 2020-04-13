package models

import "time"

// 都柏林核心（Dublin Core）元数据
// DCMI 类型
// 资源集合Collection
// 数据集 DataSet
// 事件 Event
// 图像 Image
// 交互资源 InteractiveResource
// 动态图像 MovingImage
// 物理对象 PhysicalObject
// 服务 Service
// 软件 Software
// 声音 Sound
// 静态图像 StillImage
// 文本 Text

// Dublin Metadata base
type DMB struct {
	Title        string
	Creator      string
	Description  string
	Abstract     string `json:"abstract"` // 文章摘要
	Publisher    string
	Type         string
	Contributors []string
	Format       string
	Language     string
	License      string // 最好是一个url
	Source       string
	Subject      string `json:"subject"`

	GID     string
	PID     string
	ID      uint64
	Version int8

	Modified string
	Medium   string

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Status    int8      `gorm:"column:status" json:"status"` // 记录状态: -1=删除 0=可正常使用
}
