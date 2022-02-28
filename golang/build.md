编译时将 git 相关信息传入到代码内部


```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.GoVersion=$(go version)' -X 'main.GitHash=$(git show -s --format=%H)' -X 'main.BuildTime=$(git show -s --format=%cd)'" -o _output/bin/descheduler sigs.k8s.io/descheduler/cmd/descheduler


var (
    GitHash   string
    BuildTime string
    GoVersion string
)

func init() {
    multi_cluster.BuildVersion = BuildTime
}

func main() {
    out := os.Stdout
    cmd := app.NewDeschedulerCommand(out)
    cmd.AddCommand(app.NewVersionCommand())

    logs.InitLogs()
    defer logs.FlushLogs()

    klog.Infof("Git Commit Hash: %s", GitHash)
    klog.Infof("Build TimeStamp: %s", BuildTime)
    klog.Infof("GoLang Version: %s", GoVersion)

    if err := cmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```
