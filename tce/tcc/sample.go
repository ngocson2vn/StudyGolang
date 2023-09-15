package main

import (
    "fmt"
    "context"
    "code.byted.org/gopkg/tccclient"
)

var (
    client *tccclient.ClientV2
)

func init() {
    config := tccclient.NewConfigV2()
    config.Confspace = "default" // Confspace is optional，default value is "default"
    var err error
    client, err = tccclient.NewClientV2("data.reckon.crontask_sre_test", config)
    if err != nil {
        panic(err)
    }
}

func main() {
    ctx := context.Background() // should use framework's ctx when you use some framework
    value, err := client.Get(ctx, "cluster_config")
    // err == tccclient.ConfigNotFoundError
    fmt.Println(value, err)
}
