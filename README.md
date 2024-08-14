# SysStress Usage Guide

SysStress is a powerful tool designed to simulate high loads on your system’s CPU, memory, and other resources, allowing you to test and observe system performance under stress. Below is an overview of how to use SysStrain effectively.

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
./sysstress cpu --cpu-number 10 --duration 10m
```
This command will stress the CPU for 10 minutes on 10 cores. Valid time units are "s", "m", "h", such as "300s", "1.5h" or "2h45m"

### Memory Stress Test

```
./sysstress memory --size 1G --duration 10m
```
This command will use 1 GB of available memory for a duration of 10 minutes. Valid memory units are "G","M","K", such as "1G", "125M", "32K" 

### Run in the background

If You want run these command in background,you can use:
```
nohup ./sysstress xxx &
```
also, if you don't want output, you can use:
```
nohup ./sysstress xxx & 2>&1

```
