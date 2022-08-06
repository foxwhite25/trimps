package trimps

import "time"

type Player struct {
	Save           Save
	PendingMessage []string
}

type BuildingQueItem struct {
	ID       BuildingName
	TimeLeft float64
}

type BuildingQue []BuildingQueItem

type Save struct {
	Player       *Player                    `json:"-"`
	LastSave     time.Time                  `json:"lastSave"`
	PlayerAction Action                     `json:"playerAction"`
	Resources    map[ResourceName]*Resource `json:"resources"`
	Buildings    map[BuildingName]*Building `json:"buildings"`
	Jobs         map[JobName]*Job           `json:"jobs"`
	Trimps       TrimpsData                 `json:"trimps"`
	Stats        Stats                      `json:"stats"`
	Flag         Flag                       `json:"flag"`
	BuildingQue  BuildingQue                `json:"buildingQue"`
}

type Stats struct {
	TotalJob float64 `json:"totalJob"`
}

type Resource struct {
	Player   *Player                      `json:"-"`
	Amount   float64                      `json:"amount"`
	Locked   bool                         `json:"locked"`
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
	ID              BuildingName                `json:"-"`
	DescriptionFunc func(player *Player) string `json:"-"`
}

type Cost struct {
	Base float64
	Expo float64
}

type Job struct {
	Player          *Player                     `json:"-"`
	Amount          float64                     `json:"amount"`
	Locked          bool                        `json:"locked"`
	Cost            map[ResourceName]Cost       `json:"-"`
	Name            string                      `json:"-"`
	DescriptionFunc func(player *Player) string `json:"-"`
}
