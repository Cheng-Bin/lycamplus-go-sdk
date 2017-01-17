
# LYCAM+ 直播服务端Golang SDK

该 SDK 适用于 Golang 1.5 及其以上版本 ，基于LYCAM+官方直播API构建 。 若您的服务端是一个基于 Golang 编写的网络程序，使用此 SDK ，能让您以非常便捷地方式接入我们的服务 ，同时也使得您的服务端更加轻盈 。

## 安装

```shell
go get github.com/lycam-dev/lycamplus-go-sdk/lycamplus
```

## Demo 参考
```shell
https://github.com/lycam-dev/lycamplus-go-sdk/tree/master/example
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
lycamPlusInstance := NewLycamPlus(appKey, appSecret, masterSecret)
```

## User 对象

获取 User 对象并进行操作

```javascript
var userInstance = lycamPlusInstance.UserInstance;
```

**1. `创建用户`**

创建用户到 LYCAM+ 系统中 ，以便用户操作 API 接口鉴权使用 。
```
userRequestModel := lycamplus.UserRequestModel{
    UserName: "lycamplus_tester",
    Password: "abcdefg123",
    ...
}

userResponseModel, err := userInstance.Create(&userRequestModel)
```

**请求参数**

请求参数传递 UserRequestModel 类型。 UserRequestModel 类型包含以下字段 ：

| 字段名称 | 是否必须 | 数据类型 | 参数说明 |
| ------------- | :-------------: | :-------------: | :-------------------------------------: |
| Username | false | string | 用户名，长度为6-80位，如果为空将随机生成（唯一） |
| Password | false | string | 用户密码，长度8-16位，如果为空将随机生成 |
| Email | false | string | 邮件地址 |
| Phone | false | string | 手机号码 11-20 位 |
| Description | false | string | 描述 4-300个字符 |
| DisplayName | false | string | 显示的昵称，2-20位 |

**返回数据**

接口返回 UserResponseModel 类型 和 error 类型。 error 标识是否出错，UserResponseModel标识返回详细数据。UserResponseModel 类型包含以下字段：

| 字段名称       | 数据类型         | 参数说明                                 |
| ------------- | :-------------: | :-------------------------------------: |
| Username      | string          | 用户名                                   |
| UUID          | string          | 用户唯一身份标识                           |
| Password      | string          | 用户密码（如果密码为随机生成，才返回此字段）    |
| Success       | bool            | 成功标志，成功 true，失败 false             |


**2. `用户token获取`**

用户访问 LYCAM+ 资源操作接口（比如：推流）时需要用户鉴权，我们使用token进行验证 。
```
UUID := "3725d420-dc71-11e6-b191-5f7a2ebf06ef"

tokenResponseModel, err := userInstance.Assume(UUID)
```

**请求参数**

请求参数传递 string 类型， 用于标识用户的唯一ID。

| 请求参数       | 是否必须         | 数据类型          | 参数说明                   |
| ------------- | :-------------: | :-------------: | :-----------------------: |
| uuid          | true            | string          | 用户唯一身份标识( 即uuid )   |

**返回数据**

接口返回 TokenResponseModel 类型 和 error 类型 。 error 标识是否出错，TokenResponseModel 标识返回详细数据。TokenResponseModel 类型包含以下字段：

| 字段名称       | 数据类型         | 参数说明                                          |
| ------------- | :-------------: | :---------------------------------------------: |
| Success       | bool            | 成功标志，成功 true，失败false                      |
| Scope         | string          | 授权资源范围，`*`表示所有资源                        |
| Token         | json object     | token对象 。包括 access_token，expires_in 字段等... |



## Stream 对象

获取 Stream 对象并进行操作

```javascript
var streamInstance = lycamPlusInstance.StreamInstance;
```

**1. `创建视频流`**

在 LYCAM+ 后台系统中创建一条视频流 。 用于返回给终端用户或实现您自己的业务 。
```
streamRequestModel := lycamplus.StreamRequestModel{
    Title: "test_stream",
    Description: "test_stream description"
    ...
}

streamResponseModel, err := streamInstance.Create(&streamRequestModel)
```

**请求参数**

请求参数传递 StreamRequestModel 类型。 StreamRequestModel 类型包含以下字段：

