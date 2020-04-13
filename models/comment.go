package models

// user comment some resource some context
type Comment struct {
	ID        int64  `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
	Namespace string `json:"category"` // 目录命名
	UserId    string                   // 评论者

	// 评论的对象
	ResType string
	ResID   string

	// 所有的评论都要带上这个ID
	GID       uint64 // hash(ResType,ResID)
	Format    string // image file url 物理格式
	Language  string // zh-CN ...
	SourceUrl string // url
}
