package relation


type GroupMemberRep struct {
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	Status    int    `json:"status"`
	GroupID   string `json:"group_id"`
	IsManage  int    `json:"is_manage"`
}

type GroupMemberListRep struct {
	ID        int                 `json:"id"`         // 自动编号
	Version   int                 `json:"version"`    // 数据库版本
	CreatedAt int                 `json:"created_at"` // 创建时间
	UpdatedAt int                 `json:"updated_at"` // 更新时间
	DeleteAt  int                 `json:"delete_at"`  // 删除时间
	Status    int                 `json:"status"`     // 记录状态: -1=删除 0=可正常使用 1=申请加入 2=拒绝加入
	RecID     string              `json:"rec_id"`     // 记录编号
	GroupID   string              `json:"group_id"`   // 圈子编号
	InUid     string              `json:"in_uid"`     // 圈子成员内部编号(内部流转)
	Signature string              `json:"signature"`  // 数据签名
	IsManage  int                 `json:"is_manage"`  // 是否是管理员,1=管理员
	Member    UserRep `json:"member" gorm:"ForeignKey:InUid;AssociationForeignKey:in_uid"`
}

func (member *GroupMemberRep) TableName() string {
	return "group_members"
}

func (member *GroupMemberListRep) TableName() string {
	return "group_members"
}
