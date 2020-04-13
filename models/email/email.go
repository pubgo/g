package email

type Email struct {
	ID            int    `json:"id"`             // 自动编号
	Version       int    `json:"version"`        // 数据库版本
	CreatedAt     int    `json:"created_at"`     // 创建时间
	UpdatedAt     int    `json:"updated_at"`     // 更新时间
	DeleteAt      int    `json:"delete_at"`      // 删除时间
	Status        int    `json:"status"`         // 记录状态: -1=删除 0=发送中 1=发送成功 2=发送失败
	RecID         string `json:"rec_id"`         // 日志编号
	Email         string `json:"email"`          // 用户email
	SendContent   string `json:"send_content"`   // 发送内容
	FailedResult  string `json:"failed_result"`  // 失败结果描述
	ThirdEmailId  string `json:"third_email_id"` // 第三方查询id
	ThirdProvider string `json:"third_provider"` // 第三方名称
}
