package varnames

import (
	"fmt"

	"github.com/gnboorse/centipede"
)

type VariablePrefix string

const (
	PrefixPlacement VariablePrefix = "placement_"
)

func Placement(itemId int) centipede.VariableName {
	return centipede.VariableName(fmt.Sprintf("%s%d", PrefixPlacement, itemId))
}