| 字段名称        | 是否必须         | 数据类型          | 参数说明                      |
| -------------- | :-------------: | :-------------: | :---------------------------:|
| User           | false           | string          | 用户唯一身份标识( 即uuid )      |
| Title          | false           | string          | 视频流标题                     |
| Description    | false           | string          | 视频流描述                     |
| ThumbnailUrl   | false           | string          | 视频流封面地址                  | 
| StartLat       | false           | float           | 开始视频的维度坐标               |
| StartLon       | false           | float           | 开始视频的经度坐标               |
| Country        | false           | string          | 国家                           |
| State          | false           | string          | 省份                           |
| City           | false           | string          | 城市                           |
| Privacy        | false           | bool            | 是否私有视频（ true是，false否 ） |

**返回数据**

接口返回 StreamResponseModel 类型 和 error 类型 。 error 标识是否出错，StreamResponseModel 标识返回详细数据。StreamResponseModel 类型包含以下字段：

| 字段名称            | 数据类型         | 参数说明                       |
| ------------------ | :-------------: | :--------------------------: |
| StreamId           | string          | stramId ( 视频流标识 )         |
| Status             | string          | 直播状态(live， over，ready)   |
| StreamUrls         | json object     | 视频播放资源列表                |
| UploadUrl          | string          | 推流地址                       |
| ChatUrl            | string          | 消息服务器地址                  |
| ChatChannel        | string          | 消息服务器频道                  |
| ResourceUrl        | string          | 视屏 HTML 主页地址              |
| Title              | string          | 视频流标题                     |
| Description        | string          | 视频流描述                     |
| ThumbnailUrl       | string          | 视频流封面地址                  | 
| StartLat           | float           | 开始视频的维度坐标               |
| StartLon           | float           | 开始视频的经度坐标               |
| Country            | string          | 国家                           |
| State              | string          | 省份                           |
| City               | string          | 城市                           |
| Privacy            | bool            | 是否私有视频（ true是，false否 ） |

**2. `更新指定ID视频流`**

在 LYCAM+ 后台系统中更新指定 ID 视频流信息 。
```
streamID := "b7d87ea0-dc72-11e6-98af-bb17f4293ffa"

streamRequestModel := lycamplus.StreamRequestModel{
    Title:       "test_stream",
    Description: "no Description",
    ...
}

streamResponseModel, err := streamInstance.Update(streamID, &streamRequestModel)
```

**请求参数**

请求参数传递 string 类型 和 StreamRequestModel 类型。string 类型标识被更新的视频流 ID， StreamRequestModel 类型包含以下字段：

| 字段名称        | 是否必须         | 数据类型          | 参数说明                      |
| -------------- | :-------------: | :-------------: | :---------------------------:|
| Title          | false           | string          | 视频流标题                     |
| Description    | false           | string          | 视频流描述                     |
| ThumbnailUrl   | false           | string          | 视频流封面地址                  | 
| StartLat       | false           | float           | 开始视频的维度坐标               |
| StartLon       | false           | float           | 开始视频的经度坐标               |
| EndLat         | false           | float           | 视频当前的维度坐标               |
| EndLon         | false           | float           | 视频当前的经度坐标               |
| Country        | false           | string          | 国家                           |
| State          | false           | string          | 省份                           |
| City           | false           | string          | 城市                           |
| Privacy        | false           | bool            | 是否私有视频（ true是，false否 ） |


**返回数据**

接口返回 StreamResponseModel 类型 和 error 类型 。 error 标识是否出错，StreamResponseModel 标识返回详细数据。StreamResponseModel 类型包含以下字段：

| 字段名称            | 数据类型         | 参数说明                       |
| ------------------ | :-------------: | :--------------------------: |
| StreamId           | string          | stramId ( 视频流标识 )         |
| Status             | string          | 直播状态(live， over，ready)   |
| StreamUrls         | json object     | 视频播放资源列表                |
| UploadUrl          | string          | 推流地址                       |
| ChatUrl            | string          | 消息服务器地址                  |
| ChatChannel        | string          | 消息服务器频道                  |
| ResourceUrl        | string          | 视屏 HTML 主页地址              |
| Title              | string          | 视频流标题                     |
| Description        | string          | 视频流描述                     |
| ThumbnailUrl       | string          | 视频流封面地址                  | 
| Country            | string          | 国家                           |
| State              | string          | 省份                           |
| City               | string          | 城市                           |
| Privacy            | bool            | 是否私有视频（ true是，false否 ） |
| ...                |                 | 其它视频流参数                   |


