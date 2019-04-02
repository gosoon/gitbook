Error handling practices in Go

#### 1.Return custom errors

this need define custom errors and return it.


##### (1) define an API for errors 

```
type userNotFound interface {
    UserNotFound() (bool, int64)
}

func IsUserNotFound(err error) (bool, int64) {
    if e, ok := err.(userNotFound); ok {
        return e.UserNotFound()
    }

    return false, 0
}

user, err := repository.GetUser(1)

if ok, id := IsUserNotFound(err); ok {
    // user not found, log the user id
}

```


##### (2) define an API for errors 


#### 2.Don’t just return errors




[参考链接](https://banzaicloud.com/blog/error-handling-go/)
