package lesson_6

import (
	"fmt"
)

type PlayerMy struct {
	// Ordering fields largest to smallest to minimize padding
	X        int32    // 4 bytes (offset 0)
	Y        int32    // 4 bytes (offset 4)
	Z        int32    // 4 bytes (offset 8)
	Gold     int32    // 4 bytes (offset 12)
	block1   uint32   // 4 bytes (offset 16)
	block2   uint16   // 2 bytes (offset 20)
	Username [42]byte // 42 bytes (offset 22)
}

func EncodeBlock1() uint32 {
	var encoded uint32 = 0

	mana := 850
	// Mana: bits 0-9 (10 bits)
	encoded = encoded | uint32(mana&0x3FF) // mask to 10 bits
	fmt.Printf("After mana(%d):       %032b\n", mana, encoded)

	health := 950
	// Health: bits 10-19 (10 bits)
	encoded = encoded | (uint32(health&0x3FF) << 10)
	fmt.Printf("After health(%d):     %032b\n", health, encoded)

	respect := 7
	// Respect: bits 20-23 (4 bits)
	encoded = encoded | (uint32(respect&0xF) << 20)
	fmt.Printf("After respect(%d):    %032b\n", respect, encoded)

	power := 9
	// Power: bits 24-27 (4 bits)
	encoded = encoded | (uint32(power&0xF) << 24)
	fmt.Printf("After power(%d):      %032b\n", power, encoded)

	experience := 5
	// Experience: bits 28-31 (4 bits)
	encoded = encoded | (uint32(experience&0xF) << 28)
	fmt.Printf("After experience(%d): %032b\n", experience, encoded)

	return encoded
}

func DecodeBlock1(bits uint32) {
	// Mana: bits 0-9
	manaMask := uint32(0x3FF)
	mana := uint16(bits & manaMask)
	fmt.Printf("Mana = %d (mask: %010b)\n", mana, manaMask)

	// Health: bits 10-19
	healthMask := uint32(0x3FF << 10)
	health := uint16((bits & healthMask) >> 10)
	fmt.Printf("Health = %d\n", health)

	// Respect: bits 20-23
	respectMask := uint32(0xF << 20)
	respect := byte((bits & respectMask) >> 20)
	fmt.Printf("Respect = %d\n", respect)

	// Power: bits 24-27
	powerMask := uint32(0xF << 24)
	power := byte((bits & powerMask) >> 24)
	fmt.Printf("Power = %d\n", power)

	// Experience: bits 28-31
	experienceMask := uint32(0xF << 28)
	experience := byte((bits & experienceMask) >> 28)
	fmt.Printf("Experience = %d\n", experience)
}

func EncodeBlock2() uint16 {
	var encoded uint16 = 0

	level := 10
	encoded = encoded | uint16(level)
	fmt.Printf("encoded = %016b\n", encoded)

	hasHouse := 1
	encoded = encoded | (uint16(hasHouse) << 4)
	fmt.Printf("encoded = %016b\n", encoded)

	hasWeapon := 1
	encoded = encoded | (uint16(hasWeapon) << 5)
	fmt.Printf("encoded = %016b\n", encoded)

	hasFamily := 1
	encoded = encoded | (uint16(hasFamily) << 6)
	fmt.Printf("encoded = %016b\n", encoded)

	playerType := 2 // 1, 2, 3
	encoded = encoded | (uint16(playerType) << 7)
	fmt.Printf("encoded = %016b\n", encoded)

	return encoded
}

func DecodeBlock2(bits uint16) (level byte, hasHouse, hasWeapon, hasFamily bool, playerType byte) {
	levelMask := 15
	level = byte(bits & uint16(levelMask))
	fmt.Println("level = ", level)

	hasHouseMask := 1 << 4
	hasHouse = ((bits & uint16(hasHouseMask)) >> 4) == 1
	fmt.Println("hasHouse = ", hasHouse)

	hasWeaponMask := 1 << 5
	hasWeapon = ((bits & uint16(hasWeaponMask)) >> 5) == 1
	fmt.Println("hasWeapon = ", hasWeapon)

	hasFamilyMask := 1 << 6
	hasFamily = ((bits & uint16(hasFamilyMask)) >> 6) == 1
	fmt.Println("hasFamily = ", hasFamily)

	playerTypeMask := 3 << 7
	playerType = byte((bits & uint16(playerTypeMask)) >> 7)
	fmt.Println("playerType = ", playerType)

	return
}
