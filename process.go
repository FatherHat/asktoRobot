package main //声明当前包为main包，表示是一个可运行的程序
import (
	"fmt"
	"mrfzRObot/Tool"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/bitmap"
)

// 初始化函数，目的是先打开游戏界面
// func init() {
// 	fpid, err := robotgo.FindIds("asktao")
// 	fmt.Println("pids...", fpid)
// 	if err == nil {
// 		if len(fpid) > 0 {
// 			//打开目标窗口设为活动窗口
// 			robotgo.ActivePid(17248)
// 			//挂起线程随机时间
// 			Tool.SleepRandTime()
// 		} else {
// 			fmt.Println("游戏还没打开呢!")
// 		}
// 	} else {
// 		fmt.Println("出错啦：", err)
// 	}
// }

// 是否回复宠物状态
var isRecPet = 1

func run() {
	var (
		/*
			*
			0：修山
			1：十绝
			2：寻仙
			*
		*/
		runType int
		auto    int // 自动定量，打多少次怪续上自动（增加随机量，实际是atuo+(0~2)次）
	)
	//引导用户输入配置
	fmt.Println("请输入任务的类型（0：修山；1：十绝；2：寻仙）和多少回合后续上自动，空格隔开：")
	_, typeErr := fmt.Scan(&runType, &auto)
	//输出错误
	if typeErr != nil || runType > 2 || auto <= 0 {
		fmt.Println("请输入正确的配置，错误信息：", typeErr)
		return
	}
	//与问道有关的所有进程
	fpid, err := robotgo.FindIds("asktao")
	if err == nil {
		if len(fpid) > 0 {
			//击杀次数
			var i = 0
			//增加的自动轮数，声明一个有缓冲的管道（异步）
			autoAdd := make(chan int, 3)
			for {
				//击杀一个任务怪
				baseProcess(runType)
				i++
				//生产
				if i%auto == 0 && i != 0 && len(autoAdd) <= 0 {
					//增加auto次续上自动
					//randNum := Tool.GetRandOfRange(1) + 1
					randNum := Tool.GetRandOfRange(1)
					//auto进入管道
					for i := 0; i <= randNum; i++ {
						autoAdd <- i
					}
				}
				//消费
				if len(autoAdd) > 0 {
					//最后一轮怪，打完续上自动
					if len(autoAdd) == 1 {
						//清除管道最后一个元素
						<-autoAdd
						//续上所有账号自动
						for key := range fpid {
							//挂起线程随机时间
							Tool.SleepRandTime()
							renewalAtuo(fpid[key])
						}
						//重新计算击杀轮次
						i = 0
					} else {
						//消费一个元素
						<-autoAdd
					}
				}
			}
		}
	}

}

// 脚本的基本流程
func baseProcess(runType int) {

	//与问道有关的所有进程
	fpid, err := robotgo.FindIds("asktao")
	if err == nil {
		if len(fpid) > 0 {

			//队长fid，值最小的是队长
			leader := Tool.FindMinElement(fpid)

			//指定进程为当前活动窗口
			robotgo.ActivePid(leader)

			//挂起线程随机时间
			Tool.SleepRandTime()

			//检测是否正在战斗中，如果正在战斗中就挂起3~4.75秒再次检测，直到战斗结束
			repactDecStatus(leader)
			Tool.SleepRandTime()

			//点击任务栏引导到任务NPC
			isHaveObj := ClickTaskTrace(runType, leader)
			if !isHaveObj {
				fmt.Println("已完成一轮任务!")
				//Alt+q 关闭任务栏
				arr := []string{"alt", "command"}
				robotgo.KeyTap("q", arr)
				//挂起线程随机时间
				Tool.SleepRandTime()

				//重新领取任务
				GoToReceiveTask(runType, leader)

				//随机时间间隔重复检测是否到达领取任务的NPC
				waitToRecNpc(runType)

				//领取任务
				TalkReceiveTask(leader, runType)

				//挂起线程随机时间
				Tool.SleepRandTime()
				//点击任务栏引导到任务NPC
				ClickTaskTrace(runType, leader)
			}

			//挂起线程随机时间
			Tool.SleepRandTime()

			//等待目标显示对话，随机时间间隔重复检测是否到达任务目标
			waitConversation(runType)

			//到达任务目标,点击对话进入战斗
			ClickTaskTarget(leader, runType)

			//挂起线程随机时间
			Tool.SleepRandTime()

			//检测是否正在战斗中，如果正在战斗中就挂起3~4.75秒再次检测，直到战斗结束
			repactDecStatus(leader)

			//战斗完毕回复所有人状态
			for key := range fpid {
				robotgo.ActivePid(fpid[key])
				Tool.SleepRandTime()
				recoilStatus()
			}

			//回复所有宠物状态
			if isRecPet == 1 {
				for key := range fpid {
					robotgo.ActivePid(fpid[key])
					Tool.SleepRandTime()
					recoilPetStatus()
				}
			}

		}
	}
}

// 随机时间间隔重复检测是否到达任务目标，等待角色移动到任务NPC进行对话
func waitConversation(runType int) bool {
	//随机时间间隔重复检测是否到达任务目标
	res := IsDisplayDialog(runType)
	if !res {
		//随机挂起一段时间
		Tool.SleepRandTime()
		fmt.Println("移动中，等待到达任务目标...")
		//再次重复检测
		waitConversation(runType)
	}
	return true
}

// 随机时间间隔重复检测是否到达任务目标，等待角色移动到任务NPC进行对话
func waitToRecNpc(runType int) bool {
	//随机时间间隔重复检测是否到达任务目标
	res := IsDisplayRecNpc(runType)
	if !res {
		//随机挂起一段时间
		Tool.SleepRandTime()
		fmt.Println("正在移动中，前去NPC领取任务...")
		//再次重复检测
		waitToRecNpc(runType)
	}
	return true
}

// 重复检测是否在战斗中，直到战斗结束
func repactDecStatus(obj int) {
	//目标窗口设为活动窗口
	robotgo.ActivePid(obj)
	// 是否在战斗状态
	status := isCombatStatus(obj)
	if status {
		fmt.Println("战斗中...")
		//如果在战斗状态中，挂起3~4.75秒再次检测，直到战斗结束
		Tool.SleepRandTime()
		repactDecStatus(obj)
	} else {
		fmt.Println("战斗结束!")
	}
}

// 是否在战斗状态
func isCombatStatus(obj int) bool {
	//目标窗口设为活动窗口
	robotgo.ActivePid(obj)
	//挂起线程随机时间
	Tool.SleepRandTime()
	//以左上角是否有宠物栏为判断依据，是否在战斗中
	imgPath := "icon/xinfa.png"
	x, y := bitmap.FindPic(imgPath)
	if x == -1 && y == -1 {
		return false
	}
	return true
}