**3. `获取指定ID视频流信息`**

在 LYCAM+ 后台系统中获取指定 ID 的视频流 。  用于返回给终端用户或实现您自己的业务 。

```
streamID := "b7d87ea0-dc72-11e6-98af-bb17f4293ffa"

streamResponseModel, err := streamInstance.Show(streamID)
```

**请求参数**

请求参数传递 string 类型 ， 用于标识视频流的ID。

| 请求参数        | 是否必须         | 数据类型          | 参数说明                      |
| -------------- | :-------------: | :-------------: | :---------------------------:|
| streamId       | true            | string          | stramId ( 视频流标识 )         |

**返回字段**

接口返回 StreamResponseModel 类型 和 error 类型 。 error 标识是否出错，StreamResponseModel 标识返回详细数据。StreamResponseModel 类型包含以下字段：

| 字段名称            | 数据类型         | 参数说明                       |
| ------------------ | :-------------: | :--------------------------: |
| StreamID           | string          | stramId ( 视频流标识 )         |
| Status             | string          | 直播状态(live， over，ready)   |
| StreamUrls         | json object     | 视频播放资源列表                |
| UploadUrl          | string          | 推流地址                       |
| ChatUrl            | string          | 消息服务器地址                  |
| ChatChannel        | string          | 消息服务器频道                  |
| ResourceUrl        | string          | 视屏 HTML 主页地址              |
| Title              | string          | 视频流标题                     |
| Description        | string          | 视频流描述                     |
| ThumbnailUrl       | string          | 视频流封面地址                  | 
| StartLat           | float           | 开始视频的维度坐标               |
| StartLon           | float           | 开始视频的经度坐标               |
| Country            | string          | 国家                           |
| State              | string          | 省份                           |
| City               | string          | 城市                           |
| Privacy            | bool            | 是否私有视频（ true是，false否 ） |

**4. `获取视频流列表`**

获取 Lycam+ 后台系统中视频流列表 。用于返回给终端用户或实现您自己的业务 。

```
pageModel := PageModel{
    ResultsPerPage: 2,
}
streamResponseModelList, err := streamInstance.List(&pageModel)
```

**请求参数**

请求参数传递 PageModel 类型 。 PageModel 包含以下字段：

| 字段名称        | 是否必须         | 数据类型          | 参数说明                            |
| -------------- | :-------------: | :-------------: | :---------------------------------:|
| ResultsPerPage | false           | int             | 每页返回记录数 ，默认 10 行            |
| Page           | false           | int             | 返回第几页 ，默认第 1 页               |
| Sort           | false           | string          | 排序字段（ id，description，created ）|
| Order          | false           | string          | 排序方向（ asc，desc ）                |

**返回字段**

接口返回 StreamResponseModelList 类型 和 error 类型 。 error 标识是否出错，StreamResponseModelList 标识返回详细数据。StreamResponseModelList 类型包含以下字段：

| 字段名称            | 数据类型         | 参数说明                         |
| ------------------ | :-------------: | :----------------------------: |
| TotalItems         | int             | 记录总数                         |
| ResultsPerPage     | int             | 每一页数量                       |
| NextPageAvailable  | bool            | 是否有下一页                     | 
| Items              | array           | 视频流清单数组                    |


**5. `获取指定时间前视频流列表`**

在 LYCAM+ 后台系统获取指定时间前的视频流列表 。 用于返回给终端用户或实现您自己的业务 。

```
pageModel := PageModel{
    ResultsPerPage: 2,
}

response, err := streamInstance.ListSince(time.Now().UnixNano(), &pageModel)
```

**请求参数**

请求参数传递 int64 类型 和 PageModel 类型 。int64 类型标识纳秒时间戳， PageModel 类型包含以下字段：

| 字段名称        | 是否必须         | 数据类型          | 参数说明                            |
| -------------- | :-------------: | :-------------: | :---------------------------------:|
| ResultsPerPage | false           | int             | 每页返回记录数 ，默认 10 行            |
| Page           | false           | int             | 返回第几页 ，默认第 1 页               |
| Sort           | false           | string          | 排序字段（ id，description，created ）|
| Order          | false           | string          | 排序方向（ asc，desc ）                |

**返回字段**

接口返回 StreamResponseModelList 类型 和 error 类型 。 error 标识是否出错，StreamResponseModelList 标识返回详细数据。StreamResponseModelList 类型包含以下字段：

