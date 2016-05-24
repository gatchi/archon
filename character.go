/*
* Archon PSO Server
* Copyright (C) 2014 Andrew Rodman
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// Constants and structs associated with character data.
package main

// Possible character classes as defined by the game.
type CharClass uint8

const (
	Humar     CharClass = 0x00
	Hunewearl           = 0x01
	Hucast              = 0x02
	Ramar               = 0x03
	Racast              = 0x04
	Racaseal            = 0x05
	Fomarl              = 0x06
	Fonewm              = 0x07
	Fonewearl           = 0x08
	Hucaseal            = 0x09
	Fomar               = 0x0A
	Ramarl              = 0x0B
)

// Per-player friend guildcard entries.
type GuildcardEntry struct {
	Guildcard   uint32
	Name        [24]uint16
	TeamName    [16]uint16
	Description [88]uint16
	Reserved    uint8
	Language    uint8
	SectionID   uint8
	CharClass   uint8
	padding     uint32
	Comment     [88]uint16
}

// Per-player guildcard data chunk.
type GuildcardData struct {
	Unknown  [0x114]uint8
	Blocked  [0x1DE8]uint8 //This should be a struct once implemented
	Unknown2 [0x78]uint8
	Entries  [104]GuildcardEntry
	Unknown3 [0x1BC]uint8
}

// Per-character stats.
type CharacterStats struct {
	ATP uint16
	MST uint16
	EVP uint16
	HP  uint16
	DFP uint16
	TP  uint16
	LCK uint16
	ATA uint16
}

// Common fields for representing a character's appearance.
type CharacterInfo struct {
	NameColorChksm uint32
	SectionID      byte
	CharClass      byte
	V2flags        byte
	Version        byte
	V1Flags        uint32
	Costume        uint16
	Skin           uint16
	Face           uint16
	Head           uint16
	Hair           uint16
	HairRed        uint16
	HairGreen      uint16
	HairBlue       uint16
	PropX          float32
	PropY          float32
	Name           [16]uint16
}

// Item stored in the player's inventory.
type Item struct {
	Equipped uint32
	Flags    uint32
	Data     uint32
	ItemID   uint32
	Data2    uint32 // Only for mags?
}

// A player's inventory.
type Inventory struct {
	NumItems   uint8
	HPMatsUsed uint8
	TPMatsUsed uint8
	Language   uint8
	Items      [30]Item
}

// Items stored in the player's bank.
type BankItem struct {
	Data   uint32
	ItemID uint32
	Data2  uint32
	Amount uint16
	Flags  uint16
}

// A player's bank
type Bank struct {
	NumItems uint32
	Meseta   uint32
	Items    [200]BankItem
}

// Character data sent out to other lobby members.
type Character struct {
	CharacterStats
	Unknown      [8]uint8
	Level        uint32
	Exp          uint32
	Meseta       uint32
	GuildcardStr [24]uint8
	NameColor    uint32
	Model        uint8
	Unused       [11]uint8
	Playtime     uint32
	CharacterInfo
	Config     [232]uint8
	Techniques [20]uint8
}

// Character data sent to the client via the login server when
// selecting a character from the menu.
type CharacterPreview struct {
	Experience   uint32
	Level        uint32
	GuildcardStr [16]byte
	Unknown      [2]uint32
	NameColor    uint32
	Model        byte
	Unused       [15]byte
	CharacterInfo
	Playtime uint32
}

// Full representation of a character, stored identically to the format
// expected by the E7 packet for convenience. This is a hybrid of
// tethealla/newserv/sylverant, incorporating what they agree on and
// making a best guess where they don't.
type FullCharacter struct {
	Inventory
	Character
	Unknown    [16]uint8
	Options    uint32
	QuestData1 [520]uint8
	Bank
	Guildcard     uint32
	Name          [16]uint16
	TeamName      [16]uint16
	GuildcardDesc [88]uint16
	Reserved1     uint8
	Reserved2     uint8
	SectionID     uint8
	CharClass     uint8
	Unknown2      uint32
	SymbolChats   [1248]uint8
	Shortcuts     [2624]uint8
	Autoreply     [172]uint16
	Infoboard     [172]uint16
	Unknown3      [28]uint8
	ChallengeData [320]uint8
	TechMenu      [40]uint8
	Unknown4      [44]uint8
	QuestData2    [88]uint8
	// Team config?
	Unknown5       [276]uint8
	KeyConfig      [364]uint8
	JoystickConfig [56]uint8
	Guildcard2     uint32
	TeamID         uint32
	TeamInfo       [8]uint8
	TeamPrivilege  uint16
	Reserved3      uint16
	TeamName2      [16]uint16
	TeamFlag       [2048]uint8
	TeamRewards    [2]uint32
}

// Default keyboard/joystick configuration used for players who are
// logging in for the first time.
var baseKeyConfig = [420]byte{
	0x00, 0x00, 0x00, 0x00, 0x26, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x61, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x50, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x59, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5d, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5f, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5d, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x5c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5f, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x56, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5e, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x43, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x45, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x46, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x47, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x48, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x49, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x4a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4b, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x2b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2c, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2e, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x2f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x32, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0xff, 0xff,
	0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x08, 0x00,
	0x01, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00,
	0x00, 0x02, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x01, 0x00, 0x00, 0x00,
}

var baseSymbolChats = [1248]byte{
	0x01, 0x00, 0x00, 0x00, 0x09, 0x00, 0x45, 0x00, 0x48, 0x00, 0x65, 0x00, 0x6c, 0x00, 0x6c, 0x00,
	0x6f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0x00, 0xff, 0xff, 0xff, 0xff, 0x05, 0x18, 0x1d, 0x00, 0x05, 0x28, 0x1d, 0x01,
	0x36, 0x20, 0x2a, 0x00, 0x3c, 0x00, 0x32, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02,
	0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0x01, 0x00, 0x00, 0x00, 0x09, 0x00, 0x45, 0x00,
	0x47, 0x00, 0x6f, 0x00, 0x6f, 0x00, 0x64, 0x00, 0x2d, 0x00, 0x62, 0x00, 0x79, 0x00, 0x65, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x74, 0x00, 0x00, 0x00, 0x76, 0x04, 0x0c, 0x00, 0xff, 0xff, 0xff, 0xff,
	0x06, 0x15, 0x14, 0x00, 0x06, 0x2b, 0x14, 0x01, 0x05, 0x18, 0x1f, 0x00, 0x05, 0x28, 0x1f, 0x01,
	0x36, 0x20, 0x2a, 0x00, 0x3c, 0x00, 0x32, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x02,
	0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02,
	0x01, 0x00, 0x00, 0x00, 0x09, 0x00, 0x45, 0x00, 0x48, 0x00, 0x75, 0x00, 0x72, 0x00, 0x72, 0x00,
	0x61, 0x00, 0x68, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00,
	0x62, 0x03, 0x62, 0x03, 0xff, 0xff, 0xff, 0xff, 0x09, 0x16, 0x1b, 0x00, 0x09, 0x2b, 0x1b, 0x01,
	0x37, 0x20, 0x2c, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02,
	0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0x01, 0x00, 0x00, 0x00, 0x09, 0x00, 0x45, 0x00,
	0x43, 0x00, 0x72, 0x00, 0x79, 0x00, 0x69, 0x00, 0x6e, 0x00, 0x67, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x74, 0x00, 0x00, 0x00, 0x4f, 0x07, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x06, 0x15, 0x14, 0x00, 0x06, 0x2b, 0x14, 0x01, 0x05, 0x18, 0x1f, 0x00, 0x05, 0x28, 0x1f, 0x01,
	0x21, 0x20, 0x2e, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x02,
	0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02,
	0x01, 0x00, 0x00, 0x00, 0x09, 0x00, 0x45, 0x00, 0x49, 0x00, 0x27, 0x00, 0x6d, 0x00, 0x20, 0x00,
	0x61, 0x00, 0x6e, 0x00, 0x67, 0x00, 0x72, 0x00, 0x79, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5c, 0x00, 0x00, 0x00,
	0x16, 0x01, 0x01, 0x00, 0xff, 0xff, 0xff, 0xff, 0x0b, 0x18, 0x1b, 0x01, 0x0b, 0x28, 0x1b, 0x00,
	0x33, 0x20, 0x2a, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02,
	0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0x01, 0x00, 0x00, 0x00, 0x09, 0x00, 0x45, 0x00,
	0x48, 0x00, 0x65, 0x00, 0x6c, 0x00, 0x70, 0x00, 0x20, 0x00, 0x6d, 0x00, 0x65, 0x00, 0x21, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0xec, 0x00, 0x00, 0x00, 0x5e, 0x06, 0x38, 0x01, 0xff, 0xff, 0xff, 0xff,
	0x02, 0x17, 0x1b, 0x01, 0x02, 0x2a, 0x1b, 0x00, 0x31, 0x20, 0x2c, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x02,
	0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02, 0xff, 0x00, 0x00, 0x02,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
	0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
}
