package trimps

import "math"

func (t *TrimpsData) Reduce(count float64) (ok bool) {
	if count < 0 {
		return t.Add(count * -1)
	}
	if t.Amount-count < 0 {
		return false
	} else {
		t.Amount -= count
		return true
	}
}

func (t *TrimpsData) Add(count float64) (ok bool) {
	if count < 0 {
		return t.Reduce(count * -1)
	}
	max := t.MaxFunc(t.Player)
	if t.Amount+count > max && max != -1 {
		t.Amount = max
		return false
	} else {
		t.Amount += count
		return true
	}
}

func (t *TrimpsData) WorkSpace() float64 {
	return math.Floor(t.Amount / 2)
}

func (t *TrimpsData) Tick() {
	t.Add(t.RateFunc(t.Player) / 10)
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
