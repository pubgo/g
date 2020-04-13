package api

import (
	"github.com/pubgo/g/xservice/sso/app"
	"github.com/pubgo/g/xservice/sso/model"
	"github.com/pubgo/g/xservice/sso/utils"
	"net/http"
	"strconv"

	l4g "github.com/alecthomas/log4go"

)

func InitSystem() {
	l4g.Debug(utils.T("api.system.init.debug"))

	BaseRoutes.ApiRoot.Handle("/config/client", ApiHandler(getClientConfig)).Methods("GET")
}

func getClientConfig(c *Context, w http.ResponseWriter, r *http.Request) {
	format := r.URL.Query().Get("format")

	if format == "" {
		c.Err = model.NewAppError("getClientConfig",
			"api.config.client.old_format.app_error", nil, "", http.StatusNotImplemented,
		)
		return
	}

	if format != "old" {
		c.SetInvalidParam("format")
		return
	}

	respCfg := map[string]string{}
	for k, v := range utils.ClientCfg {
		respCfg[k] = v
	}

	respCfg["NoAccounts"] = strconv.FormatBool(app.IsFirstUserAccount())
	RenderJson(w, respCfg)
}
