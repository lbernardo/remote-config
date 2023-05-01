# remote-config

Use to get  aws s3 remote config

## Install

```yaml
go get github.com/lbernardo/remote-config
```

## Usage

In your aws s3 bucket, you have config file `s3://MY_BUCKET/dev/remote-config/default/config.yaml`. In your service, you get config

```go
package main

import (
	"github.com/lbernardo/remote-config/pkg/sync"
	"github.com/lbernardo/remote-config/pkg/config"
	"fmt"
)

func main() {
	cfg := sync.New()
	cfg.SetEnvironment("dev").
		SetProject("remote-config").
		SetNamespace("default").
		SetBucket("MY_BUCKET")
	cfg.Sync()

	fmt.Println(config.GetString("name"))
}
```

### SetBucket

You can use to override default bucket, the environment variable `GO_REMOTE_CONFIG_BUCKET`

### SetProject

You can use to override default project, the environment variable `PROJECT`

### SetNamespace

You can use to override default namespace, the environment variable `NAMESPACE`

### SetEnvironment

You can use to set environment (is required)

### Sync

Use to download config (`<PROJECT>/<ENVIRONMENT>/<NAMESPACE>/config.yaml`) and read with [spf13/viper](https://github.com/spf13/viper)


## Use Config

You can use function `config.Get(), config.GetString, config.GetInt() ....` to get values in config file. Like  [spf13/viper](https://github.com/spf13/viper)