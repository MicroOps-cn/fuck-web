package endpoint

import (
	"context"
	"net/url"

	w "github.com/MicroOps-cn/fuck/wrapper"
	"github.com/go-kit/kit/endpoint"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/global"
	"github.com/MicroOps-cn/fuck-web/pkg/service"
)

func MakeGetGlobalConfigEndpoint(_ service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		adminPrefix := ctx.Value(global.HTTPWebPrefixKey).(string)
		externalURL := ctx.Value(global.HTTPExternalURLKey).(string)

		globalConfig := config.Get().GetGlobal()
		securityConfig := config.Get().GetSecurity()
		resp := &GlobalConfig{
			Title:            globalConfig.Title,
			SubTitle:         globalConfig.SubTitle,
			Logo:             globalConfig.Logo,
			Copyright:        globalConfig.Copyright,
			DefaultLoginType: LoginType(securityConfig.DefaultLoginType),
			Version:          w.Version,
			ExternalUrl:      externalURL,
			AdminUrl:         w.M(url.JoinPath(externalURL, adminPrefix)),
		}
		oauth2 := securityConfig.Oauth2
		if !securityConfig.DisableLoginForm {
			resp.LoginType = append(resp.LoginType, &GlobalLoginType{Type: LoginType_normal})
			if !config.GetRuntimeConfig().GetSecurity().ForceEnableMfa {
				resp.LoginType = append(resp.LoginType, &GlobalLoginType{Type: LoginType_email})
			}
		}
		if len(securityConfig.AllowLoginType) > 0 {
			w.Filter(resp.LoginType, func(item *GlobalLoginType) bool {
				for _, loginType := range securityConfig.AllowLoginType {
					if loginType.String() == item.Type.String() {
						return true
					}
				}
				return false
			})
		}
		for _, options := range oauth2 {
			resp.LoginType = append(resp.LoginType, &GlobalLoginType{
				Id:        options.Id,
				Type:      LoginType_oauth2,
				Name:      options.Name,
				Icon:      options.Icon,
				AutoLogin: options.AutoLogin,
			})
		}
		return resp, nil
	}
}
