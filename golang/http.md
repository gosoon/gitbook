example:

```
func (em *exampleManager) ListInfo(request *ListRequest) error {
    const path = "/api/xxx/list"

    var result ListInfoResp
    var errorResult CommonResp
    resp, err := httputils.NewRestCli().
        Post().
        Host(em.opts.ServerAddr).
        ResourcePath(path).
        Object(request).
        SetHeader("Authorization", em.opts.Token).
        Into("200", &result).
        Into("4xx", &errorResult).
        Debug(2).
        Do()
    if err != nil {
        return nil, fmt.Errorf("list xxx info failed with:%v", err)
    }

    if resp.Status != http.StatusOK {
        return nil, fmt.Errorf("status is %v,list xxx info failed with:%v",
            resp.Status, errorResult.ErrorMsg)
    }

    if result.Status != 0 {
        return nil, fmt.Errorf("%s", errorResult.ErrorMsg)
    }
    return nil
}
```


