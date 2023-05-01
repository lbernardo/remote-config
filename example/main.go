package main

import (
	"fmt"
	"github.com/lbernardo/remote-config/pkg/config"
	"github.com/lbernardo/remote-config/pkg/sync"
)

func main() {
	cfg := sync.New()
	cfg.SetEnvironment("dev").
		SetProject("teste").
		SetNamespace("default").
		SetBucket("lb-sync-remote")
	cfg.Sync()

	fmt.Println(config.GetString("name"))
}
