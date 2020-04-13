package relation

type ArticleImage struct {
	ID            int    `json:"id"`              // 自动编号
	Version       int    `json:"version"`         // 数据库版本
	CreatedAt     int    `json:"created_at"`      // 创建时间
	UpdatedAt     int    `json:"updated_at"`      // 更新时间
	DeleteAt      int    `json:"delete_at"`       // 删除时间
	Status        int    `json:"status"`          // 记录状态: -1=删除 0=可正常使用
	RecID         string `json:"rec_id"`          // 记录编号
	ArticleID     string `json:"article_id"`      // 文章编号
	ImageID       string `json:"image_id"`        // 图片编号
	ImageIndex    int    `json:"image_index"`     // 图片所在文章中的顺序
	Signature     string `json:"signature"`       // 数据签名
	DtcpId        string `json:"dtcp_id"`         // DTCP id
	DtcpDna       string `json:"dtcp_dna"`        // DTCP dna
	DtcpParentDna string `json:"dtcp_parent_dna"` // DTCP Parent dna
}

type ArticleImageRep struct {
	ID         int    `json:"id"`
	CreatedAt  int    `json:"created_at"`
	UpdatedAt  int    `json:"updated_at"`
	RecID      string `json:"rec_id"`
	ArticleID  string `json:"article_id"`
	ImageID    string `json:"image_id"`
	ImageIndex int    `json:"image_index"`
	//ImagesInfo []*image_mod.ImageRep `json:"images_info" gorm:"ForeignKey:ArticleID;AssociationForeignKey:ArticleID`
}

func (ai ArticleImageRep) TableName() string {
	return "article_images"
}
