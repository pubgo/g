package file

import "github.com/pubgo/x/models"

type File struct {
	models.BaseModel

	FileID        string `json:"file_id"`         // 文件id
	FileDna       string `json:"file_dna"`        // 文件dna
	FileParentDna string `json:"file_parent_dna"` // 文件父dna
	StorageUrl    string `json:"storage_url"`     // 文件存储全路径
	Category      string `json:"category"`        // 文件分类
	Ext           string `json:"ext"`             // 文件后缀名
	Size          int    `json:"size"`            // 文件大小
	Title         string `json:"title"`           // 文件名称
	ImageDesc     string `json:"image_desc"`      // 文件描述
	ContentHash   string `json:"content_hash"`    // 文件内容hash值
	License       string `json:"license"`         // 授权内容
	Signature     string `json:"signature"`       // 数据签名
}
