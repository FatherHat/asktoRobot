package main

import (
	"fmt"
	"mrfzRObot/Tool"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/bitmap"
)

var (
	// 宠物图片路径
	petImgPath = "icon/pet/"
	//宠物头像鼠标偏移量
	petImgOffset = 10
)

// 回复宠物状态
// func recoilPetStatus() {
// 	//打开文件夹
// 	f, err := os.Open(petImgPath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//读取文件内的所有文件
// 	files, err := f.Readdir(-1)
// 	f.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	//遍历文件夹所有图片
// 	for _, file := range files {
// 		path := petImgPath + file.Name()
// 		x, y := bitmap.FindPic(path)
// 		if x != -1 && y != -1 {
// 			//挂起线程随机时间
// 			Tool.SleepRandTime()
// 			//增加偏移量，随机点击位置
// 			x += Tool.GetRandOfRange(petImgOffset) + 20
// 			y += Tool.GetRandOfRange(petImgOffset) + 20
// 			Tool.MoveMouse(x, y) //鼠标移动到任务目标
// 			//挂起线程随机时间
// 			Tool.SleepRandTime()
// 			robotgo.Click("right")
// 			//挂起线程随机时间
// 			Tool.SleepRandTime()
// 			robotgo.Move(0, 0) //鼠标移出问道游戏界面
// 			fmt.Println("回复人宠物状态!")
// 			break
// 		}
// 	}
// }

// 回复宠物状态
func recoilPetStatus() {
	Tool.SleepRandTime()
	path := "icon/richang.png"
	x, y := bitmap.FindPic(path)
	if x != -1 && y != -1 {
		//挂起线程随机时间
		Tool.SleepRandTime()
		//增加偏移量，随机点击位置
		x += Tool.GetRandOfRange(petImgOffset) - 105
		y += Tool.GetRandOfRange(petImgOffset) - 100
		Tool.MoveMouse(x, y) //鼠标移动到任务目标
		//挂起线程随机时间
		Tool.SleepRandTime()
		robotgo.Click("right")
		Tool.SleepRandTime()
		robotgo.Move(0, 0) //鼠标移出问道游戏界面
		fmt.Println("回复人宠物状态!")
	}
}
