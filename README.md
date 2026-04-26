# CLIProxyAPI

[![Original Project](https://img.shields.io/badge/Original-router--for--me%2FCLIProxyAPI-blue?style=flat-square)](https://github.com/router-for-me/CLIProxyAPI)

> 本项目基于 [router-for-me/CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI) 修改。
>
> 原始项目作者：**router-for-me**
>
> 原始仓库：https://github.com/router-for-me/CLIProxyAPI

## 本分支新增内容

相比原版，本分支增加了以下功能：

### 🔑 API Key 备注/归属功能

在管理页面中，每个 API Key 可关联"备注/归属"信息（如"张三手机"、"办公室电脑"），方便识别各 Key 的使用者和用途。

- 支持添加/编辑 Key 时填写备注
- 备注与 Key 一起保存，完全兼容旧格式（纯字符串 Key 仍可正常使用）
- 前端和后端同步支持，无需破坏性升级

## 项目结构

```
CLIProxyAPI/
├── backend/          # 后端服务（Go）
│   ├── cmd/server/  # 程序入口
│   ├── internal/    # 核心代码
│   └── ...
└── frontend/        # 前端管理页面（React）
    ├── src/        # 源代码
    └── dist/       # 编译产物（management.html）
```

## 快速部署

### 后端

```bash
cd backend
go build -o cli-proxy-api ./cmd/server
sudo systemctl stop cli-proxy-api.service
sudo cp cli-proxy-api /opt/cli-proxy-api/cli-proxy-api
sudo systemctl start cli-proxy-api.service
```

### 前端

```bash
cd frontend
npm install
npm run build
# 将 dist/index.html 重命名为 management.html
# 上传到服务器的 /opt/cli-proxy-api/static/
```

## 原始功能（来自原版）

- 为 CLI 编程工具（Claude Code、Codex、Copilot 等）提供 OpenAI/Gemini/Claude 兼容 API 接口
- OAuth 登录支持
- 多账户轮询负载均衡
- 流式与非流式响应
- 函数调用/工具支持

## 文档

- 后端完整说明：[backend/README_CN.md](backend/README_CN.md)
- 管理 API 文档：[https://help.router-for.me/cn/management/api](https://help.router-for.me/cn/management/api)

## 在线体验

管理后台（需配合后端使用）：
```
http://你的服务器IP:8317/management.html
```

## 许可证

MIT
