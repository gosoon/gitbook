
get count
```
    sql := "select sum(xx) from xx"
    var counts []int32
    err := m.client.Table(tableName).
        Raw(sql).
        Pluck("SUM(node) as count", &counts).
        Error
```

```
    var counts int64
    db := m.client.Table(xxx).Where("timestamp < ?", startTime).Count(&counts)
```

CreateOrUpdate:
```
func (m *MysqlClient) CreateOrUpdate(xxx *types.xxx) error {
    _, err := m.GetXXXByID(xxx.ID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // create
            if err := m.client.Debug().Create(xxx).Error; err != nil {
                return err
            }
            return nil
        }
        return err
    }
    // update
    if err := m.client.Debug().Model(&types.xxx{}).Where("id = ?", xxx.ID).                             Updates(map[string]interface{}{
        "amount": xxx.Amount,
    }).Error; err != nil {
        return err
    }
    return nil
}
```
