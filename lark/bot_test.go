package lark

import (
	"fmt"
	"open-ai/g"
	"testing"
)

func TestGetBotOpenId(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "", want: "ou_8731bb5d33872fe7cb06b9f5f1a917f0"},
	}
	g.ParseWithPath("../cfg.yaml")
	Init()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetBotOpenId()
			fmt.Println(got)
			if got != tt.want {
				t.Errorf("GetBotOpenId() = %v, want %v", got, tt.want)
			}
		})
	}
}
