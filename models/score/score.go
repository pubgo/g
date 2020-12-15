package score

import (
	"github.com/pubgo/x/models"
)

type Score struct {
	models.BaseModel

	ScoreID   string `json:"score_id"`   // 积分编号
	ScoreType int    `json:"score_type"` // 积分类型：1=登录 2=发布 3=评论 4=抽奖
	Amount    int    `json:"amount"`     // 积分数量
	InUid     string `json:"in_uid"`     // 用户内部编号
	RecID     string `json:"rec_id"`     // 记录编号
}
