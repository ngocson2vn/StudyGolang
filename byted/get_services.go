package main

import (
    "fmt"
    "context"
    "time"
    "encoding/json"
    "code.byted.org/gopkg/pkg/log"
    "code.byted.org/lagrange/crontask/pkg/reckon"
    "code.byted.org/lagrange/crontask/pkg/metrics_util"
)

const (
    getServicesTTL = 150 * time.Second
)

func main() {
    metrics_util.InitMetrics("data.reckon.crontasks_sre_test")
    reckon.InitService("data.reckon.reckon_go")
    ctx, cancel := context.WithTimeout(context.Background(), getServicesTTL)
    filterOpts := []reckon.Option{reckon.WithIDC("sg1")}
    ss, err := reckon.GetServices(ctx, filterOpts...)
    cancel()
    if err != nil {
        log.Errorf("reckon api failed: %v", err)
    }

    b, err := json.MarshalIndent(ss, "", "  ")
    fmt.Println(string(b))
}