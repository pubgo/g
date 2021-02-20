package mark

import "github.com/pubgo/x/models"

type Mark struct {
	models.BaseModel

	MarkID             string `json:"mark_id"`              // 标注编号
	InUid              string `json:"in_uid"`               // 用户内部编号(内部流转)
	LinkID             string `json:"link_id"`              // 链接id
	ArticleID          string `json:"article_id"`           // 文章id
	CommentID          string `json:"comment_id"`           // 评论id
	GroupID            string `json:"group_id"`             // group id
	MarkContent        string `json:"mark_content"`         // 标注信息
	Collapsed          int    `json:"collapsed"`            // 光标是否合并: -1=删除 0=否 1=是
	StartContainerPath string `json:"start_container_path"` // 标注开始路径
	EndContainerPath   string `json:"end_container_path"`   // 标注结束路径
	StartOffset        int    `json:"start_offset"`         // 标注开始位
	EndOffset          int    `json:"end_offset"`           // 标注结束位
	Signature          string `json:"signature"`            // 数据签名
	Note               string `json:"note"`                 //
	Extra              string `json:"extra"`                //
}
