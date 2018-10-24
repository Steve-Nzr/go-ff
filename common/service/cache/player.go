package cache

import (
	"flyff/common/def/component"
)

// Player structure representing it's State
type Player struct {
	EntityID    uint32 `gorm:"unique_index"`
	NetClientID uint32 `gorm:"primary_key"`
	Type        uint8
	Name        string `gorm:"size:64"`
	Gender      uint8
	ModelID     uint32
	Position    component.Position `gorm:"embedded;EMBEDDED_PREFIX:posit_"`
	Angle       float32
	Size        uint8
	Level       uint16
	PlayerID    uint32
	Slot        uint8
	JobID       uint8
	HairColor   uint32
	HairID      uint32
	SkinSetID   uint32
	FaceID      uint32
	Statistics  component.Statistics `gorm:"embedded;EMBEDDED_PREFIX:stats_"`
	//Moving      mc.Moving            `gorm:"embedded;EMBEDDED_PREFIX:movin_"`
}
