package trimps

type Action int64

const (
	None Action = iota
	Farming
	Lumbering
	Mining
	Researching
	Trapping
)

type ResourceName int64

const (
	Food ResourceName = iota
	Wood
	Metal
	Science
)

type BuildingName int64

const (
	Trap BuildingName = iota
)

type JobName int64

const (
	Farmer = iota
	Lumberjack
	Miner
)

type Flag uint8

const (
	FoodFlag Flag = 1 << iota
	WoodFlag
	TrapFlag
	JobFlag
	ResearchFlag
	BreedFlag
)

func (b *Flag) Set(flag Flag)     { *b = *b | flag }
func (b *Flag) Clear(flag Flag)   { *b = *b &^ flag }
func (b *Flag) Toggle(flag Flag)  { *b = *b ^ flag }
func (b Flag) Has(flag Flag) bool { return b&flag != 0 }
