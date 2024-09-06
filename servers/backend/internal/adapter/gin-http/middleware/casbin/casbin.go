package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
)

type Casbin struct {
	response response.IResponse

	SyncedEnforcer *casbin.SyncedEnforcer
}

func NewCasbin(response response.IResponse, syncedEnforcer *casbin.SyncedEnforcer) *Casbin {
	return &Casbin{
		response:       response,
		SyncedEnforcer: syncedEnforcer,
	}
}

func (cb *Casbin) Middleware(c *gin.Context) {
	authority := c.GetString("authority_id")

	// 請求路徑
	path := c.Request.URL.Path
	// 請求方法
	method := c.Request.Method
	// 角色
	sub := authority

	// 檢查策略
	loadPolicyErr := cb.SyncedEnforcer.LoadPolicy()
	if loadPolicyErr != nil {
		cb.response.AuthFail(c, "load policy error")
		c.Abort()
		return
	}

	// 檢查權限
	success, enforceError := cb.SyncedEnforcer.Enforce(sub, path, method)
	if enforceError != nil {
		cb.response.AuthFail(c, "enforce error")
		c.Abort()
		return
	}

	if !success {
		cb.response.AuthFail(c, "permission denied")
		c.Abort()
		return
	}

	c.Next()
}
