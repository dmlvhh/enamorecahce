# enamorecache 是一个轻量级的缓存工具

## 项目架构
```
enamorecache/
|--lru/
|--lru.go  // lru 缓存淘汰策略
|--byteview.go // 缓存值的抽象与封装
|--cache.go    // 并发控制
|--enamorecache.go // 负责与外部交互，控制缓存存储和获取的主流程
|--http.go     // 提供被其他节点访问的能力(基于http)
```

## 主体结构 Group
Group 是 enamorecache 最核心的数据结构，负责与用户的交互，并且控制缓存值存储和获取的流程。
```
                            是
接收 key --> 检查是否被缓存 -----> 返回缓存值 ⑴
                |  否                         是
                |-----> 是否应当从远程节点获取 -----> 与远程节点交互 --> 返回缓存值 ⑵
                            |  否
                            |-----> 调用`回调函数`，获取值并添加到缓存 --> 返回缓存值 ⑶
```