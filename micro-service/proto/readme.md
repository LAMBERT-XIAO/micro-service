# 为Proto文件生成桩文件

## 生成go类型的文件

### 安装 `protoc` 工具和插件

* 下载 `protoc`, 将protoc文件拷贝到 `$GOPATH/bin` 目录下

    https://github.com/protocolbuffers/protobuf/releases

* 安装go插件 `protoc-gen-go`

    ```
    go get -a github.com/golang/protobuf/protoc-gen-go
    ```

### 生成go的桩文件

```
protoc --go_out=. *.proto
```
