package github

import (
	"github.com/pterm/pterm"
	"os/exec"
	"strconv"
	"strings"
)

func ListOrganizations(limit int) error {
	cmd := exec.Command("gh", "org", "list", "--limit", strconv.Itoa(limit))
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	lines := strings.Split(string(output), "\n")
	tableData := pterm.TableData{{"Name"}}
	for _, line := range lines {
		if line != "" {
			tableData = append(tableData, []string{line})
		}
	}

	pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
	return nil
}
