package preinstall

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// check si on est root, faut sudo
func root() error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("erreur, vous devez etre root")
	}
	return nil
}

// check linux, ne marchera pas sous windows pcq on utilise /ect et bah crowdsec sous windows mdr
func systeme() error {
	if runtime.GOOS != "linux" {
		return fmt.Errorf("erreur, vous devez etre sous Linux")
	}
	return nil
}

// check si crowdsec est deja la
func here() error {
	_, err := exec.LookPath("crowdsec")
	if err == nil {
		return fmt.Errorf("crowdsec est deja installé")
	}
	return nil
}

// Pass effectue toutes les vérifications préalables nécessaires, dans le bon ordre surtout (check le main.go)
func Pass() error {
	if err := systeme(); err != nil {
		return err
	}
	if err := root(); err != nil {
		return err
	}
	if err := here(); err != nil {
		return err
	}
	return nil
}

