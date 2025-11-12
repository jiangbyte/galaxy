使用 Apifox 测试你的配置管理接口，可以按照以下步骤进行：

## 1. 首先设置环境变量

在 Apifox 中设置：
- **Base URL**: `http://你的服务器地址:端口/api/v1`
- **Authorization**: 设置认证 token

## 2. 创建配置管理接口集合

### 认证接口（先获取token）
```
POST /auth/login
Body:
{
"username": "admin",
"password": "123456",
"captcha_id": "xxx",
"captcha_code": "xxx"
}
```

### 配置分组管理接口

#### 创建配置分组
```
POST /config/groups
Headers:
Authorization: Bearer {token}

Body:
{
"name": "系统配置",
"code": "system",
"description": "系统相关配置分组",
"sort": 1,
"is_system": true
}
```

#### 获取配置分组列表
```
GET /config/groups
Headers:
Authorization: Bearer {token}

Query Parameters:
current: 1
size: 10
sort_field: created_at
sort_order: descend
keyword: 系统
```

#### 获取配置分组详情
```
GET /config/groups/{id}
Headers:
Authorization: Bearer {token}
```

#### 更新配置分组
```
PUT /config/groups/{id}
Headers:
Authorization: Bearer {token}

Body:
{
"name": "系统配置更新",
"code": "system",
"description": "更新后的系统配置分组",
"sort": 1,
"is_system": true
}
```

#### 删除配置分组
```
DELETE /config/groups/{id}
Headers:
Authorization: Bearer {token}
```

### 配置项管理接口

#### 创建配置项
```
POST /config/items
Headers:
Authorization: Bearer {token}

Body:
{
"group_id": "分组ID",
"name": "站点名称",
"code": "site_name",
"value": "我的网站",
"component_type": "input",
"description": "网站名称配置",
"sort": 1
}
```

#### 获取配置项列表
```
GET /config/items
Headers:
Authorization: Bearer {token}

Query Parameters:
current: 1
size: 10
sort_field: sort
sort_order: ascend
keyword: 站点
```

#### 获取配置项详情
```
GET /config/items/{id}
Headers:
Authorization: Bearer {token}
```

#### 更新配置项
```
PUT /config/items/{id}
Headers:
Authorization: Bearer {token}

Body:
{
"group_id": "分组ID",
"name": "站点名称更新",
"code": "site_name",
"value": "我的网站更新版",
"component_type": "input",
"description": "更新后的网站名称配置",
"sort": 1
}
```

#### 删除配置项
```
DELETE /config/items/{id}
Headers:
Authorization: Bearer {token}
```
