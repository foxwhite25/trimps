package trimps

func (r *Resource) Reduce(count float64) (ok bool) {
	if count < 0 {
		return r.Add(count * -1)
	}
	if r.Amount-count < 0 {
		return false
	} else {
		r.Amount -= count
		return true
	}
}

func (r *Resource) Add(count float64) (ok bool) {
	if count < 0 {
		return r.Reduce(count * -1)
	}
	max := r.MaxFunc(r.Player)
	if r.Amount+count > max && max != -1 {
		r.Amount = max
		return false
	} else {
		r.Amount += count
		return true
	}
}

func (r *Resource) Tick() {
	r.Add(r.RateFunc(r.Player))
}

func (r *Resource) with(name string, max func(player *Player) float64, rate func(player *Player) float64) *Resource {
	r.RateFunc = rate
	r.MaxFunc = max
	r.Name = name
	return r
}

func getDefaultResources(player *Player) map[ResourceName]*Resource {
	return map[ResourceName]*Resource{
		Food: {
			Player: player,
			Name:   "食物",
			MaxFunc: func(player *Player) float64 {
				return 500
			},
			RateFunc: func(player *Player) float64 {
				if player.Save.PlayerAction == Farming {
					return 1
				} else {
					return 0
				}
			},
		},
		Wood: {
			Player: player,
			Name:   "木头",
			Locked: true,
			MaxFunc: func(player *Player) float64 {
				return 500
			},
			RateFunc: func(player *Player) float64 {
				if player.Save.PlayerAction == Lumbering {
					return 1
				} else {
					return 0
				}
			},
		},
		Metal: {
			Player: player,
			Name:   "钢铁",
			Locked: true,
			MaxFunc: func(player *Player) float64 {
				return 500
			},
			RateFunc: func(player *Player) float64 {
				if player.Save.PlayerAction == Mining {
					return 1
				} else {
					return 0
				}
			},
		},
		Science: {
			Player: player,
			Name:   "科研",
			Locked: true,
			MaxFunc: func(player *Player) float64 {
				return -1
			},
			RateFunc: func(player *Player) float64 {
				if player.Save.PlayerAction == Researching {
					return 1
				} else {
					return 0
				}
			},
		},
	}
}
