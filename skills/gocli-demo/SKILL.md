---
name: gocli-demo
version: 1.0.0
description: "gocli demo skill: 使用 gocli CLI 的 demo +greet 命令向指定对象打招呼。"
metadata:
  requires:
    bins: ["gocli"]
---

# gocli demo

## 前置条件

先完成登录：

```bash
gocli auth login --api-key <your-key> --api-secret <your-secret>
```

## 命令

```bash
# 向指定对象打招呼
gocli demo +greet --name "World"
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--name` | 否 | 打招呼对象，默认 "World" |

## 示例

```bash
# 默认
gocli demo +greet

# 指定名字
gocli demo +greet --name "ByteDance"
```
