package incoming

import (
	"github.com/Steve-Nzr/go-ff/pkg/feature/inventory/helpers"
	"github.com/Steve-Nzr/go-ff/pkg/service/cache"
	"github.com/Steve-Nzr/go-ff/pkg/service/external"
)

// EquipItem packet
func EquipItem(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	uniqueID := p.Packet.ReadUInt8()
	part := p.Packet.ReadInt8()

	helpers.Equip(player, uniqueID, part)
}

// MoveItem packet
func MoveItem(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	p.Packet.ReadUInt8() // skipped
	sourceSlot := p.Packet.ReadUInt8()
	destSlot := p.Packet.ReadUInt8()

	helpers.Move(player, sourceSlot, destSlot)
}

func DropItem(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	p.Packet.ReadUInt32()
	uniqueID := p.Packet.ReadUInt32()
	count := p.Packet.ReadInt16()
	helpers.Drop(player, uniqueID, count)
}
