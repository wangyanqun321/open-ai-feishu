package lark

import (
	"context"
	"encoding/json"
	"fmt"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func reply(msgId, userId, content string) {
	dataContent := make(map[string]string)
	dataContent["text"] = fmt.Sprintf("<at user_id=\"%s\"></at> %s", userId, content)
	text, _ := json.Marshal(dataContent)
	textStr := string(text)
	var msgType = larkim.MsgTypeText
	req := larkim.NewReplyMessageReqBuilder().MessageId(msgId).Body(&larkim.ReplyMessageReqBody{
		Content: &textStr,
		MsgType: &msgType,
	}).Build()

	var resp *larkim.ReplyMessageResp
	var err error
	if resp, err = client.Im.Message.Reply(context.Background(), req); err != nil {
		fmt.Println(err)
		return
	}
	if !resp.Success() {
		fmt.Println(resp.StatusCode, resp.Err)
	}
}
