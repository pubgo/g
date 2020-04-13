package image

type ImageRep struct {
	ImageID    string `json:"image_id"`    // 图片id
	StorageUrl string `json:"storage_url"` // 图片存储全路径
	Height     int    `json:"height"`      // 高度
	Width      int    `json:"width"`       // 宽度
	Category   string `json:"category"`    // 图片分类
	Ext        string `json:"ext"`         // 图片后缀名
	Size       int    `json:"size"`        // 图片大小
	Title      string `json:"title"`       // 图片标题
	ImageDesc  string `json:"image_desc"`  // 图片描述
	ArticleID  string `json:"-"`           // article id
}

type Image struct {
	ID             int    `json:"id"`               // 自动编号
	Version        int    `json:"version"`          // 数据库版本
	CreatedAt      int    `json:"created_at"`       // 创建时间
	UpdatedAt      int    `json:"updated_at"`       // 更新时间
	DeleteAt       int    `json:"delete_at"`        // 删除时间
	Status         int    `json:"status"`           // 记录状态: -1=删除 0=可正常使用
	ImageID        string `json:"image_id"`         // 图片id
	ImageDna       string `json:"image_dna"`        // 图片dna
	ImageParentDna string `json:"image_parent_dna"` // 图片父dna
	StorageUrl     string `json:"storage_url"`      // 图片存储全路径
	Category       string `json:"category"`         // 图片分类
	Height         int    `json:"height"`           // 高度
	Width          int    `json:"width"`            // 宽度
	Ext            string `json:"ext"`              // 图片后缀名
	Size           int    `json:"size"`             // 图片大小
	Title          string `json:"title"`            // 图片标题
	ImageDesc      string `json:"image_desc"`       // 图片描述
	ContentHash    string `json:"content_hash"`     // 图片内容hash值
	License        string `json:"license"`          // 授权内容
	Signature      string `json:"signature"`        // 数据签名
	DtcpId         string `json:"dtcp_id"`          // DTCP id
	DtcpDna        string `json:"dtcp_dna"`         // DTCP dna
	DtcpParentDna  string `json:"dtcp_parent_dna"`  // DTCP Parent dna
}

type PicInfo struct {
	AssetFormat string `json:"asset_format"`
	Brand       string `json:"brand"`
	BrandId     int64  `json:"brandId"`
	Cameraman   string `json:"cameraman"`
	Caption     string `json:"caption"`
	Copyright   string `json:"copyright"`
	FileType    string `json:"file_type"`
	Height      int64  `json:"height"`
	ID          int64  `json:"id"`
	ImgDate     string `json:"img_date"`
	License     string `json:"license"`
	Oss176      string `json:"oss176"`
	Restrict    string `json:"restrict"`
	Size        string `json:"size"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Width       int64  `json:"width"`
}

type ImageMetadata struct {
	Name            string `json:"name"`
	ID              string `json:"ID"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	FileSize        int    `json:"file_size"`
	License         string `json:"license"`
	Format          string `json:"format"`
	FileType        string `json:"file_type"`
	Cameraman       string `json:"cameraman"`
	Copyright       string `json:"copyright"`
	URL             string `json:"url"`
	URL176          string `json:"url176"`
	Keywords        string `json:"keywords"`
	Size            string `json:"size"`
	ImgDate         string `json:"img_date"`
	Price           string `json:"price"`
	Title           string `json:"title"`
	ResID           string `json:"res_id"`
	Restrict        string `json:"restrict"`
	DHash           string `json:"dhash"`
	URLSource       string `json:"url_source"`
	ImageSource     string `json:"image_source"`
	ImageAuthor     string `json:"image_author"`
	ImageOrg        string `json:"image_org"`
	Brand           string `json:"brand"`
	DciID           string `json:"dci_id"`
	DciName         string `json:"dci_name"`
	DciAuthor       string `json:"dci_author"`
	DciAuthorized   string `json:"dci_authorized"`
	DciType         string `json:"dci_type"`
	DciPublishTime  string `json:"dci_publish_time"`
	DciCreateTime   string `json:"dci_create_time"`
	DciRegisterTime string `json:"dci_register_time"`
}
