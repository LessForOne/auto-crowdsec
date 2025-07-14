package enrolling

import (
	"fmt"
	"os"
	"os/exec"
)

func Enroll() error {
	var enrollkey string
	fmt.Println("Entrez votre enroll key: ")
	fmt.Scan(&enrollkey)
	fmt.Println("Votre clé est: ", enrollkey)

	cmd := exec.Command("sudo", "cscli", "console", "enroll", "-e", enrollkey)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // pour le mot de passe sudo si besoin

	if err := cmd.Run(); err != nil {
		fmt.Println("Erreur: ", err)
		return err
	}
	fmt.Println("Enroll ok")
	return nil
}
