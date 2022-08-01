package trimpsGame

import (
	"encoding/json"
	"io"
)

// LoadPlayer read json file from path to load saves
func LoadPlayer(r io.Reader) (player Player, err error) {
	player = Player{}
	save := getDefaultSave(&player)
	decoder := json.NewDecoder(r)
	err = decoder.Decode(&save)
	player.Save = save
	return
}

// SavePlayer save player data to json
func SavePlayer(player Player, r io.Writer) (err error) {
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
	}
}
