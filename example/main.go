package main

import (
	_ "github.com/SupenBysz/gf-admin-community"

	"github.com/SupenBysz/gf-admin-community/example/internal/boot"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	boot.Main.Run(gctx.New())
}
