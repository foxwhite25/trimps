package trimps

import (
	"fmt"
	"math"
)

func (j *Job) GetCost(count float64) map[ResourceName]float64 {
	result := map[ResourceName]float64{}
	for resourceName, cost := range j.Cost {
		if cost.Expo == 1 {
			result[resourceName] = count * cost.Base
			continue
		}
		result[resourceName] = math.Floor((cost.Base * math.Pow(cost.Expo, j.Amount)) * ((math.Pow(cost.Expo, count)) - 1) / (cost.Expo - 1))
	}
	return result
}

func (j *Job) Hire(count float64) (ok bool) {
	costs := j.GetCost(count)
	for name, cost := range costs {
		if j.Player.Save.Resources[name].Amount < cost {
			return false
		}
	}
	for name, cost := range costs {
		j.Player.Save.Resources[name].Amount -= cost
	}
	j.Player.Save.Stats.TotalJob += 1
	j.Amount += 1
	return true
}

func (j *Job) Fire(count float64) (ok bool) {
	if count > j.Amount {
		count = j.Amount
	}
	j.Amount -= count
	return true
}

func getDefaultJobs(player *Player) map[JobName]*Job {
	return map[JobName]*Job{
		Farmer: {
			Player: player,
			Cost:   map[ResourceName]Cost{Food: {10, 1}},
			Locked: true,
			Name:   "农民",
			DescriptionFunc: func(player *Player) string {
				return fmt.Sprintf("使用古老的耕作艺术训练一个脆皮。每个农民每秒收获 %.1f 食物。", 0.5)
			},
		},
		Lumberjack: {
			Player: player,
			Cost:   map[ResourceName]Cost{Food: {10, 1}},
			Locked: true,
			Name:   "伐木工",
			DescriptionFunc: func(player *Player) string {
				return fmt.Sprintf("向一个脆皮展示如何砍伐那些怪树。每个伐木工每秒拉回 %.1f 个木头。", 0.5)
			},
		},
		Miner: {
			Player: player,
			Cost:   map[ResourceName]Cost{Food: {10, 1}},
			Locked: true,
			Name:   "矿工",
			DescriptionFunc: func(player *Player) string {
				return fmt.Sprintf("把品行不端的脆皮丢到矿井里干点“放松”的活。每个矿工每秒收获 %.1f 金属。", 0.5)
			},
		},
	}
}
