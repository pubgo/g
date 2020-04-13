package menu

// menu需要构建的点有点多 这里关联关系表直接把所有数据拿过来 用代码实现关联  后期实现主外键模式
type Menu struct {
	BaseMenu

	MenuId      string `json:"menuId"`
	AuthorityId string `json:"-"`
	Children    []Menu `json:"children"`
}

type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

// 为角色增加menu树
func (m *Menu) AddMenuAuthority(menus []BaseMenu, authorityId string) (err error) {
	return nil
}

// 查看当前角色树
func (m *Menu) GetMenuAuthority(authorityId string) (err error, menus []Menu) {
	return err, menus
}

//获取动态路由树
func (m *Menu) GetMenuTree(authorityId string) (err error, menus []Menu) {
	return err, menus
}

func getChildrenList(menu *Menu) (err error) {
	return err
}
