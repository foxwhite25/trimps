package trimps

import (
	"fmt"
	"math"
)

func (b *Building) GetCost(count float64) map[ResourceName]float64 {
	result := map[ResourceName]float64{}
	countInQue := 0.0
	for _, item := range b.Player.Save.BuildingQue {
		if item.ID == b.ID {
			countInQue += 1
		}
	}
	for resourceName, cost := range b.Cost {
		if cost.Expo == 1 {
			result[resourceName] = count * cost.Base
			continue
		}
		result[resourceName] = math.Floor((cost.Base * math.Pow(cost.Expo, b.Amount+countInQue)) * ((math.Pow(cost.Expo, count)) - 1) / (cost.Expo - 1))
	}
	return result
}

func (b *Building) Buy(count float64) (ok bool) {
	costs := b.GetCost(count)
	for name, cost := range costs {
		if b.Player.Save.Resources[name].Amount < cost {
			return false
		}
	}
	for name, cost := range costs {
		b.Player.Save.Resources[name].Amount -= cost
	}
	b.Player.Save.AddBuildingQue(b)
	return true
}

func (b *Building) Description() string {
	return b.DescriptionFunc(b.Player)
}

func getDefaultBuildings(player *Player) map[BuildingName]*Building {
	return map[BuildingName]*Building{
		Trap: {
			Player:    player,
			Locked:    true,
			ID:        Trap,
			Cost:      map[ResourceName]Cost{Wood: {10, 1}, Food: {10, 1}},
			BuildTime: 5,
			Name:      "陷阱",
			DescriptionFunc: func(player *Player) string {
				return fmt.Sprintf("每个陷阱能让你抓住 %d 个东西", 10)
			},
		},
	}
}

func (s *Save) AddBuildingQue(build *Building) {
	s.BuildingQue = append(s.BuildingQue, BuildingQueItem{
		ID:       build.ID,
		TimeLeft: build.BuildTime,
	})
}

func (s *Save) TickBuildingQue() {
	if len(s.BuildingQue) == 0 {
		return
	}
	if s.BuildingQue[0].TimeLeft < 0.1 {
		building := s.Buildings[s.BuildingQue[0].ID]
		if !s.Flag.Has(JobFlag) && building.ID == Trap {
			s.Flag.Set(JobFlag)
			s.Jobs[Lumberjack].Locked = false
			s.Jobs[Farmer].Locked = false
			s.Player.PendingMessage = append(s.Player.PendingMessage, "陷阱里有一个看起来很皮的脆弱小动物，你决定叫它“脆皮”。要不再有创意一些，考虑一下训练脆皮来帮你干活？")
		}
		building.Amount += 1
		building.Built += 1
		s.BuildingQue = s.BuildingQue[1:]
	} else {
		s.BuildingQue[0].TimeLeft -= 0.1
	}
}
