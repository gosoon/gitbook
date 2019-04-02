# gomock.md

要写出好的测试代码，必须精通相关的测试框架。对于Golang的程序员来说，至少需要掌握下面几个测试框架：

- `GoConvey`
- `GoStub`
- [gomock](github.com/golang/mock)
- `Monkey`
- `mockgen`

golang 的 mock 组件下包含 gomock 和 mockgen 两个测试需要用的库，这连个库我在工作中用的比较多

`mockgen xxx/pkg/dao Storage > mock_dao/dao.go`





mockgen 会自动将为包中的 public 函数生成 mock 函数，然后在测试时调用 mock 函数即可

mockgen 使用示例：

```

#! /bin/sh

mockgen xxx/pkg/dao Storage >mock_dao/dao.go
mockgen xxx/pkg/fedkube FederalClient >mock_fedkube/fedkube.go

```


### gomock 使用示例，部分为调用 mockgen 生成的 mock 函数：
```
func TestCloseTask(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()
    app := &types.StatefulApp{TaskID: 1}
    task := &types.Task{}
    mockDB := mock_dao.NewMockStorage(mockCtrl)
    mockDB.EXPECT().GetStatefulAppByName(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(app, nil).Times(1)
    mockDB.EXPECT().GetTaskByTaskID(gomock.Any(), gomock.Any()).Return(task, nil).Times(1)
    mockDB.EXPECT().UpdateTaskField(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
    con := console.New(&console.Options{DB: mockDB})
    userInfo := utils.UserInfo{}
    ctx := context.WithValue(context.TODO(), utils.DefaultLoginUserInfo, &userInfo)
    err := con.CloseStatefulAppTask(ctx, "feilei", "txy01", "online")
    if err != nil {
        t.Log(err)
        t.Fatal("Failed to pass CloseTask test")
    }
}

```

[参考链接](https://sourcegraph.com/github.com/xanzy/go-gitlab/-/blob/projects_test.go)
