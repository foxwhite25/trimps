package trimps

import "fmt"

func getDefaultJobs(player *Player) map[JobName]*Job {
	return map[JobName]*Job{
		Farmer: {
			Player: player,
			Cost:   map[ResourceName]Cost{Food: {10, 1}},
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
