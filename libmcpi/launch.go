package libmcpi

/*
libmcpi - the mcpi revival parts of the launcher, modularized
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

const (
	ENV_FEATURES = "MCPI_FEATURE_FLAGS"
	ENV_RENDDIST = "MCPI_RENDER_DISTANCE"
	ENV_USERNAME = "MCPI_USERNAME"
)

type LaunchProfile struct {
	FeatureFlags []string
	Username     string
	RendDistance string
	//information our implementation needs
	ExectuableName string
}

func (l *LaunchProfile) Launch() {
	//if we're not on linux, abort and dummy out because *mcpi doesn't run on windows*
	if runtime.GOOS != "linux" {
		fmt.Println("[libmcpi] [warn] launch() was called on a LaunchProfile. This has been dummied out due to the platform not being Linux")
		return
	}
	//otherwise carry on
	fmt.Printf("[libmcpi] [info] launch() has been called.\nFeature Flags:%v\nUseranme:%v\nRender Distance:%v\n", l.FeatureFlags, l.Username, l.RendDistance)

	//generate the string for ff env var from our list
	featuresEnvVal := ""
	for i, feature := range l.FeatureFlags {
		featuresEnvVal += feature
		if i > len(l.FeatureFlags)-2 {
			featuresEnvVal += "|"
		}
	}
	cmd := exec.Command(l.ExectuableName)
	//set the environment variables, and do other setup
	out, _ := cmd.StdoutPipe()
	err, _ := cmd.StderrPipe()
	go execLog(out)
	go execLog(err)
	cmd.Env = append(os.Environ(), ENV_FEATURES+"="+featuresEnvVal)
	cmd.Env = append(cmd.Env, ENV_RENDDIST+"="+l.RendDistance)
	cmd.Env = append(cmd.Env, ENV_USERNAME+"="+l.Username)
	//now, lets get this party on the road
	fmt.Println("[libmcpi] [info] starting MCPI process and detaching")
	//run async in a frankly cursed way
	//TODO: there *must* be a better way to do this...
	go cmdRunHolder(cmd)
}

func cmdRunHolder(cmd *exec.Cmd) {
	cmd.Run()
	for {
	}
}

func execLog(inPipe io.ReadCloser) {
	scanner := bufio.NewScanner(inPipe)
	for scanner.Scan() {
		fmt.Printf("[mcpi] %v\n", scanner.Text())
	}
}
