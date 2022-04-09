package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/otiai10/gosseract/v2"
)

type StuData struct {
	stuId  string
	stuPwd string
}

var stuData = StuData{}
var loginData = url.Values{}
var reportData = url.Values{}

var timeStr = time.Now().Format("2006-01-02")
var maxRetry = 10

func main() {
	// 加载学生数据
	err := stuData.loadStuData()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 开始打卡
	log.Println("打卡开始！")
	report()
}

// 加载学生数据
func (stuData *StuData) loadStuData() error {
	var isValid bool
	stuData.stuId, isValid = os.LookupEnv("STU_ID")
	if !isValid {
		return errors.New("无法读取学号！启动失败！ :(")
	}
	log.Println("成功读取学号！")
	stuData.stuPwd, isValid = os.LookupEnv("STU_PWD")
	if !isValid {
		return errors.New("无法读取密码！启动失败！ :(")
	}
	log.Println("成功读取密码！")
	return nil
}

func report() {
	var retryCount = 0
	err := reportTry()
	// 若打卡失败，则重试，直到超过阈值为止
	for err != nil {
		log.Println("打卡失败, 正在重试中...")
		if retryCount > maxRetry {
			log.Fatal(err.Error())
			return
		}
		err = reportTry()
		retryCount++
	}
	log.Println("打卡成功! :)")
}

func reportTry() error {
	var err error

	openPage := colly.NewCollector()
	imageDownloader := openPage.Clone()
	login := imageDownloader.Clone()
	reportPage := login.Clone()
	reporter := reportPage.Clone()

	// 进入登陆页面，获取验证码
	openPage.OnResponse(func(r *colly.Response) {
		imageDownloader.Request("GET", "http://smst.hhu.edu.cn/Vcode.ASPX", nil, r.Ctx, nil)
	})

	// 识别验证码
	imageDownloader.OnResponse(func(r *colly.Response) {
		client := gosseract.NewClient()
		defer client.Close()
		client.SetImageFromBytes(r.Body)
		code, _ := client.Text()
		loginData.Set("vcode", code)
	})

	// 解析登陆页面，构建登陆的报文数据
	openPage.OnHTML("#form1", func(e *colly.HTMLElement) {
		viewState := e.ChildAttr("#__VIEWSTATE", "value")
		viewStateGenerator := e.ChildAttr("#__VIEWSTATEGENERATOR", "value")
		loginDataInit(&loginData)
		loginData.Set("__VIEWSTATE", viewState)
		loginData.Set("__VIEWSTATEGENERATOR", viewStateGenerator)
	})

	// 发出登陆的POST请求
	openPage.OnScraped(func(r *colly.Response) {
		login.Request("POST", "http://smst.hhu.edu.cn/login.aspx", strings.NewReader(loginData.Encode()), r.Ctx, nil)
	})

	// 进入健康打卡页面
	login.OnResponse(func(r *colly.Response) {
		reportPage.Request("GET", "http://smst.hhu.edu.cn/txxm/rsbulid/r_3_3_st_jkdk.aspx?xq=2021-2022-2&nd=2018&msie=1", nil, r.Ctx, nil)
	})

	// 解析健康打卡页面，构建健康打卡的报文数据
	reportPage.OnHTML("#form1", func(e *colly.HTMLElement) {
		viewState := e.ChildAttr("#__VIEWSTATE", "value")
		viewStateGenerator := e.ChildAttr("#__VIEWSTATEGENERATOR", "value")
		stuName := e.ChildAttr("#xm", "value")
		stuInfo := e.ChildAttr("#bjhm", "value")
		reportDataInit(&reportData)
		reportData.Set("uname", stuName)
		reportData.Set("xm", stuName)
		reportData.Set("bjhm", stuInfo)
		reportData.Set("__VIEWSTATE", viewState)
		reportData.Set("__VIEWSTATEGENERATOR", viewStateGenerator)
	})

	// 发出健康打卡的POST请求
	reportPage.OnScraped(func(r *colly.Response) {
		reporter.Request("POST", "http://smst.hhu.edu.cn/txxm/rsbulid/r_3_3_st_jkdk.aspx?xq=2021-2022-2&nd=2018&msie=1", strings.NewReader(reportData.Encode()), r.Ctx, nil)
	})

	// 获取打卡结果，完成打卡
	reporter.OnHTML("#cw", func(e *colly.HTMLElement) {
		result := e.Attr("value")
		if result != "新建成功!" && result != "保存修改成功!" {
			err = errors.New("打卡失败！ :(")
		}
	})

	openPage.Visit("http://smst.hhu.edu.cn/login.aspx")
	return err
}

// 初始化登陆报文的数据
func loginDataInit(data *url.Values) {
	data.Set("yxdm", "10294")
	data.Set("__VIEWSTATEENCRYPTED", "")
	data.Set("userbh", stuData.stuId)
	data.Set("cw", "")
	data.Set("xzbz", "1")
	data.Set("pas2s", generateCryptedPwd(stuData.stuPwd))
}

// 初始化健康打卡报文的数据
func reportDataInit(reportData *url.Values) {
	reportData.Set("__EVENTTARGET", "databc")
	reportData.Set("__EVENTARGUMENT", "")
	reportData.Set("__VIEWSTATEENCRYPTED", "")
	reportData.Set("tbrq", timeStr)
	reportData.Set("twqk", "正常（37.3℃及以下）")
	reportData.Set("twqkdm", "1")
	reportData.Set("sfzx", "校内")
	reportData.Set("sfzxdm", "1")
	reportData.Set("brjkqk", "健康")
	reportData.Set("brjkqkdm", "1")
	reportData.Set("tzrjkqk", "健康")
	reportData.Set("tzrjkqkdm", "1")
	reportData.Set("sfjczgfx", "否")
	reportData.Set("sfjczgfxdm", "2")
	reportData.Set("jkmys", "绿码")
	reportData.Set("jkmysdm", "1")
	reportData.Set("xcmqk", "否")
	reportData.Set("xcmqkdm", "2")
	reportData.Set("brcnnrss", "ON")
	reportData.Set("ck_brcnnrss", "false")
	reportData.Set("pzd_lock", "uname,")
	reportData.Set("xdm", "06")
	reportData.Set("xh", stuData.stuId)
	reportData.Set("qx_r", "1")
	reportData.Set("qx_i", "1")
	reportData.Set("qx_u", "1")
	reportData.Set("qx_d", "0")
	reportData.Set("databcxs", "1")
	reportData.Set("pkey", timeStr)
	reportData.Set("dcbz", "1")
	reportData.Set("xqbz", "1")
	reportData.Set("st_xq", "2021-2022-2")
	reportData.Set("msie", "1")
	reportData.Set("tkey", "false")
	reportData.Set("ck_brcnnrss", "tbrq")
}

// 生成密码加密后的哈希值
func generateCryptedPwd(pwd string) string {
	hash := md5.Sum([]byte(strings.ToUpper(pwd)))
	return strings.ToUpper(fmt.Sprintf("%x", hash))
}
