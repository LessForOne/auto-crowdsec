package main

import (
	"auto-crowdsec/pkg/enrolling"
	install "auto-crowdsec/pkg/installation"
	preinstall "auto-crowdsec/pkg/preinstallation"
	"auto-crowdsec/pkg/reboot"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Verif des prérequis en cours..")

	if err := preinstall.Pass(); err != nil {
		fmt.Println("Erreur: ", err)
		os.Exit(1)
	}
	fmt.Println("Linux: ok")
	fmt.Println("Permission: ok")
	fmt.Println("Crowdsec n'est pas dejé installé: ok")
	fmt.Println("Tout est bon !")

	if err := install.Install(); err != nil {
		fmt.Println("Erreur: ", err)
		os.Exit(1)
	}

	if err := enrolling.Enroll(); err != nil {
		fmt.Println("Erreur: ", err)
		os.Exit(1)
	}

	if err := reboot.Reb(); err != nil {
		fmt.Println("Erreur: ", err)
		os.Exit(1)
	}
}
