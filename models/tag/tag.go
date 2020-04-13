package tag

type Tag struct {
	ID   uint64 `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`
	Name string `json:"name"` // unique_index hash index
}
