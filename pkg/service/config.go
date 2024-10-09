package service

import "context"

func (s Set) PatchSystemConfig(ctx context.Context, prefix string, patch map[string]interface{}) error {
	return s.commonService.PatchSystemConfig(ctx, prefix, patch)
}
