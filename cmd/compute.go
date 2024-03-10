/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// computeCmd represents the compute command
var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "computes impact of manifest definition and its dependencies.",
	Long:  `This command takes a superset manifest yaml defined by IF specification. It computes the impact of the infra and its dependencies. It outputs the computed infra and its dependencies in a YAML file.`,
	Run: func(cmd *cobra.Command, args []string) {
		main()

	},
}
var manifest *string

func init() {
	rootCmd.AddCommand(computeCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	manifest = computeCmd.PersistentFlags().String("manifest", "manifest.yaml", "Path for manifest YAML file. The file defines your infra and its dependencies.")
	computeCmd.MarkPersistentFlagRequired("manifest")
	computeCmd.PersistentFlags().String("output", "", "Path for the output file. It will be a YAML file containing the computed infra and its dependencies.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//computeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const url = "https://charm.sh/"

type model struct {
	status int
	err    error
}

func checkServer() tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}
	defer res.Body.Close() // nolint:errcheck

	return statusMsg(res.StatusCode)
}

type statusMsg int

type errMsg struct{ err error }

// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e errMsg) Error() string { return e.err.Error() }

func (m model) Init() tea.Cmd {
	return checkServer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		m.status = int(msg)
		return m, tea.Quit

	case errMsg:
		m.err = msg
		return m, tea.Quit

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err)
	}

	s := fmt.Sprintf("Checking %s ... ", url)
	if m.status > 0 {
		s = fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
	}
	return "\n" + s + "\n\n"
}

func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
