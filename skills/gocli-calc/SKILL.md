---
name: gocli-calc
version: 1.0.0
description: "gocli 计算器 skill: 使用 gocli CLI 的 math +calc 命令执行四则运算。"
metadata:
  requires:
    bins: ["gocli"]
---

# gocli math calc

简单的四则运算命令。

## 前置条件

先完成登录：

```bash
gocli auth login --api-key <your-key> --api-secret <your-secret>
```

## 命令

```bash
gocli math +calc --expr "1+2*3"
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--expr` | 是 | 数学表达式，支持 +、-、*、/ |

## 示例

```bash
gocli math +calc --expr "100/3"
gocli math +calc --expr "42*2+1"
```
