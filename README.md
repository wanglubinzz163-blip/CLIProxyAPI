# CLIProxyAPI

为 CLI 编程工具（Claude Code、Codex、Copilot 等）提供 OpenAI/Gemini/Claude 兼容 API 接口的代理服务器。

## 项目结构

```
CLIProxyAPI/
├── backend/          # 后端服务（Go）
│   ├── cmd/server/   # 程序入口
│   ├── internal/     # 核心代码
│   └── ...
└── frontend/        # 前端管理页面（React）
    ├── src/         # 源代码
    └── dist/        # 编译产物（management.html）
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
# 将 dist/index.html 重命名为 management.html，上传到服务器的 /opt/cli-proxy-api/static/
```

## 功能

- API Key 备注/归属（可标记每个 Key 给谁/哪台终端使用）
- 兼容 OpenAI/Gemini/Claude 格式
- OAuth 登录支持
- 多账户轮询负载均衡

## 文档

- 后端完整说明：[backend/README_CN.md](backend/README_CN.md)
- 管理 API 文档：[/opt/cli-proxy-api/README_CN.md](https://help.router-for.me/cn/management/api)

## 在线体验

管理后台（需配合后端使用）：
```
http://你的服务器IP:8317/management.html
```

## 许可证

MIT
