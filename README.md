# frpc_use (Silent Guardian)

A high-performance Go-based daemon for **frpc** on Windows. It ensures persistent connectivity with zero console footprint and automated crash recovery.

基于 Go 语言实现的 Windows 平台 **frpc** 守护进程。支持完全静默运行、后台自动保活及无人值守部署。

## Key Features / 核心特性

* **Auto-Recovery**: 5-second polling heartbeat to restart frpc if it crashes or is killed.
* **自动恢复**: 每 5 秒心跳检测，异常退出或被关闭时自动拉起进程。

* **I/O Safety**: Redirects all streams to prevent pipe buffer leaks.
* **I/O 安全**: 显式重定向输出流，防止因缓冲区溢出导致的进程挂起。

## Build / 编译

**Prerequisites**: Go 1.26+

1. **Clone / 克隆**:
```bash
git clone https://github.com/Minakanmi-Yuki/frpc_use.git

```

2. **Compile (Critical Flags) / 编译（关键标志）**:
To hide the daemon's own console window, compile with the following flags:
为了隐藏守护进程自身的窗口，编译时必须添加：
```bash
go build -ldflags="-H windowsgui -s -w" -o dwm.exe main.go

```
