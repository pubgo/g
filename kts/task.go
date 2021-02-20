package kts

import (
	"encoding/json"
	"github.com/pubgo/xerror"
)

// Task xtask define
type Task struct {
	CreatedAt  uint   `json:"created_at,omitempty" db:"created_at" gorm:"index;not null"`                    // 每级任务生成时间
	FinishedAt uint   `json:"finished_at,omitempty" db:"finished_at" gorm:"index;not null"`                  // 每级任务完成时间
	Status     string `json:"status,omitempty" db:"status" gorm:"type:varchar(20);index;not null"`           // 任务状态
	Type       string `json:"type,omitempty" db:"type" gorm:"type:varchar(20);not null"`                     // 任务类型(article, image)
	CurService string `json:"cur_service"`                                                                   // 当前服务
	GID        string `json:"task_id,omitempty" db:"task_id" gorm:"type:varchar(100);unique_index;not null"` // 全局任务 ID, 通过该ID能够得到任务树, 得到所有相关的任务
	PID        string `json:"task_id,omitempty" db:"task_id" gorm:"type:varchar(100);unique_index;not null"` // 父任务 ID, 通过该ID可以获取子任务
	TID        string `json:"task_id,omitempty" db:"task_id" gorm:"type:varchar(100);unique_index;not null"` // 任务 ID, 本次任务的ID
	Input      string `json:"input,omitempty" db:"input" gorm:"type:text;not null"`                          // 任务参数
	Output     string `json:"output,omitempty" db:"output" gorm:"type:text;not null"`                        // 任务参数
	Priority   uint8  `json:"priority,omitempty" db:"priority" gorm:"not null"`                              // 任务优先度 1-9
	RetryNum   int    `json:"retry_num,omitempty" db:"retry_num" gorm:"not null"`                            // 任务重试次数
}

// Marshal xtask marshal
func (t Task) Marshal() []byte {
	return xerror.PanicErr(json.Marshal(t)).([]byte)
}
