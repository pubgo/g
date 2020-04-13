package models

// Resource model
type Resource struct {
	Status      int8   `json:"status"`  // 记录状态: -1=删除 0=可正常使用
	Version     int8   `json:"version"` // 版本
	ID          int64  `json:"id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
	Slug        string `json:"slug"`      // unique_index 人类可读的ID
	Namespace   string `json:"namespace"` // 目录命名
	Title       string `json:"title"`     // 文件名称
	CoverUrl    string                    // url 封面
	Abstract    string `json:"abstract"`  // 文章摘要
	Description string                    // 描述
	//Type        string                 // 实际类型 逻辑类型格式 比如log 通讯录 文章等
	Format    string // image file url 物理格式
	Language  string // zh-CN ...
	SourceUrl string // url
	Creator   string // 创建者
	Tags      string
	Category  string

	// 地理位置
	Geo string

	// 来源链接
	OriginUrl string

	PublishTime int    `json:"publish_time"` // 文章的发布时间
	Publisher   string `json:"publisher"`    // 文章作者
	Group       string
	ContentHash string `json:"content_hash"` // 文章的hash值

	// 额外的数据
	Extra string

	//YuanChuang bool
	// 是否是原创

	// 是否公开
	// 允许订阅
	// 允许评论
	// 允许申请加入
	// 允许打赏
}
