package environment

import (
	"fmt"
	"os/exec"
	"strings"
	"unisun/api/unisun-authen-listener/src/config/environment/models"

	"github.com/narawichsaphimarn/goyamlenv"
)

var ENV models.Configs

func LoadEnvironment() error {
	basepath, err := getBasePath()
	if err != nil {
		return err
	}
	basepath = strings.Join([]string{basepath, "/", "resources"}, "")
	err = goyamlenv.NewWithBasePath(basepath)
	if err != nil {
		return err
	}
	err = goyamlenv.LoadConfig(&ENV)
	if err != nil {
		return err
	}
	return nil
}

func getBasePath() (string, error) {
	prg := "pwd"
	cmd := exec.Command(prg)
	stdout, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%s{%v}", "Error step get base path application.Error msg : ", err)
	}
	return strings.Replace(string(stdout), "\n", "", 1), nil
}
