package main

import (
    "bytes"
    // "io"
    // "os"
    "strings"
    "fmt"
    "os/exec"
)

func main() {
    manifest := exec.Command(`/usr/bin/printf`, `apiVersion: batch/v1\nkind: Job\nmetadata:\n  name: sample-job\nspec:\n  template:\n    spec:\n      containers:\n        - name: finc-point\n          image: 698186686422.dkr.ecr.us-east-1.amazonaws.com/finc_point:deployeks-staging20190125091754\n          envFrom:\n            - secretRef:\n                name: finc-point-secret\n          command: [\"bash\", \"-l\", \"-c\", \"while true; do date; sleep 1; done\"]\n      restartPolicy: Never`)
    var out bytes.Buffer
    manifest.Stdout = &out
    manifest.Start()
    manifest.Wait()
    fmt.Println(out.String())

    batch_command := "bash -lc \"while true; do date; sleep 1; done\""
    cmd := fmt.Sprintf("command: %s", strings.Split(batch_command, " "))
    fmt.Println(cmd)
}

