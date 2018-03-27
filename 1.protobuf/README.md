[TOC]

# protobuf安装和使用

## protobuf安装

1. 首先到 [下载地址](https://github.com/google/protobuf/releases) 下载操作系统对应最新版本并安装，同时设置`protoc`目录到`path`；
2. 安装`proto`和`protoc-gen-go`
```
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
```

## protobuf使用

### proto文件编写


#### proto与golang的参数对应

|proto||golang|java|
|-|-|-|-|
|double||float64|double|
|float||float32|float|
|int32|使用变长编码，对于负值的效率很低，如果你的域有可能有负值，请使用sint64替代|int32|int|
|uint32|使用变长编码|uint32|int|
|uint64|使用变长编码|uint64|long|
|sint32|使用变长编码，这些编码在负值时比int32高效的多|int32|int|
|sint64|使用变长编码，有符号的整型值。编码时比通常的int64高效|uint64|long|
|fixed32|总是4个字节，如果数值总是比总是比228大的话，这个类型会比uint32高效|uint32|int|
|fixed64|总是8个字节，如果数值总是比总是比256大的话，这个类型会比uint64高效|uint64|long|
|sfixed32|总是4个字节|int32|int|
|sfixed64|总是8个字节|int64|long|
|bool||bool|boolean|
|string|一个字符串必须是UTF-8编码或者7-bit ASCII编码的文本|string|String|
|bytes|可能包含任意顺序的字节数据|[]byte|ByteString|


### 使用下面命令生成`*.pb.go`文件

```
protoc --go_out=. *.proto
```

### 在程序中使用生成好的结构体

## 实例代码

1. 最简易版本，请查看`1.protobuf/simple`目录

