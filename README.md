# hhu-auto-report

河海大学自动健康打卡，适用于最新的打卡方式（2022年4月份），基于Go语言实现



默认为每天**中午12点**进行自动打卡，若打卡失败会自动重试10次

如果重试10次仍失败，大概率是学校服务器挂掉了，这种情况只能手动打卡

（未来会加入打卡时间和打卡次数的配置）



## 集成组件

- 爬虫框架[scrapy](https://github.com/scrapy/scrapy)，用于访问网址、发起请求
- OCR工具[gosseract](https://github.com/otiai10/gosseract)，用于验证码识别
- 定时任务[cron](https://github.com/robfig/cron)，用于每天定时完成打卡任务
- 配置加载工具[ini](https://github.com/go-ini/ini)，用于加载数据源配置



## 使用方式

### 方式一：Docker（推荐）

```bash
# 拉取镜像
docker pull davyxx3/hhu-auto-report

# 按照如下格式配置参数，启动容器
docker run -dit -e STU_ID="xxx" \
    -e STU_PWD="xxx" \
    davyxx3/hhu-auto-report
```

其中的参数说明

- STU_ID：填写学号
- STU_PWD：填写HHU奥兰系统的密码



### 方式二：手动编译（不推荐）

#### 1. 安装依赖

项目使用了[tesseract](https://github.com/tesseract-ocr/tesseract)来识别验证码，需要先安装所需要的依赖



##### Ubuntu

在Ubuntu下，可以直接从默认包管理工具直接安装：

```bash
sudo apt install tesseract-ocr
sudo apt install libtesseract-dev
```



##### Debian

在Debian下，需要切换[阿里云镜像站](https://mirrors.aliyun.com/debian/)，再进行安装

```bash
cd /etc/apt
sudo > sources.list
echo "deb http://mirrors.aliyun.com/debian  stable main contrib non-free" >> sources.list
echo "deb http://mirrors.aliyun.com/debian  stable-updates main contrib non-free" >> sources.list
sudo apt update && sudo apt upgrade
sudo apt install tesseract-ocr
sudo apt install libtesseract-dev
```



##### 其他操作系统

请参考[tesseract安装指导](https://tesseract-ocr.github.io/tessdoc/Installation.html)



#### 2. 导入项目

```bash
git clone https://github.com/davyxx3/hhu-auto-report.git
cd ./hhu-auto-report
```



#### 3. 配置学号和密码

在config.ini文件中，配置自己的学号和密码

如果缺少config.ini文件，程序将无法运行

```ini
[student]
stu_id = 填写你的学号
stu_pwd = 填写你的密码
```



#### 4. 启动程序

```bash
go run ./hhu-auto-report.go
```

可以将程序设为开机启动并放到后台，这样就可以实现每天的自动打卡了