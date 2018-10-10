package packets

import "flyff/core"

// CreatePlayer packet struct
type CreatePlayer struct {
	Username          string
	Password          string
	Slot              uint8
	Name              string
	FaceID            int8
	CostumeID         int8
	SkinSet           int8
	HairMeshID        int8
	HairColor         uint32
	Gender            uint8
	Job               uint8
	HeadMesh          int8
	BankPassword      int32
	AuthenticationKey int32
}

// DeletePlayer packet struct
type DeletePlayer struct {
	Username        string
	Password        string
	PasswordConfirm string
	CharacterID     uint32
}

// PreJoin packet struct
type PreJoin struct {
	Username      string
	CharacterID   uint32
	CharacterName string
}

// Construct ...
func (c *CreatePlayer) Construct(p *core.Packet) {
	c.Username = p.ReadString()
	c.Password = p.ReadString()
	c.Slot = p.ReadUInt8()
	c.Name = p.ReadString()
	c.FaceID = p.ReadInt8()
	c.CostumeID = p.ReadInt8()
	c.SkinSet = p.ReadInt8()
	c.HairMeshID = p.ReadInt8()
	c.HairColor = p.ReadUInt32()
	c.Gender = p.ReadUInt8()
	c.Job = p.ReadUInt8()
	c.HeadMesh = p.ReadInt8()
	c.BankPassword = p.ReadInt32()
	c.AuthenticationKey = p.ReadInt32()
}

// Construct ...
func (d *DeletePlayer) Construct(p *core.Packet) {
	d.Username = p.ReadString()
	d.Password = p.ReadString()
	d.PasswordConfirm = p.ReadString()
	d.CharacterID = p.ReadUInt32()
}

// Construct ...
func (pj *PreJoin) Construct(p *core.Packet) {
	pj.Username = p.ReadString()
	pj.CharacterID = p.ReadUInt32()
	pj.CharacterName = p.ReadString()
}