| 字段名称            | 数据类型         | 参数说明                         |
| ------------------ | :-------------: | :----------------------------: |
| TotalItems         | int             | 记录总数                         |
| ResultsPerPage     | int             | 每一页数量                       |
| NextPageAvailable  | bool            | 是否有下一页                     | 
| Items              | array           | 视频流清单数组                    |


**6. `通过关键词搜索视频流`**

通过关键词在 LYCAM+ 后台系统获取视频流列表 。 用于返回给终端用户或实现您自己的业务 。

```
keywordModel := KeywordModel{
    Keyword: "lycamplus666"
}

response, err := streamInstance.SearchByKeyword(&keywordModel)
```

**请求参数**

请求参数传递 KeywordModel 类型 。 KeywordModel 包含以下字段：

| 字段名称        | 是否必须         | 数据类型          | 参数说明                            |
| -------------- | :-------------: | :-------------: | :---------------------------------:|
| Keyword        | true            | string          | 搜索关键词                           |
| ResultsPerPage | false           | int             | 每页返回记录数 ，默认 10 行            |
| Page           | false           | int             | 返回第几页 ，默认第 1 页               |
| Sort           | false           | string          | 排序字段（ id，description，created ）|
| Order          | false           | string          | 排序方向（ asc，desc ）               |

**返回字段**

接口返回 StreamResponseModelList 类型 和 error 类型 。 error 标识是否出错，StreamResponseModelList 标识返回详细数据。StreamResponseModelList 类型包含以下字段：

| 字段名称            | 数据类型         | 参数说明                         |
| ------------------ | :-------------: | :----------------------------: |
| TotalItems         | int             | 记录总数                         |
| ResultsPerPage     | int             | 每一页数量                       |
| NextPageAvailable  | bool            | 是否有下一页                     | 
| Items              | array           | 视频流清单数组                    |

**7. `通过地理位置搜索视频流`**

通过地理位置在 LYCAM+ 后台系统获取视频流列表 。 用于返回给终端用户或实现您自己的业务 。

```
locationModel := lycamplus.LocationModel {
    Lon:    90,
    Lat:    90,
    Radius: 100,
}

streamResponseModelList, err := streamInstance.SearchByLocation(&locationModel)
```

**请求参数**

请求参数传递 LocationModel 类型 。LocationModel 包含以下字段：

| 字段名称        | 是否必须         | 数据类型          | 参数说明                            |
| -------------- | :-------------: | :-------------: | :---------------------------------:|
| Lon            | true            | float           | 经度
| Lat            | true            | float           | 纬度
| Radius         | true            | float           | 搜索半径
| ResultsPerPage | false           | int             | 每页返回记录数 ，默认 10 行            |
| Page           | false           | int             | 返回第几页 ，默认第 1 页               |
| Sort           | false           | string          | 排序字段（ id，description，created ）|
| Order          | false           | string          | 排序方向（ asc，desc ）               |

**返回字段**

接口返回 StreamResponseModelList 类型 和 error 类型 。 error 标识是否出错，StreamResponseModelList 标识返回详细数据。StreamResponseModelList 类型包含以下字段：

| 字段名称            | 数据类型         | 参数说明                         |
| ------------------ | :-------------: | :----------------------------: |
| TotalItems         | int             | 记录总数                         |
| ResultsPerPage     | int             | 每一页数量                       |
| NextPageAvailable  | bool            | 是否有下一页                     | 
| Items              | array           | 视频流清单数组                    |


**8. `销毁指定ID视频流`**

销毁在 LYCAM+ 后台系统中指定ID的视频流 。

```
streamID := "b7d87ea0-dc72-11e6-98af-bb17f4293ffa"

successModel, err := streamInstance.Delete(streamID)
```

**请求参数**

请求参数传递 string 类型 ， 用于标识视频流的ID。

| 请求参数       | 是否必须         | 数据类型          | 参数说明                   |
| ------------- | :-------------: | :-------------: | :-----------------------: |
| streamId      | true            | string          | stramId ( 视频流标识 )      |

**返回字段**

接口返回 SuccessModel 类型 和 error 类型 。 error 标识是否出错，SuccessModel 标识返回详细数据。SuccessModel 类型包含以下字段：

| 字段名称       | 数据类型         | 参数说明                         |
| ------------- | :-------------: | :----------------------------: |
| Success       | bool            | 成功标志 。成功 true，失败 false   |