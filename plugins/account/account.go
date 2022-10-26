package account

import (
	"fmt"
	plugin "github.com/itering/subscan-plugin"
	"github.com/itering/subscan-plugin/router"
	"github.com/itering/subscan-plugin/storage"
	"github.com/itering/subscan/plugins/account/dao"
	"github.com/itering/subscan/plugins/account/http"
	"github.com/itering/subscan/plugins/account/model"
	"github.com/itering/subscan/plugins/balance/service"
	"github.com/itering/subscan/util"
	"github.com/shopspring/decimal"
	"strings"
)

var srv *service.Service

type Account struct {
	d storage.Dao
}

func New() *Account {
	return &Account{}
}

func (a *Account) InitDao(d storage.Dao) {
	srv = service.New(d)
	a.d = d
	a.Migrate()
}

func (a *Account) InitHttp() []router.Http {
	return http.Router(srv)
}

func (a *Account) ProcessExtrinsic(*storage.Block, *storage.Extrinsic, []storage.Event) error {
	return nil
}

func (a *Account) ProcessEvent(block *storage.Block, event *storage.Event, fee decimal.Decimal) error {
	if event == nil {
		return nil
	}
	var paramEvent []storage.EventParam
	util.UnmarshalAny(&paramEvent, event.Params)

	switch fmt.Sprintf("%s-%s", strings.ToLower(event.ModuleId), strings.ToLower(event.EventId)) {
	case strings.ToLower("System-NewAccount"):
		return dao.NewAccount(a.d, util.ToString(paramEvent[0].Value))
	}

	return nil
}

func (a *Account) SubscribeExtrinsic() []string {
	return nil
}

func (a *Account) SubscribeEvent() []string {
	return []string{"system"}
}

func (a *Account) Version() string {
	return "0.1"
}

func (a *Account) UiConf() *plugin.UiConfig {
	conf := new(plugin.UiConfig)
	conf.Init()
	conf.Body.Api.Method = "post"
	conf.Body.Api.Url = "api/plugin/account/accounts"
	conf.Body.Api.Adaptor = fmt.Sprintf(conf.Body.Api.Adaptor, "list")
	conf.Body.Columns = []plugin.UiColumns{
		{Name: "extrinsic_id", Label: "Id"},
		{Name: "block", Label: "Block"},
		{Name: "time", Label: "Time"},
		{Name: "from", Label: "From"},
		{Name: "to", Label: "To"},
		{Name: "value", Label: "Value"},
		{Name: "result", Label: "Result"},
	}
	return conf
}

func (a *Account) Migrate() {
	_ = a.d.AutoMigration(&model.AccountHistory{})
	_ = a.d.AddUniqueIndex(&model.AccountHistory{}, "address", "address")
}
