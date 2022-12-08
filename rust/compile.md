1.为什么 Rust 比 Golang 编译速度慢？
    Go 代码的编译速度不止比 Rust 快，还比 C/C++、Java 等众多语言都快，Go 有简洁的模块依赖关系、优雅的包管理方式、极致的类型系统、高效的编译器，编译速度是 Go 的亮点之一，在语言设计时 Go 团队就为编译速度做过深度的考虑。而 Rust 在编译时解决掉了很多缺陷，如资源释放安全、并发安全和错误处理方面的缺陷，Rust 借鉴了 Haskell，有完整的系统类型，支持泛型。为了性能的考虑，Rust 在处理泛型函数的时候会做单体化，泛型函数里每个用到的类型会编译出一份代码，所以 Rust 在编译的时候速度比较慢。

[Why Go compiles so fast](https://devrajcoder.medium.com/why-go-compiles-so-fast-772435b6bd86)
