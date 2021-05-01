package block

import (
	"github.com/df-mc/dragonfly/server/item/tool"
)

// Glass is a decorative, fully transparent solid block that can be dyed into stained glass.
type Glass struct {
	solid
	transparent
	clicksAndSticks
}

// BreakInfo ...
func (g Glass) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 0.3,
		Drops:    simpleDrops(),
		Harvestable: func(t tool.Tool) bool {
			return true
		},
		Effective: nothingEffective,
	}
}

// EncodeItem ...
func (Glass) EncodeItem() (id int32, name string, meta int16) {
	return 20, "minecraft:glass", 0
}

// EncodeBlock ...
func (Glass) EncodeBlock() (string, map[string]interface{}) {
	return "minecraft:glass", nil
}