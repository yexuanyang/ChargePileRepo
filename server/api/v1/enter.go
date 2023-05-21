package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/admin"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/user"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	UserApiGroup    user.ApiGroup
	AdminApiGroup   admin.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
