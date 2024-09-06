package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const (
	casbinText = `
[request_definition]
    r = sub, obj, act

[policy_definition]
    p = sub, obj, act

[role_definition]
    g = _, _

[policy_effect]
    e = some(where (p.eft == allow))

[matchers]
    m = g(r.sub, p.sub) == true \
        && keyMatch2(r.obj, p.obj) == true \
        && regexMatch(r.act, p.act) == true \
        || r.sub == "root"
`
)

type Option struct {
	Prefix string
}

func NewOption(conf *viper.Viper) Option {
	return Option{
		Prefix: conf.GetString("casbin.prefix"),
	}
}

func NewCasbin(gorm *gorm.DB, option Option) *casbin.SyncedEnforcer {
	adapter, adapterErr := gormadapter.NewAdapterByDBUseTableName(gorm, option.Prefix, "casbin")
	if adapterErr != nil {
		panic(adapterErr)
	}

	casbinModel, err := model.NewModelFromString(casbinText)
	if err != nil {
		panic(err)
	}

	syncedCachedEnforcer, enforcerErr := casbin.NewSyncedEnforcer(casbinModel, adapter)
	if enforcerErr != nil {
		panic(enforcerErr)
	}

	loadPolicyErr := syncedCachedEnforcer.LoadPolicy()
	if loadPolicyErr != nil {
		panic(loadPolicyErr)
	}

	return syncedCachedEnforcer
}
