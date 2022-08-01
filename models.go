package trimpsGame

type Player struct {
	Save Save
}

type Save struct {
	Player       *Player                   `json:"-"`
	PlayerAction Action                    `json:"playerAction"`
	Resources    map[ResourceName]Resource `json:"resources"`
	Buildings    map[BuildingName]Building `json:"buildings"`
}

type Resource struct {
	Player   *Player                      `json:"-"`
	Amount   float64                      `json:"amount"`
	MaxFunc  func(player *Player) float64 `json:"-"`
	RateFunc func(player *Player) float64 `json:"-"`
	Name     string                       `json:"-"`
}

type TrimpsData Resource

type Building struct {
	Player          *Player                     `json:"-"`
	Amount          float64                     `json:"amount"`
	Built           float64                     `json:"built"`
	Locked          bool                        `json:"locked"`
	Cost            map[ResourceName]Cost       `json:"-"`
	BuildTime       float64                     `json:"-"`
	Name            string                      `json:"-"`
	DescriptionFunc func(player *Player) string `json:"-"`
}

type Cost struct {
	Base float64
	Expo float64
}
