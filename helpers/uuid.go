package helpers

import(
	"fmt"
    "os/exec"
    "strings"
)

func UuidGenerate() (string, error){
	out, err := exec.Command("uuidgen").Output()
	uuid := strings.Join(strings.Fields(fmt.Sprintf("%s", out)), "")
	return uuid, err
}