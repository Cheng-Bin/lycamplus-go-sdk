
# LYCAM+ 直播服务端Golang SDK

该 SDK 适用于 Golang 1.5 及其以上版本 ，基于LYCAM+官方直播API构建 。 若您的服务端是一个基于 Golang 编写的网络程序，使用此 SDK ，能让您以非常便捷地方式接入我们的服务 ，同时也使得您的服务端更加轻盈 。

## 安装

```shell
go get github.com/lycam-dev/lycamplus-go-sdk/lycamplus
```

## 使用

### 配置参数 并创建 SDK 实例

设置全局参数 ，包括必须的 appKey ，appSecret 和 masterSecret ，配置参数将会延至所有空间 。
```
appKey       :=  <您申请到的 AppKey>

appSecret    :=  <您申请到的 AppSecret>

masterSecret :=  <您申请到的 masterSecret>
```

创建 SDK 实例 。

```
lycamPlus := NewLycamPlus(appKey, appSecret, masterSecret)
```

## User 对象

获取 User 对象并进行操作

```javascript
var userInstance = lycamPlus.UserInstance;
```


## Stream 对象

获取 Stream 对象并进行操作

```javascript
var streamInstance = lycamPlus.StreamInstance;
```

