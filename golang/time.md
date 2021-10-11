#### 1、time 与 string 之间的格式转换

time 格式转换为 string 格式：

```
time.Format("2006-01-02T15:04:05Z")

time.Format("01/02/2006")
```



string 转 time 格式：

```
// 从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
t, err := time.Parse("2006-01-02T15:04:05Z", "2019-03-05T09:01:11Z")

t, err := time.Parse("01/02/2006", "02/08/2015")
```



 time 与 string 之间可以转换为多种格式，golang time 包中定义有如下多种格式：

```
const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)
```



#### 2、求两个时间点的差值

```
    // 两个时间点的差值
    t, _ := time.Parse("2006-01-02T15:04:05Z", s)
    
    now := time.Now()
    sub := now.Sub(t)
    subHours := sub.Hours()
    
    // 描述几天前的时间
    sevenDaysAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
    today := time.Now().Format("2006-01-02")
```

#### 3、构建自己的 time 包

```
// Timestamp wraps around time.Time and offers utilities to format and parse
// the time using RFC3339Nano
type Timestamp struct {
	time time.Time
}

// NewTimestamp returns a Timestamp object using the current time.
func NewTimestamp() *Timestamp {
	return &Timestamp{time.Now()}
}

// ConvertToTimestamp takes a string, parses it using the RFC3339Nano layout,
// and converts it to a Timestamp object.
func ConvertToTimestamp(timeString string) *Timestamp {
	parsed, _ := time.Parse(time.RFC3339Nano, timeString)
	return &Timestamp{parsed}
}

// Get returns the time as time.Time.
func (t *Timestamp) Get() time.Time {
	return t.time
}

// GetString returns the time in the string format using the RFC3339Nano
// layout.
func (t *Timestamp) GetString() string {
	return t.time.Format(time.RFC3339Nano)
}
```

#### 4、

```
1.time.Now().Add(t.ttl).Unix()
2.    
  t := time.Now() 
  tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
  tm2 := tm1.AddDate(0, 0, -1-i)
  startTime := tm2.UnixNano() / 1e6
```

#### 5、timer 对象池

https://github.com/nats-io/nats.go/blob/main/timer.go

