package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/admin"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/user"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	User    user.RouterGroup
	Admin   admin.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
