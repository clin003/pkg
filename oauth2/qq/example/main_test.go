package main

import (
	"testing"

	. "gitee.com/lyhuilin/pkg/oauth2"
	"gitee.com/lyhuilin/pkg/oauth2/qq"
)

func TestQqAuth(t *testing.T) {
	qqConf := &OAuthConfg{
		ClientID:     "xxx",
		ClientSecret: "xxx",
		RedirectURL:  "https://auth.nav.xin/oauth/qq",
		Scopes:       []string{"get_user_info", "getUnionId"},
		AuthURL:      "https://graph.qq.com/oauth2.0/authorize",
		TokenURL:     "https://graph.qq.com/oauth2.0/token",
	}
	qqAuth := qq.NewAuth(qqConf)
	t.Log(qqAuth.AuthCodeURL("state")) //获取第三方登录地址
	code := "xxx"

	// 获取accessToken相关信息
	if qqToken, err := qqAuth.Token(code); err != nil {
		t.Log(err)
	} else {
		t.Log("qqToken", qqToken)
		t.Log("AccessToken", qqToken.AccessToken)

		// 获取openInfo信息
		if qqOpenIDInfo, err := qqAuth.OpenInfo(qqToken.AccessToken, true); err != nil {
			t.Log(err)
		} else {
			t.Log("qqOpenIDInfo", qqOpenIDInfo)
			t.Log("OpenID", qqOpenIDInfo.OpenID)
			t.Log("UnionID", qqOpenIDInfo.UnionID)

			// 获取userinfo 信息
			if qqUserInfo, err := qqAuth.UserInfo(qqToken.AccessToken, qqOpenIDInfo.OpenID); err != nil {
				t.Log(err)
			} else if qqUserInfo.Ret < 0 {
				t.Log("qqUserInfo", qqUserInfo)
				t.Log(qqUserInfo.Msg)
			} else {
				t.Log("qqUserInfo", qqUserInfo)
				t.Log(qqUserInfo.NickName)
				t.Log(qqUserInfo.Figureurl_qq_1)
			}

		}
	}

}
