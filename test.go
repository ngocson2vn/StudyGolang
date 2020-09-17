package main

import (
    "bytes"
    "io"
    "os"
    "fmt"
    "os/exec"
    "strings"
)

func main() {
    jobName := "finc-point-job"
    batch_command := []string {"bash", "-lc", "\"i=0; while [ $i -lt 120 ]; do date; i=$((i+1)); sleep 1; done\""}
    manifest_template := `apiVersion: batch/v1\nkind: Job\nmetadata:\n  name: %s\nspec:\n  template:\n    spec:\n      containers:\n        - name: %s\n          image: %s\n          envFrom:\n            - secretRef:\n                name: %s-secret\n          command: [%s]\n      restartPolicy: Never\n  backoffLimit: 1`
    manifest := exec.Command(`/usr/bin/printf`, 
        fmt.Sprintf(manifest_template,
            jobName,
            "finc-point",
            "698186686422.dkr.ecr.us-east-1.amazonaws.com/finc_point:deployeks-staging20190125091754",
            "finc-point",
            strings.Join(batch_command[:], ",")))

    // Delete old job
    kubectl := exec.Command("/usr/local/bin/kubectl", "delete", "job", jobName)
    kubectl.Env = append(os.Environ(), "KUBECONFIG=/Users/nguyen.son/.kube/kubeconfig_staging-workers")
    err := kubectl.Run()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // Create new job
    kubectl = exec.Command("/usr/local/bin/kubectl", "apply", "-f", "-")
    kubectl.Env = append(os.Environ(), "KUBECONFIG=/Users/nguyen.son/.kube/kubeconfig_staging-workers")

    r, w := io.Pipe()
    manifest.Stdout = w
    kubectl.Stdin = r

    var manifest_stderr bytes.Buffer
    manifest.Stderr = &manifest_stderr

    var kubectl_stdout bytes.Buffer
    var kubectl_stderr bytes.Buffer
    kubectl.Stdout = &kubectl_stdout
    kubectl.Stderr = &kubectl_stderr

    manifest.Start()
    kubectl.Start()

    err = manifest.Wait()
    if err != nil {
        fmt.Println(err.Error())
        fmt.Println(manifest_stderr.String())
        return
    }
    w.Close()

    err = kubectl.Wait()
    if err != nil {
        fmt.Println(err.Error())
        fmt.Println(kubectl_stderr.String())
        return
    }
    r.Close()

    fmt.Println(kubectl_stdout.String())
}