package rpm

import (
	"bytes"
	"errors"
        "os/exec"
	"strings"
)

func GetSourceRPMName(rpm string) (srcname string, err error) {

        var out bytes.Buffer
        var stderr bytes.Buffer

        err = nil

        cmd := exec.Command("rpm", "-qp", "--qf", "%{SOURCERPM}", rpm)
        cmd.Stdout = &out
        cmd.Stderr = &stderr

        // log.Printf("Executing %s: %v", cmd.Path, cmd.Args)

        if err = cmd.Run(); err != nil {
                return stderr.String(), err
        }

        return out.String(), err
}

func GetRPMFilelist(rpm string) (list []string, err error) {

        var out bytes.Buffer
        var stderr bytes.Buffer

        err = nil

        cmd := exec.Command("rpm", "-qlp", rpm)
        cmd.Stdout = &out
        cmd.Stderr = &stderr

        // log.Printf("Executing %s: %v", cmd.Path, cmd.Args)

        if err = cmd.Run(); err != nil {
                return nil, err
        }

        return strings.Split(out.String(), "\n"), err
}

func SplitRPMname(rpm string) (name string, version string, release string, arch string, err error) {

	if !strings.HasSuffix(rpm, ".rpm") {
		return "", "", "", "", errors.New("No RPM name")
	}

	str := strings.TrimSuffix(rpm, ".rpm")

	// get architecture
	lastInd := strings.LastIndex(str, ".")
	arch = str[lastInd+1:]
	str = str[:lastInd]

	// get release
	lastInd = strings.LastIndex(str, "-")
	release = str[lastInd+1:]
	str = str[:lastInd]

	// get version and package name
	lastInd = strings.LastIndex(str, "-")
	version = str[lastInd+1:]
	name = str[:lastInd]

	return name, version, release, arch, nil
}
