# WhatIsThis

一个基于 Wails 和 Vue.js 开发的跨平台桌面应用程序，用于快速识别和分析剪贴板内容。

## 功能特点

- 系统托盘运行，轻量且便捷
- 支持快捷键呼出窗口
- 自动识别剪贴板内容
- GPT 智能分析功能
- 可自定义配置
- 历史记录存储

## 技术栈

- 后端：Go + Wails
- 前端：Vue.js
- 数据存储：NoSQL
- AI 功能：GPT API

## 系统要求

- Windows/macOS/Linux
- Go 1.16+
- Node.js 14+

## 安装说明

1. 克隆项目
```bash
git clone https://github.com/glwlg/whatisthis.git
cd whatisthis
```

2. 安装依赖
```bash
# 安装依赖
go mod tidy
```

3. 运行开发环境
```bash
wails dev
```

4. 构建应用
```bash
wails build
```

## 配置说明

配置文件位于 `configs/config.yaml`，你可以在此设置：

- GPT API 配置
- 快捷键设置
- 界面配置
- 其他系统设置

## 使用说明

1. 启动应用后，程序将在系统托盘运行
2. 使用配置的快捷键呼出主窗口
3. 复制需要分析的内容到剪贴板
4. 程序会自动识别并分析剪贴板内容
5. 可以在设置中自定义配置

## 目录结构

```
whatisthis/
├── frontend/          # 前端 Vue.js 代码
├── internal/          # 内部包
├── pkg/              # 外部包
├── configs/          # 配置文件
├── docs/            # 文档
├── scripts/         # 脚本
├── app.go           # 应用主逻辑
└── main.go          # 程序入口
```

## 贡献指南

欢迎提交 Issue 和 Pull Request 来帮助改进项目。

## 许可证

[MIT License](LICENSE)
