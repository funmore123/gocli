---
name: gocli-ark
version: 1.0.0
description: "火山引擎 ARK 模型管理：列出 ARK 平台可用的基础模型，支持按名称筛选。当用户需要查看可用模型、搜索特定模型时使用。"
metadata:
  requires:
    bins: ["gocli"]
---

# gocli ark

## 前置条件

使用火山引擎 AK/SK 登录：

```bash
gocli auth login --api-key <AK> --api-secret <SK>
```

## 命令

### 列出基础模型

```bash
gocli ark +list-models [--name <filter>]
```

| 参数 | 必填 | 说明 |
|------|------|------|
| `--name` | 否 | 按模型名称筛选（子串匹配） |

## 示例

```bash
# 列出所有模型
gocli ark +list-models

# 筛选 doubao 系列模型
gocli ark +list-models --name "doubao"
```

## 编排示例

当用户需要了解平台模型能力时：
1. 先用 `gocli ark +list-models` 获取完整列表
2. 根据返回的 `Items[].FoundationModelTag.Domains` 分类展示
3. 按 `Items[].DisplayName` 和 `Items[].Description` 总结推荐
