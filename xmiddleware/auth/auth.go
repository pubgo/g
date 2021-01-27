package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pubgo/x/xmiddleware"
)

// 网页上用cookie
// 移动端用token

const Client = "X-Client"
const TokenEntropy = 32

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端类型，网站，移动端，桌面客户端
		// 如果不符合，就认证失败
		//_client := xmiddleware.GetHeader(c, Client)
		// 获取cookie 中的组合认证
		// sessionID，跟登陆用户绑定，有时效和权限等
		// 网页的ID
		// 用户权限ID
		// 如果是移动端，就用token
		// 如果是web 那么，是谁

		reqToken := xmiddleware.GetHeader(c, "Authorization", "WWW-Authenticate", "Token")
		if reqToken == "" {
			c.AbortWithStatusJSON(401, gin.H{"data": "Authorization Error"})
			return
		}

		// 获取类型终端类型
		// 获取
		// 根据不同的终端类型获取cookie和Authorization
		// cookie中获取网页ID和sessionID
		// 如果认证失败，那么，从定向到从新登陆
		// 获取sessionID就传递下去，并获取
		// 获得token key
		// 然后获得session
		// token key 可以刷新session
		// session过期很快
		// token key过期慢
		// 用户名和密码不保存本地，token key可以保存
		// 用户名可以是，email，phone，name

	}
}
