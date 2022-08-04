package trimps

import (
	"fmt"
	"math"
)

func (b *Building) GetCost(count float64) map[ResourceName]float64 {
	result := map[ResourceName]float64{}
	for resourceName, cost := range b.Cost {
		if cost.Expo == 1 {
			result[resourceName] = count * cost.Base
			continue
		}
		result[resourceName] = math.Floor((cost.Base * math.Pow(cost.Expo, b.Amount)) * ((math.Pow(cost.Expo, count)) - 1) / (cost.Expo - 1))
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
	b.Built += 1
	b.Amount += 1
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
			Cost:      map[ResourceName]Cost{Wood: {10, 1}, Food: {10, 1}},
			BuildTime: 5,
			Name:      "陷阱",
			DescriptionFunc: func(player *Player) string {
				return fmt.Sprintf("每个陷阱能让你抓住 %d 个东西", 10)
			},
		},
	}
}
