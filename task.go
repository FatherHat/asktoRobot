package main

import (
	"fmt"
	"mrfzRObot/Tool"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/bitmap"
)

var path = [3]string{
	"icon/xiuxing/", //修行
	"icon/shijue/",  //十绝
	"icon/xunxian/", //寻仙
} //修行有关图片路径

/*
*
健值0 是修行
健值1 是十绝
健值2 是寻仙
*
*/
var target = [3][]string{
	{"leishen", "huashen", "longshen", "yanshen", "shanshen"},                         //修行
	{"hongshazhenzhu", "fengkongzhenzhu"},                                             //十绝阵
	{"tieguali", "ruidongbing", "hanxiangzi", "zhangguolao", "hexiangu", "caoguojiu"}, //寻仙
}
var offset = [5]int{67, 67, 82, 67, 82} //任务目标对话偏移距离，对应上面怪物顺序

// 领取任务的NPC
var receiveNPC = [3]string{
	"liuruchen",      //修行
	"nanhuazhenren",  //十绝
	"taiqingzhenjun", //寻仙
}

// 领取任务NPC对话鼠标偏移量
var recOffset = [3]int{
	106, //修行
	106, //十绝
	106, //寻仙
}

var offsetSizeOfTask = 5 //点击任务栏图片x轴和y轴的偏移量

var offsetSizeX = 10 //对话鼠标x轴可偏移量
var offsetSizeY = 5  //对话鼠标y轴可偏移量
var imgType = "png"  //图片类型

// 点击任务目标进行引导追踪
func ClickTaskTrace(runType int, fpid int) bool {
	//指定进程为当前活动窗口
	robotgo.ActivePid(fpid)
	//Alt+q 打开任务栏
	arr := []string{"alt", "command"}
	robotgo.KeyTap("q", arr)
	//挂起线程随机时间
	Tool.SleepRandTime()
	//打开任务栏自选界面
	openMySelectTask()
	//挂起线程随机时间
	Tool.SleepRandTime()
	target := target[runType]
	for key := range target {
		path := path[runType] + "rw_" + target[key] + "." + imgType
		x, y := bitmap.FindPic(path)
		if x != -1 && y != -1 {
			//挂起线程随机时间
			Tool.SleepRandTime()
			//增加偏移量，随机点击位置
			x += Tool.GetRandOfRange(offsetSizeOfTask)
			y += Tool.GetRandOfRange(offsetSizeOfTask)
			Tool.MoveMouse(x, y) //鼠标移动到任务目标
			Tool.SleepRandTime() //挂起线程随机时间
			robotgo.Click()      //点击任务目标进行引导追踪
			fmt.Println("任务进行追踪!")
			return true
		}
	}
	return false
}

// 是否显示NPC对话
func IsDisplayDialog(runType int) bool {
	target := target[runType]
	//遍历修行的所有怪
	for i := 0; i < len(target); i++ {
		path := path[runType] + "dh_" + target[i] + "." + imgType
		res := Tool.IsExistForImg(path)
		//如果存在对话跳出循环，跳到breakHere标签
		if res {
			goto breakHere
		}
	}
	return false
breakHere:
	return true
}

// 移动到NPC处，领取任务
func GoToReceiveTask(runType int, fpid int) bool {
	//设为活动窗口
	robotgo.ActivePid(fpid)
	//挂起线程随机时间
	Tool.SleepRandTime()
	//alt+q 打开任务栏
	arr := []string{"alt", "command"}
	robotgo.KeyTap("q", arr)
	//挂起线程随机时间
	Tool.SleepRandTime()
	//打开任务栏自选界面
	openMySelectTask()
	//挂起线程随机时间
	Tool.SleepRandTime()
	//领取任务的NPC图片
	path := path[runType] + "rw_" + receiveNPC[runType] + "." + imgType
	x, y := bitmap.FindPic(path)
	if x != -1 && y != -1 {
		//增加偏移量，随机点击位置
		x += Tool.GetRandOfRange(offsetSizeOfTask)
		y += Tool.GetRandOfRange(offsetSizeOfTask)
		Tool.MoveMouse(x, y) //鼠标移动到任务目标
		Tool.SleepRandTime() //挂起线程随机时间
		robotgo.Click()      //点击任务目标进行引导追踪
		fmt.Println("前去NPC处领取任务...")
		return true
	}
	return false
}

