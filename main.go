package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func getLuckyNumbers(t int) ([33]int, [16]int, [6]int, int) {
	var redBall int
	var numbers [6]int
	var lucky int
	var redCount [33]int
	var blueCount [16]int

	//红球算法
	var redBalls [33]int
	for i := 0; i < 33; i++ {
		redBalls[i] = i + 1
	}
	for i := 0; i < 33; i++ {
		rand.Seed(time.Now().UnixNano())
		redBall = rand.Intn(33)
		redBalls[i], redBalls[redBall] = redBalls[redBall], redBalls[i]
	}
	for i := 0; i < 6; i++ {
		numbers[i] = redBalls[i]
		redCount[numbers[i]-1]++
	}

	//蓝球算法
	rand.Seed(time.Now().UnixNano() + 1)
	lucky = rand.Intn(16) + 1
	blueCount[lucky-1]++
	switch t {
	case 1:
		fmt.Printf("本期中奖号码为：%02d,%02d\n", numbers, lucky)
	case 2:
		fmt.Printf("机器选择号码为：%02d,%02d\n", numbers, lucky)
	default:
		fmt.Println("无结果")
	}
	return redCount, blueCount, numbers, lucky
}
func getLuckyNumbers2(t int) ([33]int, [16]int, [6]int, int) {
	var redBall int
	var numbers [6]int
	var lucky int
	var redCount [33]int
	var blueCount [16]int

	//红球算法
	for i := 0; i < 6; i++ {
		for {
			var isExist = false
			rand.Seed(time.Now().UnixNano())
			redBall = rand.Intn(33) + 1
			for j := 0; j < 6; j++ {
				if redBall == numbers[j] {
					isExist = true
					break
				}
			}
			if !isExist {
				numbers[i] = redBall
				redCount[redBall-1]++
				break
			}
		}

	}
	//蓝球算法
	rand.Seed(time.Now().UnixNano())
	lucky = rand.Intn(16) + 1
	blueCount[lucky-1]++
	switch t {
	case 1:
		fmt.Printf("本期中奖号码为：%02d,%02d\n", numbers, lucky)
	case 2:
		fmt.Printf("机器选择号码为：%02d,%02d\n", numbers, lucky)
	default:
		fmt.Println("无结果")
	}
	return redCount, blueCount, numbers, lucky
}
func selfChoose() {
	var i, a, b, c, d, e, f, blue, redCorrect, blueCorrect int
	size := 1
	var redCount [33]int
	var blueCount [16]int
	var redCount1 [33]int
	var blueCount1 [16]int
	var winningRed [6]int
	var winningBlue int
	fmt.Println("请输入您选择的六个红球号码（1-33），以空格区分，回车结束：")
	n, err := fmt.Scanln(&a, &b, &c, &d, &e, &f)
	if n == 6 && err == nil && a > 0 && b > 0 && c > 0 && d > 0 && e > 0 && f > 0 && a < 34 && b < 34 && c < 34 && d < 34 && e < 34 && f < 34 {
		var array = [6]int{a, b, c, d, e, f}
		fmt.Println("请再选择一个蓝球号码（1-16）：")
		n, err := fmt.Scanln(&blue)
		if n == 1 && err == nil && blue > 0 && blue < 17 {
			fmt.Printf("您选择的号码为：%02d %02d\n", array, blue)
			for i = 0; i < size; i++ {
				redCount1, blueCount1, winningRed, winningBlue = getLuckyNumbers2(1)
				for j := 0; j < 33; j++ {
					redCount[j] = redCount[j] + redCount1[j]
				}
				for j := 0; j < 16; j++ {
					blueCount[j] = blueCount[j] + blueCount1[j]
				}
				for i := 0; i < 6; i++ {
					for j := 0; j < 6; j++ {
						if array[i] == winningRed[j] {
							redCorrect++
						}
					}
				}
				if blue == winningBlue {
					blueCorrect++
				}
				result := conculatPrize(redCorrect, blueCorrect)
				var detail details
				detail.period = period
				detail.chooseRB = array
				detail.chooseBB = blue
				detail.winningRB = winningRed
				detail.winningBB = winningBlue
				detail.prize = result
				history(1, detail)
			}
		} else {
			fmt.Println("非法数据")
			os.Exit(1)
		}
	} else {
		fmt.Println("非法数据")
		os.Exit(1)
	}
	// fmt.Printf("统计%d次开奖结果各个数出现的次数及概率：\n", size)
	// fmt.Println("蓝球:")
	// for i = 0; i < 33; i++ {
	// 	var red = float32(redCount[i])
	// 	var size = float32(size)
	// 	fmt.Printf("[%d]:%d,%.01f%%\n", i+1, redCount[i], red/size*100)
	// }
	// fmt.Println("红球")
	// for i = 0; i < 16; i++ {
	// 	var blue = float32(blueCount[i])
	// 	var size = float32(size)
	// 	fmt.Printf("[%d]:%d,%.01f%%\n", i+1, blueCount[i], blue/size*100)
	// }
}
func machineChoose(t, index int) {
	var redCorrect, blueCorrect int
	if t == 1 {
		fmt.Printf("【第%d期】\n", index)
	}
	_, _, array, blue := getLuckyNumbers2(2)
	_, _, winningRed, winningBlue := getLuckyNumbers2(1)
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if array[i] == winningRed[j] {
				redCorrect++
			}
		}
	}
	if blue == winningBlue {
		blueCorrect++
	}
	result := conculatPrize(redCorrect, blueCorrect)
	var detail details
	detail.period = period
	detail.chooseRB = array
	detail.chooseBB = blue
	detail.winningRB = winningRed
	detail.winningBB = winningBlue
	detail.prize = result
	history(1, detail)
}
func conculatPrize(red, blue int) string {
	var result string
	switch blue {
	case 0:
		if red == 6 {
			result = "二等奖\n"
		} else if red == 5 {
			result = "四等奖\n"
		} else if red == 4 {
			result = "五等奖\n"
		} else {
			result = "您未中奖，感谢您对福利事业的支持！\n"
		}
	case 1:
		if red == 6 {
			result = "一等奖\n"
		} else if red == 5 {
			result = "三等奖\n"
		} else if red == 4 {
			result = "四等奖\n"
		} else if red == 3 {
			result = "五等奖\n"
		} else {
			result = "六等奖\n"
		}
	default:
		result = "非法数据\n"
	}
	fmt.Println(result)
	return result
}
func conculatProbability(list []details) ([6]int, int) {
	var red [33]int
	var blue [16]int
	var newRed [6]int
	var newBlue int
	oldRed := red
	oldBlue := blue
	for i := 0; i < 33; i++ {
		for j := 32; j > i; j-- {
			if red[i] < red[j] {
				red[i], red[j] = red[j], red[i]
			}
		}
	}
	for i := 0; i < 16; i++ {
		for j := 15; j > i; j-- {
			if blue[i] < blue[j] {
				blue[i], blue[j] = blue[j], blue[i]
			}
		}
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			for k := 0; k < 33; k++ {
				if red[i] == oldRed[k] {
					newRed[i] = k
				}
			}

		}
	}
	for i := 0; i < 16; i++ {
		if oldBlue[i] == blue[0] {
			newBlue = i
		}
	}
	fmt.Println("概率最高的红球和蓝球分别是:")
	fmt.Println(newRed)
	fmt.Println(newBlue)
	return newRed, newBlue
}
func history(t int, detail details) {
	switch t {
	case 1:
		list = append(list, detail)
		period++
	case 2:
		for i := 0; i < period-1; i++ {
			fmt.Printf("期数:%02d	选择号码:%02d,%02d	中奖号码:%02d,%02d	中奖情况:%s\n", list[i].period, list[i].chooseRB, list[i].chooseBB, list[i].winningRB, list[i].winningBB, list[i].prize)
		}
		fmt.Printf("%d次抽奖数据汇总:\n", period)
		var a [6]int
		for i := 0; i < period-1; i++ {
			if list[i].prize == "一等奖\n" {
				a[0] = a[0] + 1
			} else if list[i].prize == "二等奖\n" {
				a[1] = a[1] + 1
			} else if list[i].prize == "三等奖\n" {
				a[2] = a[2] + 1
			} else if list[i].prize == "四等奖\n" {
				a[3] = a[3] + 1
			} else if list[i].prize == "五等奖\n" {
				a[4] = a[4] + 1
			} else if list[i].prize == "六等奖\n" {
				a[5] = a[5] + 1
			}
		}
		fmt.Printf("一等奖(%d)	二等奖(%d)	三等奖(%d)	四等奖(%d)	五等奖(%d)	六等奖(%d)\n", a[0], a[1], a[2], a[3], a[4], a[5])
		fmt.Println()
	case 3:
		fmt.Printf("%d次抽奖数据汇总:\n", detail.period)
		var a [6]int
		for i := 0; i < period-1; i++ {
			if list[i].prize == "一等奖\n" {
				a[0] = a[0] + 1
			} else if list[i].prize == "二等奖\n" {
				a[1] = a[1] + 1
			} else if list[i].prize == "三等奖\n" {
				a[2] = a[2] + 1
			} else if list[i].prize == "四等奖\n" {
				a[3] = a[3] + 1
			} else if list[i].prize == "五等奖\n" {
				a[4] = a[4] + 1
			} else if list[i].prize == "六等奖\n" {
				a[5] = a[5] + 1
			}
		}
		fmt.Printf("一等奖(%d)	二等奖(%d)	三等奖(%d)	四等奖(%d)	五等奖(%d)	六等奖(%d)\n", a[0], a[1], a[2], a[3], a[4], a[5])
		fmt.Println()
	default:
		fmt.Println("wrong")
	}
}

type details struct {
	period              int
	chooseRB, winningRB [6]int
	chooseBB, winningBB int
	prize               string
}

var (
	list   []details
	period = 1
)

func main() {
	for {
		time.Sleep(2e9)
		fmt.Println("***双色球***")
		fmt.Print("菜单：")
		fmt.Println(`
	1.自选
	2.机选
	3.模拟多次机选开奖(试手气)
	4.历史数据
	5.退出系统
		`)
		var choseType int
		fmt.Println("请选择:")
		fmt.Scanln(&choseType)
		switch choseType {
		case 1:
			fmt.Println("【自选】")
			selfChoose()
		case 2:
			fmt.Println("【机选】")
			machineChoose(0, 0)
		case 3:
			fmt.Println("请输入需要模拟的开奖次数")
			var n int
			fmt.Scanln(&n)
			if n > n {
				fmt.Println("【模拟开奖】")
				for i := 0; i < n; i++ {
					machineChoose(1, i+1)
				}
				var detail details
				detail.period = n
				history(3, detail)
			} else {
				os.Exit(1)
			}

		case 4:
			fmt.Println("【历史数据】")
			var detail details
			history(2, detail)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("无效输入")
		}
	}
}
