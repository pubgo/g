package notify

import "github.com/pubgo/x/models"

type Notify struct {
	models.BaseModel

	NotifyID string `json:"notify_id"` // 消息ID
	Sender   string `json:"sender"`    // 发送方
	Receiver string `json:"receiver"`  // 接收方
	Type     int    `json:"type"`      // 类型：0子转让通知
	Resource string `json:"resource"`  // 资源

	MessageInfo  interface{} `json:"message_info"`
	ArticleCount int         `json:"article_count,omitempty"`
}

const (
	NotifyComment                   = 0  // 评论通知
	NotifyTransferGroup             = 1  // 圈子转让通知
	NotifyCommentMention            = 2  // 评论中@
	NotifyArticleMention            = 3  // 文章中@
	NotifySubComment                = 4  // 子评论
	NotifyCommentReply              = 5  // 评论回复
	NotifyInviteJoinGroup           = 6  // 邀请消息
	NotifyJoinGroup                 = 7  // 加入通知
	NotifyInviteJoinGroupNeedVerify = 8  // 邀请加入验证通知
	NotifyBoost                     = 9  // boost
	NotifyArticlePush               = 10 // 文章推送
	NotifyReward                    = 11 // 打赏
)
