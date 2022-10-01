package utils

import "os/exec"

func Copy(src string, dest string) {
	cmd := exec.Command("cp", "-R", src, dest)

	cmd.Run()
}
