# hhu-auto-report

河海大学自动健康打卡，适用于最新的打卡方式（2022年4月份），基于Go语言实现

## 集成组件

- 爬虫框架[scrapy](https://github.com/scrapy/scrapy)，用于访问网址、发起请求
- OCR工具[gosseract](https://github.com/otiai10/gosseract)，用于验证码识别

## 使用方式

Github Action全自动操作（测试中）

建立secrets，格式如下：

```ini
[student]
stu_id = xxx
stu_pwd = xxx
```

分别换成自己的学号和奥兰系统登陆密码

