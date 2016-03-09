/**
 * Copyright 2015 @ z3q.net.
 * name : default_c.go
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package ols

import (
	"github.com/jsix/gof/web/session"
	"go2o/src/cache"
	"go2o/src/core/domain/interface/member"
	"go2o/src/core/domain/interface/partner"
	"go2o/src/core/service/dps"
	"go2o/src/x/echox"
	"net/http"
	"net/url"
)

// 获取商户编号
func GetSessionPartnerId(ctx *echox.Context) int {
	return ctx.Get("partner_id").(int)
}

func getPartner(ctx *echox.Context) *partner.ValuePartner {
	partnerId := ctx.Get("partner_id").(int)
	return cache.GetValuePartnerCache(partnerId)
}

// 获取商户API信息
func getPartnerApi(ctx *echox.Context) *partner.ApiInfo {
	return dps.PartnerService.GetApiInfo(GetSessionPartnerId(ctx))
}

// 获取商户站点设置
func getSiteConf(ctx *echox.Context) *partner.SiteConf {
	return cache.GetPartnerSiteConf(GetSessionPartnerId(ctx))
}

// 获取会员
func GetMember(ctx *echox.Context) *member.ValueMember {
	memberIdObj := ctx.Session.Get("member")
	if memberIdObj != nil {
		if o, ok := memberIdObj.(*member.ValueMember); ok {
			return o
		}
	}
	return nil
}

// 检查会员是否登陆
func CheckMemberLogin(ctx *echox.Context) bool {
	if ctx.Session.Get("member") == nil {
		ctx.Response().Header().Add("Location", "/user/login?return_url="+
			url.QueryEscape(ctx.Request().RequestURI))
		ctx.Response().WriteHeader(302)
		return false
	}
	return true
}

// 获取商户编号
func GetPartnerId(r *http.Request, s *session.Session) int {
	return 104
	currHost := r.Host
	host := s.Get("webui_host")
	pid := s.Get("webui_pid")
	if host == nil || pid == nil || host != currHost {
		partnerId := dps.PartnerService.GetPartnerIdByHost(currHost)
		if partnerId != -1 {
			s.Set("webui_host", currHost)
			s.Set("webui_pid", partnerId)
			s.Save()
		}
		return partnerId
	}
	return pid.(int)
}