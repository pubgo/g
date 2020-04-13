package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

const activeName = "WeChat"

// 屏幕大小
var screenPoint = point(1440, 900)

// 最大化, 最小化
// robotgo.KeyTap("f", "lctrl", "lcmd")

// 公众号回退按钮
var gzhBackPoint = point(25, 50)

// 公众号前进按钮
var gzhGoPoint = point(65, 50)

// 公众号文章列表间隔
const gzhListStep = 105

// 微信公众号第一篇文章的位置
var gzhFirstList = point(715, 385)

// 聊天列表 187,95
var contactListPoint = point(187, 95)

// 聚焦通信列表 35 167
var contactListFocusPoint = point(35, 167)

// 公众号关闭 175 189
var gzhgbanPoint = point(175, 189)

// 公众号 第一篇 位置 597 293 竖列+10
var gzhgbanPoint1 = point(175, 189)

// 文章回退 185 223
// 文章刷新 253 222
// 复制网址 1006 223
// 文章正上点击 621 291
// data_ok
// url_ok
// 公众号返回 391 28
// 公众号列表1409 33 --  1362 290

func init() {
	title := robotgo.GetTitle()
	fmt.Println("title@@@ ", title)
}
