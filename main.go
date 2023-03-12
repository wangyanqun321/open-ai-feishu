package main

import (
	"open-ai-feishu/g"
	"open-ai-feishu/lark"
)

func main() {
	g.Parse()
	lark.StartEventDispatcher()
}
