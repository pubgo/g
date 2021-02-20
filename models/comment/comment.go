package comment

import "github.com/pubgo/x/models"

type Comment struct {
	models.BaseModel

	CommentID   string `json:"comment_id"`   // 评论编号
	CommentType int    `json:"comment_type"` //

	UserId string // 评论者

	// 评论的对象
	ResType string
	ResID   string

	// 所有的评论都要带上这个ID
	GID string

	CommentContent string `json:"comment_content"` // 用户发布评论
}
