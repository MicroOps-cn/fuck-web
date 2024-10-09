package endpoint

import (
	"context"

	"github.com/MicroOps-cn/fuck/conv"
	"github.com/go-kit/kit/endpoint"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/service"
)

func MakeGetSecurityConfigEndpoint(_ service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		resp := SimpleResponseWrapper[*config.RuntimeSecurityConfig]{}
		resp.Data = config.GetRuntimeConfig().Security
		return resp, nil
	}
}

func MakePatchSecurityConfigEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Requester).GetRequestData().(*PatchSecurityConfigRequest)
		resp := SimpleResponseWrapper[interface{}]{}
		config.SetRuntimeConfig(func(c *config.RuntimeConfig) {
			dst := map[string]interface{}{}
			if resp.Error = conv.JSON(req, &dst); resp.Error != nil {
				return
			}
			if resp.Error = svc.PatchSystemConfig(ctx, "security", dst); resp.Error != nil {
				return
			}
			if c.Security == nil {
				c.Security = &config.RuntimeSecurityConfig{}
			}
			resp.Error = conv.JSON(req, c.Security)
		})
		return resp, nil
	}
}
