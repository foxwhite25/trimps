package trimps

import (
	"encoding/json"
	"io"
)

// LoadPlayer read json file from path to load saves
func LoadPlayer(r io.Reader) (player *Player, err error) {
	player = &Player{PendingMessage: []string{}}
	save := getDefaultSave(player)
	decoder := json.NewDecoder(r)
	err = decoder.Decode(&save)
	player.Save = save
	if !player.Save.Flag.Has(FoodFlag) {
		player.PendingMessage = append(player.PendingMessage, "一团绿光闪过后，你落到了地面上。你看起来饿极了……\n使用 /玩家动作 采集 来搞点东西吃。使用 /资源 查看你当前的资源。")
		player.Save.Flag.Set(FoodFlag)
	}
	return
}

// SavePlayer save player data to json
func SavePlayer(player *Player, r io.Writer) (err error) {
	encoder := json.NewEncoder(r)
	err = encoder.Encode(&player.Save)
	return
}

func getDefaultSave(player *Player) Save {
	return Save{
		Player:       player,
		PlayerAction: None,
		Resources:    getDefaultResources(player),
		Buildings:    getDefaultBuildings(player),
		Jobs:         getDefaultJobs(player),
		Trimps:       getDefaultTrimps(player),
		Stats:        Stats{TotalJob: 0},
	}
}

// Tick the game where one tick is 100ms
func (p *Player) Tick(n int) {
	for i := 0; i < n; i++ {
		for _, resource := range p.Save.Resources {
			resource.Tick()
		}
		p.Save.Trimps.Tick()
		p.Save.TickBuildingQue()
	}
}
