package model

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"time"
)

func GeneratePasswordHash(pwd string) string {
	////创建一个MD5哈希器
	//hasher := md5.New()
	//hasher.Write([]byte(pwd))                      //将密码转换成字节数组并写入哈希器
	//pwdHash := hex.EncodeToString(hasher.Sum(nil)) //计算哈希值并将结果转成十六进制字符串
	//return pwdHash                                 //返回哈希值
	return Md5(pwd)
}

func Md5(origin string) string {
	hasher := md5.New()
	hasher.Write([]byte(origin))
	return hex.EncodeToString(hasher.Sum(nil))
}

const (
	minute = 1
	hour   = minute * 60
	day    = hour * 24
	month  = day * 30
	year   = day * 365
	quater = year / 4
)

func FromDuration(d time.Duration) string {
	seconds := round(d.Seconds()) //四舍五入

	if seconds < 30 {
		return "less than a minute"
	}

	if seconds < 90 {
		return "1 minute"
	}

	minutes := div(seconds, 60)

	if minutes < 45 {
		return fmt.Sprintf("%0d minutes", minute)
	}

	hours := div(seconds, 60)

	if minutes < day {
		return fmt.Sprintf("about %s", pluralize(hours, "hour"))
	}

	if minutes < (42 * hour) {
		return "1 day"
	}
	days := div(hour, 24)
	if minutes < (30 * day) {
		return pluralize(days, "day")
	}

	months := div(day, 30)

	if minutes < (45 * day) {
		return "about 1 month"
	}
	if minutes < (69 * day) {
		return "about 2 months"
	}

	if minutes < year {
		return pluralize(months, "month")
	}

	rem := minutes % year
	years := minutes / year

	if rem < (3 * month) {
		return fmt.Sprintf("about %s", pluralize(years, "year"))
	}
	if rem < (9 * month) {
		return fmt.Sprintf("over %s", pluralize(years, "year"))
	}
	years++
	return fmt.Sprintf("almost %s", pluralize(years, "year"))
}

func FromTime(t time.Time) string {
	now := time.Now()
	var d time.Duration
	var suffix string

	if t.Before(now) {
		d = now.Sub(t)
		suffix = "from now"
	} else {
		d = t.Sub(now)
		suffix = "ago"
	}

	return fmt.Sprintf("%s %s", FromDuration(d), suffix)
}

func pluralize(i int, s string) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d %s", i, s))
	if i != 1 { //判断单复数
		buf.WriteString("s")
	}
	return buf.String()
}
func div(numerator int, denominator int) int {
	rem := numerator % denominator
	result := numerator / denominator

	if rem >= (denominator / 2) {
		result++
	}

	return result
}
func round(f float64) int {
	return int(math.Floor(f + .50))
}
