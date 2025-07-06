package models

import (
	"fmt"
	"time"
)

// Birthday represents a person's birthday
type Birthday struct {
	Name    string
	Day     int
	Month   int
	Year    int
	HasYear bool
}

// NewBirthday creates a new Birthday instance
func NewBirthday(name string, day, month, year int, hasYear bool) *Birthday {
	return &Birthday{
		Name:    name,
		Day:     day,
		Month:   month,
		Year:    year,
		HasYear: hasYear,
	}
}

// ToString returns a string representation of the birthday
func (b *Birthday) ToString() string {
	if b.HasYear {
		return fmt.Sprintf("%02d.%02d.%04d %s", b.Day, b.Month, b.Year, b.Name)
	}
	return fmt.Sprintf("%02d.%02d. %s", b.Day, b.Month, b.Name)
}

// ToDate returns a time.Time representation of the birthday
func (b *Birthday) ToDate() time.Time {
	year := b.Year
	if !b.HasYear {
		year = time.Now().Year()
	}
	return time.Date(year, time.Month(b.Month), b.Day, 0, 0, 0, 0, time.UTC)
}

// GetUID returns a unique identifier for the birthday
func (b *Birthday) GetUID() string {
	cleanName := b.Name
	// Replace spaces with underscores for the UID
	for i := 0; i < len(cleanName); i++ {
		if cleanName[i] == ' ' {
			cleanName = cleanName[:i] + "_" + cleanName[i+1:]
		}
	}
	return fmt.Sprintf("birthday-%02d%02d-%s@birthday-md", b.Day, b.Month, cleanName)
}