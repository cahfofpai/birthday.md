package ics

import (
	"fmt"
	"os"
	"time"

	"github.com/cahfofpai/birthday.md/internal/models"
)

// Generator is responsible for generating ICS files
type Generator struct {
	outputPath string
	birthdays  []*models.Birthday
}

// NewGenerator creates a new Generator instance
func NewGenerator(outputPath string, birthdays []*models.Birthday) *Generator {
	return &Generator{
		outputPath: outputPath,
		birthdays:  birthdays,
	}
}

// Generate generates an ICS file from the birthdays
func (g *Generator) Generate() error {
	file, err := os.Create(g.outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	// Write ICS header
	header := "BEGIN:VCALENDAR\r\n" +
		"VERSION:2.0\r\n" +
		"PRODID:-//github.com/cahfofpai/birthday.md//NONSGML v1.0//EN\r\n" +
		"CALSCALE:GREGORIAN\r\n" +
		"METHOD:PUBLISH\r\n"

	_, err = file.WriteString(header)
	if err != nil {
		return fmt.Errorf("failed to write ICS header: %w", err)
	}

	// Write events for each birthday
	for _, birthday := range g.birthdays {
		event, err := g.generateEvent(birthday)
		if err != nil {
			return fmt.Errorf("failed to generate event for %s: %w", birthday.Name, err)
		}

		_, err = file.WriteString(event)
		if err != nil {
			return fmt.Errorf("failed to write event for %s: %w", birthday.Name, err)
		}
	}

	// Write ICS footer
	footer := "END:VCALENDAR\r\n"
	_, err = file.WriteString(footer)
	if err != nil {
		return fmt.Errorf("failed to write ICS footer: %w", err)
	}

	return nil
}

// generateEvent generates an ICS event for a birthday
func (g *Generator) generateEvent(birthday *models.Birthday) (string, error) {
	now := time.Now().UTC().Format("20060102T150405Z")
	date := birthday.ToDate()
	dateStr := date.Format("20060102")

	// Create the event
	event := "BEGIN:VEVENT\r\n" +
		fmt.Sprintf("UID:%s\r\n", birthday.GetUID()) +
		fmt.Sprintf("DTSTAMP:%s\r\n", now) +
		fmt.Sprintf("DTSTART;VALUE=DATE:%s\r\n", dateStr) +
		"TRANSP:TRANSPARENT\r\n" +
		fmt.Sprintf("SUMMARY:%s's Birthday\r\n", birthday.Name) +
		"RRULE:FREQ=YEARLY\r\n" +
		"CATEGORIES:Birthday\r\n" +
		"END:VEVENT\r\n"

	return event, nil
}