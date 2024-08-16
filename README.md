# SysStress Usage Guide

SysStress is a powerful tool designed to simulate high loads on your system’s CPU, memory, and other resources, allowing you to test and observe system performance under stress. Below is an overview of how to use SysStress effectively.

## 1. Installation
To install SysStress, you can download the binary from the official repository or compile it from source. For example:
```
git clone git@github.com:baixiaozhou/SysStress.git
cd SysStress
GOOS=linux GOARCH=amd64 go build -o sysstress
```

## 2. Basic Commands

### CPU Stress Test

```
./sysstress cpu --cpu-number 10 --duration 10m [--force true|false]
```
This command will stress the CPU for 10 minutes on 10 cores. Valid time units are "s", "m", "h", such as "300s", "1.5h" or "2h45m"

When the requested number of cores exceeds the system’s CPU core count, the operation will be denied. However, you can bypass this restriction by using the `--force` option

### Memory Stress Test

```
./sysstress memory --size 1G --duration 10m [--force true|false]
```
This command will use 1 GB of available memory for a duration of 10 minutes. Valid memory units are "G","M","K", such as "1G", "125M", "32K" 

When the requested memory exceeds the total system memory, the operation will be denied. If the requested memory exceeds the available memory, a warning will be issued that this operation is not allowed unless the `--force true` parameter is added. This consideration is based on the fact that available memory is a dynamic value, and exceeding it during allocation could result in an OOM (Out of Memory) error.

### Run in the background

If You want run these command in background,you can use:
```
nohup ./sysstress xxx &
```
also, if you don't want any output, you can use:
```
nohup ./sysstress xxx > /dev/null 2>&1 &

```
