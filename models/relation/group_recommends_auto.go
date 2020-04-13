package relation

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type GroupRecommend struct {
	ID        int    `json:"id"`         // 自动编号
	CreatedAt int    `json:"created_at"` // 创建时间
	UpdatedAt int    `json:"updated_at"` // 更新时间
	Status    int    `json:"status"`     // 记录状态:  0=可正常使用 1=取消推荐
	RecID     string `json:"rec_id"`     // 记录编号
	GroupID   string `json:"group_id"`   // 圈子编号
	InUid     string `json:"in_uid"`     //  用户内部编号(内部流转)
}

var GroupRecommendReal = &GroupRecommend{}

func GetGroupRecommendByWheres(wheres map[string]interface{}) (*GroupRecommend, error) {
	if len(wheres) == 0 {
		return nil, errors.New("param[wheres] length is zero")
	}

	var object GroupRecommend
	var resultObjs []GroupRecommend
	err := mysql.SearchObject(&object, wheres, &resultObjs)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	if len(resultObjs) == 0 {
		return nil, nil
	}

	resultObj := resultObjs[0]
	if err != nil {
		return nil, err
	}

	return &resultObj, nil
}

func GetGroupRecommendMultiByWheres(wheres map[string]interface{}) ([]GroupRecommend, error) {
	if len(wheres) == 0 {
		return nil, errors.New("param[wheres] length is zero")
	}

	var object GroupRecommend
	var resultObjs []GroupRecommend
	err := mysql.SearchObject(&object, wheres, &resultObjs)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return resultObjs, nil
}

func GetGroupRecommendMultiByInWheresLimit(wheres, ins map[string]interface{}, orders string,
	limit, offset int) ([]GroupRecommend, error) {

	if limit < 0 {
		return nil, errors.New("param[limit] less than zero")
	}

	if offset < 0 {
		return nil, errors.New("param[offset] less than zero")
	}

	var object GroupRecommend
	var resultObjs []GroupRecommend

	err := mysql.SearchObjectByOrder(&object, wheres, ins, orders, limit, offset, &resultObjs)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	if len(resultObjs) == 0 {
		return nil, nil
	}

	return resultObjs, nil
}

func CreateGroupRecommend(value *GroupRecommend) error {
	if value == nil {
		return errors.New("param[value] is nil")
	}

	if value.RecID == "" {
		return errors.New("value field[groupDna] is empty")
	}

	err := mysql.CreateObject(value)
	if err != nil {
		return err
	}

	return nil
}

func UpdateGroupRecommend(wheres, updates map[string]interface{}) error {
	if len(wheres) == 0 {
		return errors.New("param[wheres] length is zero")
	}
	if len(updates) == 0 {
		return errors.New("param[updates] length is zero")
	}

	var object GroupRecommend
	err := mysql.UpdateObject(&object, wheres, updates)
	if err != nil {
		return err
	}

	return nil
}
