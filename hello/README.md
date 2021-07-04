# Hello

## 1 代理设置

### 1.1 Go语言在 1.13 版本之后 GOPROXY 默认值为 https://proxy.golang.org，在国内可能会存在下载慢或者无法访问的情况，所以十分建议大家将 GOPROXY 设置为国内的 goproxy.cn。
### go env -w GOPROXY=https://goproxy.cn,direct

## 2 go module 创建

### 2.1 在终端进入到你自身模块的根目录hello，执行命令：go mod init hello， 在目录下会生成go.mod文件

### 2.2 在目录下创建一个go文件，引入需要的包，比如import "rsc.io/quote"

### 2.2 执行go mod tidy，就会自动把包依赖下载下来，然后在go.mod下面会生成一个go.sum文件，里面就是本次依赖包的信息

## 3 执行代码

### 3.1 在终端运行命令 go run .\hello.go 执行代码成功