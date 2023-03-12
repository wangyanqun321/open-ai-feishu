package lark

import (
	"context"
	"encoding/json"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/core/httpserverext"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"net/http"
	"open-ai-feishu/chatgpt"
	"open-ai-feishu/g"
	"strings"
)

func StartEventDispatcher() {
	// 初始化lark客户端
	Init()
	// 注册消息处理器
	handler := dispatcher.NewEventDispatcher(g.GetConfig().VerificationToken, g.GetConfig().EventEncryptKey).
		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
			go P2MessageReceiveV1Handler(event)
			return nil
		})

	// 注册 http 路由
	http.HandleFunc(g.GetConfig().Path, httpserverext.NewEventHandlerFunc(handler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))

	// 启动 http 服务
	err := http.ListenAndServe(g.GetConfig().Addr, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %v", err))
	}
}

func P2MessageReceiveV1Handler(event *larkim.P2MessageReceiveV1) {
	// fmt.Println(larkcore.Prettify(event))
	data := event.Event
	eventType := event.EventV2Base.Header.EventType
	if eventType != "im.message.receive_v1" {
		return
	}
	message := data.Message
	if message == nil {
		return
	}
	if *message.ChatType != "group" {
		return
	}
	if *message.MessageType != "text" {
		return
	}
	content := make(map[string]string)
	_ = json.Unmarshal([]byte(*message.Content), &content)
	text := content["text"]
	msgId := message.MessageId
	openId := data.Sender.SenderId.OpenId
	mentions := message.Mentions
	var hasBot bool
	for _, mention := range mentions {
		if *mention.Id.OpenId != "ou_8731bb5d33872fe7cb06b9f5f1a917f0" {
			continue
		}
		hasBot = true
		key := mention.Key
		text = strings.Replace(text, *key, "", 1)
	}
	// 只处理at机器人的消息
	if !hasBot {
		return
	}
	text = strings.TrimSpace(text)
	if text == "" {
		reply(*msgId, *openId, "请输入问题")
	} else {
		answer := chatgpt.Ask(text)
		reply(*msgId, *openId, answer)
	}
}
