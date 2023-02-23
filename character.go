package main

import (
	"fmt"
	"mrfzRObot/Tool"
	"os"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/bitmap"
)

var (
	// 人物图片路径
	characterImgPath = "icon/character/"
	//人物头像鼠标偏移量
	characterImgOffset = 10
	//自动按钮随机偏移量
	characterAtuoOffsetX = 6
	characterAtuoOffsetY = 3
)

// 回复状态
func recoilStatus() {
	//打开文件夹
	f, err := os.Open(characterImgPath)
	if err != nil {
		panic(err)
	}
	//读取文件内的所有文件
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		panic(err)
	}
	//遍历文件夹所有图片
	for _, file := range files {
		path := characterImgPath + file.Name()
		x, y := bitmap.FindPic(path)
		if x != -1 && y != -1 {
			//挂起线程随机时间
			Tool.SleepRandTime()
			//增加偏移量，随机点击位置
			x += Tool.GetRandOfRange(characterImgOffset) + 20
			y += Tool.GetRandOfRange(characterImgOffset) + 20
			Tool.MoveMouse(x, y) //鼠标移动到任务目标
			//挂起线程随机时间
			Tool.SleepRandTime()
			robotgo.Click("right")
			//挂起线程随机时间
			Tool.SleepRandTime()
			robotgo.Move(0, 0) //鼠标移出问道游戏界面
			fmt.Println("回复人物状态!")
			break
		}
	}
}

// 续上自动
func renewalAtuo(fpid int) {
	robotgo.ActivePid(fpid)
	Tool.SleepRandTime()
	path := "icon/auto.png"
	x, y := bitmap.FindPic(path)
	if x != -1 && y != -1 {
		//挂起线程随机时间
		Tool.SleepRandTime()
		//增加偏移量，随机点击位置
		x += Tool.GetRandOfRange(characterAtuoOffsetX)
		y += Tool.GetRandOfRange(characterAtuoOffsetY)
		Tool.MoveMouse(x, y) //鼠标移动到任务目标
		//挂起线程随机时间
		Tool.SleepRandTime()
		robotgo.Click()
		//挂起线程随机时间
		Tool.SleepRandTime()
		fmt.Println("续上自动！编号：", fpid)
		robotgo.Move(0, 0) //鼠标移出问道游戏界面
	}

}
