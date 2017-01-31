package rtsengine

/*
 Implements the farm unit

*/

// Farm is an IUnit that maintains a farm and adds food resources to an IPlayer
type Farm struct {
	Poolable
}

func (farm *Farm) name() string {
	return "Farm"
}
