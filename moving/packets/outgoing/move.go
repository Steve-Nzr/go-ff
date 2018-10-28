package outgoing

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/moving/def/packets"
)

// DestPos packet
func DestPos(p *cache.Player) *external.Packet {
	packet := external.StartMergePacket(p.EntityID, uint16(0x00C1), 0x0000FF00).
		WriteFloat32(float32(p.Moving.Destination.X)).
		WriteFloat32(float32(p.Moving.Destination.Y)).
		WriteFloat32(float32(p.Moving.Destination.Z)).
		WriteUInt8(1)

	return packet.Finalize()
}

// Move packet emitter
func Move(p *cache.Player, b *packets.Behaviour) *external.Packet {
	return external.StartMergePacket(p.EntityID, uint16(0x00ca), 0x0000FF00).
		Write3DVector(b.V).
		Write3DVector(b.Vd).
		WriteFloat32(b.Angle).
		WriteUInt32(b.State).
		WriteUInt32(b.StateFlag).
		WriteUInt32(b.Motion).
		WriteInt32(b.MotionEx).
		WriteInt32(b.Loop).
		WriteUInt32(b.MotionOptions).
		WriteInt64(b.TickCount)
}