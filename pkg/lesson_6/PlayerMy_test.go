package lesson_6

import (
	"testing"
	"unsafe"
)

func TestPlayerMy_Size(t *testing.T) {
	var p PlayerMy
	size := unsafe.Sizeof(p)

	if size != 64 {
		t.Errorf("PlayerMy size = %d bytes, want exactly 64 bytes", size)
	}

	t.Logf("✓ PlayerMy struct is exactly %d bytes", size)
}

func TestEncodeDecodeBlock1(t *testing.T) {
	// Test encoding and decoding block1
	encoded := EncodeBlock1()

	// Decoding - prints values to stdout
	DecodeBlock1(encoded)

	// Verify the encoded value has expected bits set
	// Expected: mana=850, health=950, respect=7, power=9, experience=5
	// We can manually verify the bit pattern
	expectedMana := uint32(850 & 0x3FF)
	expectedHealth := uint32(950&0x3FF) << 10
	expectedRespect := uint32(7&0xF) << 20
	expectedPower := uint32(9&0xF) << 24
	expectedExperience := uint32(5&0xF) << 28

	expected := expectedMana | expectedHealth | expectedRespect | expectedPower | expectedExperience

	if encoded != expected {
		t.Errorf("Encoded = %032b, want %032b", encoded, expected)
	}
}

func TestEncodeDecodeBlock2(t *testing.T) {
	// Test encoding and decoding block2
	encoded := EncodeBlock2()

	level, hasHouse, hasWeapon, hasFamily, playerType := DecodeBlock2(encoded)

	if level != 10 {
		t.Errorf("Level = %d, want 10", level)
	}
	if !hasHouse {
		t.Error("HasHouse = false, want true")
	}
	if !hasWeapon {
		t.Error("HasWeapon = false, want true")
	}
	if !hasFamily {
		t.Error("HasFamily = false, want true")
	}
	if playerType != 2 {
		t.Errorf("PlayerType = %d, want 2", playerType)
	}
}
