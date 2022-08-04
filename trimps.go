package trimps

import "math"

func (t *TrimpsData) WorkSpace() float64 {
	return math.Floor(t.Amount / 2)
}

func getDefaultTrimps(player *Player) TrimpsData {
	return TrimpsData{
		Player: player,
		MaxFunc: func(player *Player) float64 {
			return 10
		},
		RateFunc: func(player *Player) float64 {
			free := player.Save.Trimps.Amount - player.Save.Stats.TotalJob
			if free > 2 {
				return 0.01 * free
			} else {
				return 0
			}
		},
		Name: "脆皮",
	}
}
