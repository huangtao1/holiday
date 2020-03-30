package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", InfoShow)
	r.GET("/holiday", JudgeDayIsHoliday)
	r.Run()
}

//HolidayInfo 假日表
type HolidayInfo struct {
	Holidays []string
	Workdays []string
}

//判断某个元素是否存在
func isExist(i interface{}, l []string) bool {
	for _, v := range l {
		if i == v {
			return true
		}
	}
	return false
}

//InfoShow 打印如何使用
func InfoShow(c *gin.Context) {
	c.JSON(http.StatusOK, "传入日期如:/holday?d=20200320")
}

//JudgeDayIsHoliday 判断传入日期是否为节假日
func JudgeDayIsHoliday(c *gin.Context) {
	//获取传入的日期
	day := c.Query("d")
	log.Println(day)
	// 判断日期是否在json文件中，如果在直接返回对应的值
	fileObj, err := ioutil.ReadFile("holiday.json")
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"code": 502, "message": "找不到配置文件！！！"})
		return
	}
	holidayInfo := HolidayInfo{}
	//创建json解码器
	jsonErr := json.Unmarshal(fileObj, &holidayInfo)
	if jsonErr != nil {
		log.Println(jsonErr)
	}
	//首先判断日期是否在json中定义，如果定义了直接返回
	if isExist(day, holidayInfo.Holidays) {
		c.JSON(http.StatusOK, 1)
		return
	}
	if isExist(day, holidayInfo.Workdays) {
		c.JSON(http.StatusOK, 0)
		return
	}
	//加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	//加载时区失败
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err})
		return
	}
	today, terr := time.ParseInLocation("20060102", day, loc)
	if terr != nil {
		log.Println(terr)
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请传入正确的时间参数"})
		return
	}
	//判断今天是否为工作日
	switch today.Weekday().String() {
	case "Saturday", "Sunday":
		c.JSON(http.StatusOK, 1)
		return
	default:
		c.JSON(http.StatusOK, 0)
		return
	}
}
