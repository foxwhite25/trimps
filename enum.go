package trimpsGame

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