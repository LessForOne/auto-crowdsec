package reboot

import (
	"fmt"
	"os"
	"os/exec"
)

func Reb() error {
	cmd := exec.Command("sudo", "systemctl", "restart", "crowdsec")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Erreur: ", err)
		return err
	}
	return nil
}
