# hhu-auto-report

河海大学自动健康打卡，适用于最新的打卡方式（2022年4月份），基于Go语言实现



## 集成组件

- 爬虫框架[scrapy](https://github.com/scrapy/scrapy)，用于访问网址、发起请求
- OCR工具[gosseract](https://github.com/otiai10/gosseract)，用于验证码识别

- 配置加载工具[ini](https://github.com/go-ini/ini)，用于加载数据源配置



## 使用方式

### 方式一：Github Action全自动操作（推荐）

建立secrets，格式如下：

```ini
[student]
stu_id = xxx
stu_pwd = xxx
```

分别换成自己的学号和奥兰系统登陆密码



### 方式二：Docker

```bash
# 拉取镜像
docker pull davyxx3/hhu-auto-report

# 按照如下格式配置参数，启动镜像
docker run -dit -e STU_ID="xxx" \
    -e STU_PWD="xxx" \
    davyxx3/hhu-auto-report
```

其中的参数说明

- STU_ID：填写学号
- STU_PWD：填写HHU奥兰系统的密码

