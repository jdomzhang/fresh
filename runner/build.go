package runner

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func build() (string, bool) {
	buildLog("Building...")

	cmd := exec.Command("go", "build", "-o", buildPath(), mainPath())

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fatal(err)
	// }

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	// io.Copy(os.Stdout, stdout)
	cmd.Stdout = os.Stdout
	errBuf, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()
	if err != nil {
		return string(errBuf), false
	}

	return "", true
}
