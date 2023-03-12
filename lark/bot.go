package lark

import (
	"context"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"open-ai-feishu/g"
)

var client *lark.Client

// Init 在g.parse()之后初始化
func Init() {
	client = lark.NewClient(g.GetConfig().Lark.AppId, g.GetConfig().Lark.AppSecret)
}

type BotResp struct {
	ActivateStatus int           `json:"activate_status"`
	AppName        string        `json:"app_name"`
	AvatarUrl      string        `json:"avatar_url"`
	IpWhiteList    []interface{} `json:"ip_white_list"`
	OpenId         string        `json:"open_id"`
}

type Resp struct {
	Bot struct {
		ActivateStatus int           `json:"activate_status"`
		AppName        string        `json:"app_name"`
		AvatarUrl      string        `json:"avatar_url"`
		IpWhiteList    []interface{} `json:"ip_white_list"`
		OpenId         string        `json:"open_id"`
	} `json:"bot"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func GetBotOpenId() string {
	resp, err := client.Get(context.Background(), "/open-apis/bot/v3/info", nil, larkcore.AccessTokenTypeApp)
	if err != nil {
		return ""
	}
	var botResp Resp

	err = resp.JSONUnmarshalBody(&botResp, &larkcore.Config{Serializable: &larkcore.DefaultSerialization{}})
	return botResp.Bot.OpenId
}
