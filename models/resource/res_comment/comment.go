package comment

import "github.com/pubgo/g/models"

type Comment struct {
	models.BaseModel

	CommentID int `json:"comment_id"` // 评论编号
	ResId     string
	ResType   string
	Category  string
	GroupID   string `json:"group_id"` // 圈子编号
}
