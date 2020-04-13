package report

import "github.com/pubgo/g/models"

// 举报
type Report struct {
	models.BaseModel

	ReportID     string `json:"report_id"`      // 举报编号
	InUid        string `json:"in_uid"`         // 举报人ID
	InOutUid     string `json:"in_out_uid"`     // 举报人OutID
	ReportUid    string `json:"report_uid"`     // 被举报人ID
	ReportOutUid string `json:"report_out_uid"` // 被举报人OutID
	ReportType   int    `json:"report_type"`    // 举报类型 1=圈子类型 2=文章类型 3=评论类型 4=用户
	ResourceID   string `json:"resource_id"`    // 举报资源ID 参照report_type为 圈子ID 文章ID 评论ID 用户InUid
	ReasonID     int    `json:"reason_id"`      // 举报原因ID,详情参照 举报原因列表
	Note         string `json:"note"`           // 举报原因说明
	Extra        string `json:"extra"`          // 额外信息 如图片，文件等附件，应为json结构
	Signature    string `json:"signature"`      // 数据签名
	OutUid       string `json:"out_uid"`        // 举报人ID
	Reason       string `json:"reason"`         // 举报原因ID,详情参照 举报原因列表
}
