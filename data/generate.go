package data

import (
	"fmt"
	"math/rand"
	"strings"
)

func Generate(dataType string) string {
	switch dataType {
	case TYPE_NAME:
		return generateName()
	case TYPE_DATE:
		return generateDate()
	case TYPE_ADDRESS:
		return generateAddress()
	case TYPE_PHONE:
		return generatePhone()
	}

	return ""
}

// generateName membuat nama dari kumpulan kata yang sudah ada
func generateName() string {
	nameLen := len(name)

	return name[rand.Intn(nameLen)]
}

// generateDate membuat tanggal acak
func generateDate() string {
	year := rand.Intn(100) + 1950
	month := rand.Intn(12) + 1
	day := rand.Intn(28) + 1

	return fmt.Sprintf("%02d-%02d-%d", day, month, year)
}

// generateAddress membuat alamat acak dari kumpulan kata yang sudah ada dan nomor acak
func generateAddress() string {
	streetLen := len(address[SUBTYPE_ADDRESS_STREET])
	cityLen := len(address[SUBTYPE_ADDRESS_CITY])

	return fmt.Sprintf("Jl. %s No. %d, %s", address[SUBTYPE_ADDRESS_STREET][rand.Intn(streetLen)], rand.Intn(100), address[SUBTYPE_ADDRESS_CITY][rand.Intn(cityLen)])
}

// generatePhone generates a random phone number with 12 digits limit
func generatePhone() string {
	prefixLen := 6 + rand.Intn(4) // 6-9 digits

	// prepare suffix
	var sb strings.Builder
	sb.WriteString("081")

	for i := 0; i < prefixLen; i++ {
		sb.WriteString(fmt.Sprintf("%d", rand.Intn(10)))
	}

	return sb.String()
}
