1. get package info

```
type empty struct{}

var packagePath = reflect.TypeOf(empty{}).PkgPath()

func getPackageName() string {
    structPackageName := reflect.TypeOf(empty{}).String()
    return strings.Split(structPackageName, ".")[0]
}
```



