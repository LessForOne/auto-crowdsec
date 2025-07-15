package bouncer

import (
	"fmt"
	"os"
	"os/exec"
)

func Bouncer() error {
	cmd := exec.Command("sudo", "apt", "install", "crowdsec-firewall-bouncer-iptables")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Erreur: ", err)
	}
	fmt.Println("Bouncer install ok")
	return nil
}
