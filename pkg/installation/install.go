package install

import (
	"fmt"
	"os"
	"os/exec"
)

func Install() error {
	fmt.Println("Installation de CrowdSec en cours...")

	//recup les répo
	cmd := exec.Command("sh", "-c", "curl -s https://install.crowdsec.net | sudo sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("erreur: %w", err)
	}

	//mise a jour des paquets
	cmd = exec.Command("sudo", "apt", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("erreur: %w", err)
	}

	//installation de crowdsec
	cmd = exec.Command("sudo", "apt", "insttall", "-y", "crowdsec")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("erreur: %w", err)
	}
	
	fmt.Println("OK")
	return nil
}

