###  项目介绍及使用
 本工程包含两个功能：二维码识别，二维码解析 均为最新代码(请查看日期)
#### 功能
+ 二维码识别 see: [main.go](./main.go)
    - 使用 `github.com/tuotoo/qrcode`
+ 二维码解析 see: [main.go](./main.go)
    - 使用 `github.com/skip2/go-qrcode`

#### 环境
+ go `VERSION>=1.12.X` 含go module
+ upx(可选)
    - mac : `brew install upx`
    - window see: [upx download](https://github.com/upx/upx/releases/tag/v3.95)
    - linux see: [upx download](https://github.com/upx/upx/releases/tag/v3.95)

#### 初始化(init)
+ 项目结构(project tree)
```
[FILE_BASE_DIR]/go-project-example
                └── src
                    └── qrCodes
                        ├── README.md
                        ├── controller
                        │   └── PageContrller.go
                        ├── go.mod
                        ├── go.sum
                        ├── main.exe
                        ├── main.go
                        └── template
                            └── page
                                ├── index.html
                                └── qrCode.html
```
+ 安装依赖(setup dependency)
    - `go get github.com/skip2/go-qrcode`
    - `go get github.com/tuotoo/qrcode`

#### 运行(run)
+ `cd [FILE_BASE_DIR]/go-project-example/src/qrCodes`
+ `go run main.go`
+ 访问(visit)：`http://127.0.0.1:12345`

#### 打包(package)
+ 构建原始包
    - linux `GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" ./main.go`
    - window `GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" ./main.go`
    - mac `GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" ./main.go`
+ 使用upx加壳
    - `upx --backup --brute [main.exe(windows) or main(linux、mac)]`