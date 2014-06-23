package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

const processName string = "calc.exe"

type processInfo struct {
	Name string
	pid  int
}

func main() {
	cmd_re, err := exec.Command("tasklist").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", cmd_re)
	re := regexp.MustCompile("(.*)+\n")
	psList := re.FindAllString(string(cmd_re), -1)
	//fmt.Println(psList)
	//findProcess(psList)
	pInfo, err := findProcess(psList)
	if err != nil {
		//log.Fatal(err)
		err := run()
		if err == nil {
			fmt.Println(processName + " start up success")
		} else {
			fmt.Println(processName + " start up fail")
		}
	}
	fmt.Println(pInfo)

	times := nowUnix()
	fmt.Println(times)
	//nowUnix()

}

//返回当前时间戳
func nowUnix() (times int64) {
	timeNow := time.Now().Unix()
	return timeNow
}

//在进程列表寻找制定进程
func findProcess(cmd_result []string) (info processInfo, err error) {
	var p processInfo
	for i, n := range cmd_result {
		fmt.Println(i, n)
		re := regexp.MustCompile("([A-Za-z0-9.,-])+")
		slice := re.FindAllString(string(n), -1)
		if len(slice) > 1 { //过滤说明信息
			if slice[0] == processName {
				p.Name = slice[0]
				p.pid, _ = strconv.Atoi(slice[1])
				return p, nil
			}
		}
	}
	return p, errors.New("can not find " + processName)

}

//启动进程
func run() (err error) {
	re := exec.Command(processName).Start()
	if re != nil {
		return errors.New(processName + "start up fail")
	} else {
		return nil //启动成功
	}
}
