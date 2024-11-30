package helper

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func ConvertTimeToSeconds(time string) (int, error) {
	// Memisahkan waktu berdasarkan tanda ':'
	parts := strings.Split(time, ":")

	// Pastikan formatnya valid (harus ada 3 bagian: jam, menit, detik)
	if len(parts) != 3 {
		return 0, errors.New("invalid time format, must be in HH:MM:SS format")
	}

	// Mengonversi jam, menit, dan detik ke integer
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid hour value: %v", err)
	}
	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minute value: %v", err)
	}
	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, fmt.Errorf("invalid second value: %v", err)
	}

	// Menghitung total detik
	totalSeconds := hours*3600 + minutes*60 + seconds

	return totalSeconds, nil
}

func SearchData(db *gorm.DB, filters map[string]interface{}) *gorm.DB {
	// Mulai query dasar
	query := db

	// Menggunakan filter dinamis
	for key, value := range filters {
		// Menambahkan kondisi WHERE untuk setiap filter
		query = query.Or("UPPER("+key+") LIKE ?", "%"+strings.ToUpper(value.(string))+"%")
	}

	// Menjalankan query dan mendapatkan hasil
	return query
}
