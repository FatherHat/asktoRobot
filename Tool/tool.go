package Tool

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/bitmap"
)

func IsExistForImg(path string) bool {
	x, y := bitmap.FindPic(path)
	if x == -1 && y == -1 {
		return false
	}
	return true
}

// 因为鼠标在问道游戏界面直接移动会偏移，所以鼠标需要移出到游戏界面才在移动
func MoveMouse(x int, y int) {
	robotgo.Move(0, 0) //鼠标移出问道游戏界面
	SleepRandTime()
	robotgo.MoveSmooth(x, y, 0.4, 0.65) //移动到对话坐标
}

// 挂起线程随机时间
func SleepRandTime() {
	num := GetDelayTime() //线程挂起时间,0~1.25秒
	time.Sleep(time.Duration(num) * time.Millisecond)
	time.Sleep(time.Duration(num) * time.Millisecond)
}

// 返回100~750随机数
func GetDelayTime() int {
	rand.Seed(time.Now().UnixNano()) //设置随机数种子，time.Nw().UnixNano()返回的是当前操作系统时间的毫秒
	num := rand.Intn(650)            //返回0~650以内的随机
	return num + 100
}

// 返回0~num随机
func GetRandOfRange(num int) int {
	rand.Seed(time.Now().UnixNano()) //设置随机数种子,timeNow().UnixNano()返回的是当前操作系统时间的毫秒值
	res := rand.Intn(num)            //返回0~num以内的随机
	return res
}

// 返回数组中最大的值
func FindMaxElement(arr []int) int {
	max_num := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max_num {
			max_num = arr[i]
		}
	}
	return max_num
}

// 返回数组中最小的值
func FindMinElement(arr []int) int {
	min_num := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min_num {
			min_num = arr[i]
		}
	}
	return min_num
}
