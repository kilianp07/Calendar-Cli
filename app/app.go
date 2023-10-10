package app

import (
	"fmt"
	"os"

	"github.com/kilianp07/Calendar-Cli/calib"
)

func Start(month int, year int, format string, output string) {

	cal := calib.NewCalendar(month, year)

	// Format the calendar
	formatted, err := formatCal(format, cal)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write to file
	err = writeToFile(formatted, output, format)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func formatCal(format string, cal *calib.Calendar) (string, error) {
	switch format {
	case "json":
		return cal.ToJSON()
	case "html":
		return cal.ToHtml()
	default:
		return "", fmt.Errorf("unknown format %s", format)
	}
}

func writeToFile(content string, output string, format string) error {

	// Create file
	name := fmt.Sprintf("%s.%s", output, format)
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	// Write to file
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	// Close file
	return f.Close()

}
