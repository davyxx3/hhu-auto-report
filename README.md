# hhu-auto-report

河海大学自动健康打卡，适用于最新的打卡方式（2022年4月份），基于Go语言实现



默认为每天**中午12点**进行自动打卡，若打卡失败会自动重试10次

如果重试10次仍失败，大概率是学校服务器挂掉了，这种情况只能手动打卡

（未来会加入打卡时间和打卡次数的配置）



## 集成组件

- 爬虫框架[scrapy](https://github.com/scrapy/scrapy)
- OCR工具[gosseract](https://github.com/otiai10/gosseract)，用于验证码识别
- 定时任务[cron](https://github.com/robfig/cron)



## 使用方式

### 方式一：Docker安装（推荐）

```bash
# 拉取镜像
docker pull davyxx3/hhu-auto-report

# 按照如下格式配置参数，启动镜像
docker run -dit -e STU_NAME="xxx" \
    -e STU_ID="xxx" \
    -e STU_PWD="xxx" \
    -e STU_INFO="xxx" davyxx3/hhu-auto-report
```

其中的参数说明

- STU_NAME：填写姓名
- STU_ID：填写学号
- STU_PWD：填写HHU奥兰系统的密码
- STU_INFO：填写专业、年级和班级信息（如：电信18_2），一定要按照格式填写，如果不清楚自己的专业简称，可以去奥兰系统打卡界面自行查看



### 方式二：手动安装

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

sudo echo "deb http://mirrors.aliyun.com/debian  stable main contrib non-free" >> sources.list

echo "deb http://mirrors.aliyun.com/debian  stable-updates main contrib non-free" >> sources.list

sudo apt update && sudo apt upgrade

sudo apt install tesseract-ocr

sudo apt install libtesseract-dev
```



##### 其他操作系统

请参考[tesseract安装指导](https://tesseract-ocr.github.io/tessdoc/Installation.html)



#### 2. 启动项目

```bash
git clone https://github.com/davyxx3/hhu-auto-report.git

cd ./hhu-auto-report

go run ./hhu-auto-report.go
```

