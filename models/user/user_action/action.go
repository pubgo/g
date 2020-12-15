package user_action

import "github.com/pubgo/x/models"

// 用户对某一资源进行动作，以及所带有的数据
// 点赞, 评论, 转发, 收藏, 喜欢, 打赏, 标记, 订阅, 举报, 分享, ...
// 登陆
// 所有的动作都会被记录下来
type Action struct {
	models.BaseModel

	UserID  int64                    // 个人或者团体
	Action  int8                     // 点赞, 评论, 转发, 收藏, 喜欢, 打赏, 标记, 订阅, 举报, 分享, ...
	Payload string                   // 动作数据
	ResType string `json:"res_type"` // 资源类型
	ResID   int64  `json:"res_id"`   // 资源ID
}
