package dns

import (
	"fmt"
	"os/exec"
	"strings"
)

func (m *model) ResetDns() error {

	cmdRestore := exec.Command("sudo", "mv", "/etc/resolv.conf.backup", "/etc/resolv.conf")
	if err := cmdRestore.Run(); err != nil {
		return fmt.Errorf("failed to restore /etc/resolv.conf: %v", err)
	}

	fmt.Println("DNS settings restored from backup.")
	return nil

	// if m.osName == "windows" {

	// } else if m.osName == "linux" {

	// return setDNSLinux([]string{"8.8.8.8", "8.8.4.4"})

	// } else if m.osName == "darwin" {
	// fmt.Printf("This is an unknown operating system: %s\n", m.osName)
	// }
}

// TODO : implement for other os

func (m *model) SetDns(dnsName string) error {

	// if m.osName == "windows" {

	// } else if m.osName == "linux" {

	return setDNSLinux(m.dnsList[dns(dnsName)])

	// } else if m.osName == "darwin" {
	// fmt.Printf("This is an unknown operating system: %s\n", m.osName)
	// }

	// return nil
}

func setDNSLinux(dnsServers []string) error {

	cmdBackup := exec.Command("sudo", "cp", "/etc/resolv.conf", "/etc/resolv.conf.backup")
	if err := cmdBackup.Run(); err != nil {
		return fmt.Errorf("failed to backup /etc/resolv.conf: %v", err)
	}

	resolvConfContent := "nameserver " + strings.Join(dnsServers, "\nnameserver ") + "\n"

	cmd := exec.Command("sudo", "tee", "/etc/resolv.conf")
	cmd.Stdin = strings.NewReader(resolvConfContent)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to set DNS on Linux: %s\n%s", err, output)
	}

	fmt.Printf("DNS successfully set on Linux (may be overwritten):\n%s\n", output)
	return nil
}

// func setDNSWindows(adapterName string, dnsServers []string) error {

// 	primaryDNS := ""
// 	secondaryDNS := ""
// 	if len(dnsServers) > 0 {
// 		primaryDNS = dnsServers[0]
// 	}
// 	if len(dnsServers) > 1 {
// 		secondaryDNS = dnsServers[1]
// 	}

// 	// Command to set primary DNS
// 	cmdPrimary := exec.Command("netsh", "interface", "ip", "set", "dns", adapterName, "static", primaryDNS)
// 	outputPrimary, errPrimary := cmdPrimary.CombinedOutput()
// 	if errPrimary != nil {
// 		return fmt.Errorf("failed to set primary DNS on Windows: %s\n%s", errPrimary, outputPrimary)
// 	}
// 	fmt.Printf("Primary DNS set for %s:\n%s\n", adapterName, outputPrimary)

// 	// Command to set secondary DNS (if provided)
// 	if secondaryDNS != "" {
// 		cmdSecondary := exec.Command("netsh", "interface", "ip", "add", "dns", adapterName, secondaryDNS, "index=2")
// 		outputSecondary, errSecondary := cmdSecondary.CombinedOutput()
// 		if errSecondary != nil {
// 			// This might fail if the adapter already has a secondary DNS,
// 			// or if index=2 is not allowed without first deleting.
// 			// You might need more complex logic here to handle existing configurations.
// 			return fmt.Errorf("failed to add secondary DNS on Windows: %s\n%s", errSecondary, outputSecondary)
// 		}
// 		fmt.Printf("Secondary DNS added for %s:\n%s\n", adapterName, outputSecondary)
// 	}

// 	return nil
// }

// func setDNSDarwin() error {
// 	return nil
// }
