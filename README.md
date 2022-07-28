# QQBot

一个可配置的 QQ 机器人, 需自行搭建 [opqbot](https://opqbot.com/) 配合使用

golang >= 1.18

## 使用方式

直接运行二进制即可, **必须**携带**配置文件路径**参数, **注意, 是配置文件路径, 不包含配置文件名称**

```bash
./qqbot -c {config_path}
```

## 计划支持的功能有

### 基础功能

- 向 个人/群 发送设定的文字
- 向 个人/群 发送萌图
- 向 个人/群 发送功能菜单

## 配置文件

配置文件使用 yaml, 项目正处于开发过程中, **配置文件结构可能会经常变动**

### 配置文件一览

配置文件分为 `system` 和 `function` 两个部分, 你可以在 `config/template.yaml` 中查看样例

```
system: 系统配置相关
    qq: 机器人的 qq 号
    opq: opq 服务器地址
    admin: 管理员 qq, 只能设置一个, 且必须是个人QQ, 而不是群号
    headmsg: 欢迎语设置, 会显示在功能菜单之前
    errmsg: 错误指令回复
function: 数组类型, 数组内每一个对象表示机器人的一个功能
    name: 功能名称
    trigger: 触发条件: 
    type: 功能类型
    desc: 功能描述
    admin-only: 是否仅管理员可见
    config: 功能相关配置
```

### system

| 字段名     | 字段类型   | 字段描述                                |
|---------|--------|-------------------------------------|
| qq      | int64  | 机器人的 qq 号, 此处需要与 OPQ 服务器上登录的 QQ 相匹配 |
| opq     | string | opq 服务器地址                           |
| admin   | string | 管理员 qq, 只能设置一个, 且必须是个人QQ, 而不是群号     |
| headmsg | string | 欢迎语, 会显示在功能菜单之前                     |
| errmsg  | string | 错误指令回复                              |

### function

| 字段名        | 字段类型              | 字段描述                                                      |
|------------|-------------------|-----------------------------------------------------------|
| name       | string            | 功能名称, 会显示在功能列表中, 建议使用英文                                   |
| trigger    | string            | 功能触发机制<br/>all: 全部<br/>user: 私聊<br/>group: 群聊(仅限 被at)     |
| type       | string            | 功能类型<br/> text: 回复普通文本<br/>moetu: 发送一张萌图<br/>hitokoto: 一言 |
| desc       | string            | 功能描述, 会显示在功能列表中<br/>                                      |
| admin-only | bool              | 是否是仅管理员可见的功能                                              |
| config     | map[string]string | 功能配置项, 与 功能类型 相关, 具体见下面内容                                 |

下面给出各类功能的配置格式:

#### 普通文本消息

```yaml
config:
  # 设置固定文字内容
  value: 文本内容
```

#### 图片 API

```yaml
config:
  # 萌图 API, 此 API 应直接返回图片
  api: api 地址
```

#### 一言

```yaml
config:
  # hitokoto API, 此 API 应直接返回一句话
  api: api 地址
```

# 依赖

QQBot 使用了以下的优秀开源项目:

- [OPQBot](https://github.com/mcoo/OPQBot)
- [requests](https://github.com/mcoo/requests)
- [Viper](https://github.com/spf13/viper)