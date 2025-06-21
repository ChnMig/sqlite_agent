# Sqlite Agent 中文文档

[English](./README.md) | [中文](./README_ZH.md)

为 Sqlite 提供 API 调用能力

![image](/doc/img/sqlite_agent.jpg)

## 构建

参考：https://github.com/mattn/go-sqlite3?tab=readme-ov-file#cross-compile

### Linux 构建命令

```
CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static"
```

## 运行

```
./sqlite-agent --config=config.yaml
```

## 帮助

```bash
./sqlite-agent --help   
Usage of ./sqlite-agent:
  -config string
        配置文件路径
  -dev
        以开发模式运行
```

## 配置

请参考 example.yaml，并查看注释以理解各配置项含义，配置简单明了。

## QA

- 请务必修改配置文件中的 api_auth，防止接口被滥用！！！
