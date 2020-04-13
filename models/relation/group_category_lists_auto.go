package relation

type GroupCategoryList struct {
	ID                int    `json:"id"`                  // 自动编号
	Version           int    `json:"version"`             // 数据库版本
	CreatedAt         int    `json:"created_at"`          // 创建时间
	UpdatedAt         int    `json:"updated_at"`          // 更新时间
	DeleteAt          int    `json:"delete_at"`           // 删除时间
	Status            int    `json:"status"`              // 记录状态: -1=删除 0=可正常使用
	CategoryID        string `json:"category_id"`         // 圈子类别编号
	CategoryDna       string `json:"category_dna"`        // 圈子类别dna
	CategoryParentDna string `json:"category_parent_dna"` // 圈子类别父dna
	NameZh            string `json:"name_zh"`             // 中文名称
	NameEn            string `json:"name_en"`             // 英文名称
	NameJa            string `json:"name_ja"`             // 日文名称
	NameKo            string `json:"name_ko"`             // 韩文名称
	Signature         string `json:"signature"`           // 数据签名
}