// 打开自选任务栏界面
func openMySelectTask() {
	selectOn := "icon/task_select_off.png"
	x, y := bitmap.FindPic(selectOn)
	//没有打开自选任务栏
	if x != -1 && y != -1 {
		//增加偏移量，随机点击位置
		x += Tool.GetRandOfRange(5)
		y += Tool.GetRandOfRange(10)
		Tool.MoveMouse(x, y) //鼠标移动到任务目标
		Tool.SleepRandTime() //挂起线程随机时间
		robotgo.Click()      //点击任务目标进行引导追踪
		fmt.Println("打开任务栏自选界面！")
	}
}

// 是否显示领取任务NPC对话
func IsDisplayRecNpc(runType int) bool {
	path := path[runType] + "dh_" + receiveNPC[runType] + "." + imgType
	res := Tool.IsExistForImg(path)
	//如果存在对话跳出循环，跳到breakHere标签
	if res {
		return true
	}
	return false
}

// 领取任务
func TalkReceiveTask(fpid int, runType int) {
	robotgo.ActivePid(fpid)
	path := path[runType] + "dh_" + receiveNPC[runType] + "." + imgType
	x, y := bitmap.FindPic(path)
	if x != -1 && y != -1 {
		fmt.Println("到达任务目标位置...")
		//找到目标图片位置                                       //挂起线程随机时间
		x += Tool.GetRandOfRange(offsetSizeX)                      //x轴增加随机量
		y += recOffset[runType] + Tool.GetRandOfRange(offsetSizeY) //y轴增加随机量
		Tool.MoveMouse(x, y)                                       //移动到对话坐标
		Tool.SleepRandTime()                                       //挂起线程随机时间
		robotgo.Click()                                            //选择难度
		Tool.SleepRandTime()                                       //挂起线程随机时间
		robotgo.Click()                                            //确认领取
		Tool.SleepRandTime()                                       //挂起线程随机时间
		robotgo.Click()                                            //点击关闭对话
		//点击全民升级取消按钮
		Tool.SleepRandTime()
		ClickCancel()
	}

}

// 点击NPC对话，进入战斗状态
func ClickTaskTarget(fpid int, runType int) {
	//设为活动窗口
	robotgo.ActivePid(fpid)
	target := target[runType]
	//遍历修行的所有怪
	for key := range target {
		path := path[runType] + "dh_" + target[key] + "." + imgType
		x, y := bitmap.FindPic(path)
		if x != -1 && y != -1 {
			fmt.Println("到达任务目标位置...")
			//找到目标图片位置
			Tool.SleepRandTime()                                //挂起线程随机时间
			x += Tool.GetRandOfRange(offsetSizeX)               //x轴增加随机量
			y += offset[key] + Tool.GetRandOfRange(offsetSizeY) //y轴增加随机量
			Tool.MoveMouse(x, y)                                //移动到对话坐标
			Tool.SleepRandTime()                                //挂起线程随机时间
			fmt.Println("进入战斗!")
			robotgo.Click() //点击对话进入战斗
			break
		}
	}
}

// 点击取消按钮
func ClickCancel() {
	path := "icon/cancel.png"
	x, y := bitmap.FindPic(path)
	if x != -1 && y != -1 {
		//找到目标图片位置
		x += Tool.GetRandOfRange(3) //x轴增加随机量
		y += Tool.GetRandOfRange(6) //y轴增加随机量
		Tool.MoveMouse(x, y)        //移动到对话坐标
		Tool.SleepRandTime()        //挂起线程随机时间
		robotgo.Click()             //选择难度
	}
}
