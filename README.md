# hhu-auto-report

河海大学自动健康打卡，适用于最新的打卡方式（2022年4月份），基于Go语言实现



默认为每天**中午12点**进行自动打卡，若打卡失败会自动重试10次

如果重试10次仍失败，大概率是学校服务器挂掉了，这种情况只能手动打卡




## 集成组件

- 爬虫框架[colly](https://github.com/gocolly/colly)，用于访问网址、发起请求
- OCR工具[gosseract](https://github.com/otiai10/gosseract)，用于验证码识别
- 定时任务[cron](https://github.com/robfig/cron)，用于每天定时完成打卡任务
- 配置加载工具[ini](https://github.com/go-ini/ini)，用于加载数据源配置




## 使用方式

### 方式一：Github Actions全自动打卡（强烈推荐）

强烈推荐没有个人服务器的同学使用Github Actions进行全自动打卡，非常简单快捷

步骤如下：

1. Fork本仓库，并打开
2. 在选项卡中找到Settings

![image-20220410185059221](https://cdn.jsdelivr.net/gh/davyxx3/pics/blog_img/image-20220410185059221.png)

3. 在Security栏中，找到Secrets选项，点开Actions子选项卡

![image-20220410185408965](https://cdn.jsdelivr.net/gh/davyxx3/pics/blog_img/image-20220410185408965.png)

4. 在页面的右上角点击New repository secret

![image-20220410185705289](https://cdn.jsdelivr.net/gh/davyxx3/pics/blog_img/image-20220410185705289.png)

5. 添加两个secret，Name分别是STU_ID和STU_PWD（一定要叫这个名字，否则无法读取你的学号和密码！），Value就填写自己的学号和奥兰系统密码，完成之后应该是这样的：

![image-20220410190041715](https://cdn.jsdelivr.net/gh/davyxx3/pics/blog_img/image-20220410190041715.png)

6. 点击选项卡中的Actions

![](https://cdn.jsdelivr.net/gh/davyxx3/pics/blog_img/image-20220410191049222.png)

7. 同意workflows的运行

![image-20220410190928640](https://cdn.jsdelivr.net/gh/davyxx3/pics/blog_img/image-20220410190928640.png)

8. 在侧边Workflows栏中选到hhu-auto-report，然后点黄色框框里的Enable workflow

![image-20220410191405423](https://cdn.jsdelivr.net/gh/davyxx3/pics/blog_img/image-20220410191405423.png)至此，自动打卡就已经打开了，默认设置了中午12点进行打卡。

你也可以通过手动启动workflow的方式来进行手动打卡，强烈推荐先手动打一次卡，确认没有问题之后，便可以放心退出了



手动打卡的步骤：

1. 点击蓝色提示框中的Run workflow，然后点击绿色的Run workflow按钮

![image-20220410192107889](https://cdn.jsdelivr.net/gh/davyxx3/pics/blog_img/image-20220410192107889.png)

2. 稍等几秒钟，页面会自动刷新，workflow开始运行

![image-20220410192830304](https://cdn.jsdelivr.net/gh/davyxx3/pics/blog_img/image-20220410192830304.png)

3. 点开hhu-auto-report，可以看见打卡正在运行中

![image-20220410192915290](https://cdn.jsdelivr.net/gh/davyxx3/pics/blog_img/image-20220410192915290.png)

4. 当看到如下提示时，说明打卡成功

![image-20220410193038751](https://cdn.jsdelivr.net/gh/davyxx3/pics/blog_img/image-20220410193038751.png)

因为网络和验证码识别的问题，打卡有概率失败，程序会自动重试10次，但只要最后看到打卡成功的提示就完全OK了



### 方式二：Docker（推荐）

Docker部署方式推荐有个人服务器的同学使用，只需要拉取镜像、启动容器两个步骤，就可以实现每天的自动打卡

```bash
# 拉取镜像
docker pull davyxx3/hhu-auto-report

# 按照如下格式配置参数，启动容器
docker run -d -e STU_ID="xxx" \
    -e STU_PWD="xxx" \
    davyxx3/hhu-auto-report
```

其中的参数说明

- STU_ID：填写学号
- STU_PWD：填写HHU奥兰系统的密码



### 方式三：手动编译（不推荐）

#### 1. 安装依赖

项目使用了[tesseract](https://github.com/tesseract-ocr/tesseract)来识别验证码，需要先安装所需要的依赖



##### Ubuntu/Debian

在Ubuntu/Debian下，可以直接从默认包管理工具直接安装：

```bash
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

因为操作略繁杂，还是推荐Docker的方式部署到自己的服务器上