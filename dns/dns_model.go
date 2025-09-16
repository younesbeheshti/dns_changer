package dns

import (
	"fmt"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	firstList state = iota
	secondList
)

type dns string

const (
	dnsShecan     dns = "shecan"
	dns403            = "403"
	dnsBegzar         = "begzar"
	dnsRadarGame      = "radar.game"
	dnsElectro        = "elctro"
	dnsShelter        = "shelter"
	dnsBeshcan        = "beshcan"
	dnsLevel3         = "level3"
	dnsCloudflare     = "cloudflare"
	dnsGoogle         = "google"
)

type model struct {
	firstItems  []string
	secondItems []string
	dnsList     map[dns][]string
	state       state
	cursor      int
	osName      string
}

func (m model) Init() tea.Cmd {
	return nil
}


func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.state == firstList {
				if m.cursor > 0 {
					m.cursor--
				}
			} else {
				if m.cursor > 0 {
					m.cursor--
				}
			}
		case "down", "j":
			if m.state == firstList {
				if m.cursor < len(m.firstItems)-1 {
					m.cursor++
				}
			} else {
				if m.cursor < len(m.secondItems)-1 {
					m.cursor++
				}
			}
		case "enter":
			if m.state == firstList && m.cursor == 0 {
				m.state = secondList
				m.cursor = 0
				fmt.Println("DNS List")
				return m, nil
			} else if m.state == firstList && m.cursor == 1 {
				
				if err := m.ResetDns(); err != nil {
					fmt.Printf("[Error] - fail to reset dns: %v", err)
				}

				fmt.Println("DNS reset successfully")
				return m, tea.Quit


			} else {
				selectedDNS := m.secondItems[m.cursor]
				fmt.Printf("Selected DNS: %s, \n", selectedDNS)
				err := m.SetDns(selectedDNS)
				if err != nil {
					fmt.Printf("[Error]: %v", err)
				}
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.state == firstList {
		s := "\nMain Menu\n\n"
		for i, item := range m.firstItems {
			cursor := "  "
			if i == m.cursor {
				cursor = "> "
			}
			s += fmt.Sprintf("%s%s\n", cursor, item)
		}
		s += "\nUse arrow keys to navigate, Enter to select, 'q' to quit\n"
		return s
	} else {
		s := "\nDNS List\n\n"
		for i, item := range m.secondItems {
			cursor := "  "
			if i == m.cursor {
				cursor = "> "
			}
			s += fmt.Sprintf("%s%s\n", cursor, item)
		}
		s += "\nUse arrow keys to navigate, Enter to select, 'q' to quit\n"
		return s
	}
}

func InitialModel() model {
	dnsList := map[dns][]string{
		dnsShecan:     {"178.22.122.100", "185.51.200.2"},
		dns403:        {"10.202.10.202", "10.202.10.102"},
		dnsBegzar:     {"185.55.225.25", "185.55.226.26"},
		dnsRadarGame:  {"10.202.10.10", "10.202.10.11"},
		dnsElectro:    {"78.157.42.100", "78.157.42.101"},
		dnsShelter:    {"94.103.125.157", "94.103.125.158"},
		dnsBeshcan:    {"181.41.194.177", "181.41.194.186"},
		dnsLevel3:     {"209.244.0.3", "209.244.0.4"},
		dnsCloudflare: {"1.1.1.1", "1.0.0.1"},
		dnsGoogle:     {"8.8.8.8", "8.8.4.4"},
	}

	return model{
		osName: runtime.GOOS,
		firstItems: []string{
			"active dns",
			"deactive dns",
		},
		secondItems: []string{
			"shecan",
			"403",
			"begzar",
			"radar.game",
			"elctro",
			"shelter",
			"beshcan",
			"level3",
			"cloudflare",
			"google",
		},
		dnsList: dnsList,
		state:   firstList,
		cursor:  0,
	}
}
