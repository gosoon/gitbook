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

t, err := time.ParseInLocation("2006-1-2", "2022-10-21", time.Local)
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

#### 4、time 包相关的一些用法

```
1.time.Now().Add(t.ttl).Unix()

2.获取某天的时间戳

  t := time.Now() 
  tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
  tm2 := tm1.AddDate(0, 0, -1-i)
  startTime := tm2.UnixNano() / 1e6

3.获取过去一段时间的时间戳
	now := time.Now()
	recordTimestamp := now.Unix()
	latestHourTimestamp := time.Now().Add(-1*time.Hour).UnixNano() / 1e9
	latestDayTimestamp := time.Now().Add(-24*time.Hour).UnixNano() / 1e9
	latestWeekTimestamp := time.Now().Add(-7*24*time.Hour).UnixNano() / 1e9

4.获取某月第一天与最后一天 

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
    d = d.AddDate(0, 0, -d.Day()+1)
    return GetZeroTime(d)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
    return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
    return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

```

#### 5、timer 对象池

https://github.com/nats-io/nats.go/blob/main/timer.go



#### 6、计算两个日期相差多少个自然月和每个自然月的天数

```
type SubDaysByMonth struct {
    Date     string
    UsedDays int
}

func getSubDaysByMonth(returnTime, forecastTime time.Time) []SubDaysByMonth {
    var subDaysByMonthList []SubDaysByMonth
    subMonth := SubNatureMonth(returnTime, forecastTime)
    for i := 0; i <= subMonth; i++ {
        // next month start date
        forecastTime := forecastTime.AddDate(0, i, 0)

        subMonth := SubNatureMonth(returnTime, forecastTime)
        var subDays int
        // 需求和退还时间都在同月
        if i == 0 && subMonth == 0 {
            subDays = SubDays(returnTime, forecastTime) + 1
            // 需求和退还时间不在同月，计算第一个月的天数
        } else if i == 0 {
            lastDateOfMonth := GetLastDateOfMonth(forecastTime)
            subDays = SubDays(lastDateOfMonth, forecastTime) + 1
            // 最后一个月的
        } else if subMonth == 0 {
            firstDateOfMonth := GetFirstDateOfMonth(forecastTime)
            subDays = SubDays(returnTime, firstDateOfMonth) + 1
        } else {
            // 整月的
            firstDateOfMonth := GetFirstDateOfMonth(forecastTime)
            lastDateOfMonth := GetLastDateOfMonth(forecastTime)
            subDays = SubDays(lastDateOfMonth, firstDateOfMonth) + 1
        }
        subDaysByMonth := SubDaysByMonth{Date: GetFirstDateOfMonth(forecastTime).Format("2006-01-02"), UsedDays: subDays}
        subDaysByMonthList = append(subDaysByMonthList, subDaysByMonth)
    }
    return subDaysByMonthList
}


// 计算日期相差多少自然月
func SubNatureMonth(t1, t2 time.Time) (month int) {
    y1 := t1.Year()
    y2 := t2.Year()
    m1 := int(t1.Month())
    m2 := int(t2.Month())

    yearInterval := y1 - y2
    monthInterval := m1 - m2
    return yearInterval*12 + monthInterval
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
    d = d.AddDate(0, 0, -d.Day()+1)
    return GetZeroTime(d)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
    return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
    return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

```

