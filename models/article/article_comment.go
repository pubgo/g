package article

import "github.com/pubgo/x/models"

type Comment struct {
	models.BaseModel

	CommentID      string   `json:"comment_id"`       // 评论编号
	InUid          string   `json:"in_uid"`           // 用户内部编号(内部流转)
	ToUid          string   `json:"to_uid"`           //回复对象
	ArticleID      string   `json:"article_id"`       // 文章编号
	GroupID        string   `json:"group_id"`         // 圈子编号
	MarkID         string   `json:"mark_id"`          // mark id
	CommentContent string   `json:"comment_content"`  // 用户发布评论
	ContentHash    string   `json:"content_hash"`     // 用户发布评论的hash值,可以重复
	SourceParentID string   `json:"source_parent_id"` // 父评论编号
	Signature      string   `json:"signature"`        // 数据签名
	CommentType    int      `json:"comment_type"`     // 评论类型 1=正常评论 2=boost
	Mentions       string   `json:"-"`                // 通知
	MentionsInfo   []string `json:"mentions"`
}
