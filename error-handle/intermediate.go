package main

import "os/exec"

// IntermediateErr ...
type IntermediateErr struct {
	error
}

// RunJob a simulate work
func RunJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := isGloballyExec(jobBinPath)
	if err != nil {
		return err
	} else if isExecutable == false {
		return wrapError(nil, "job binary is not executable")
	}
	return exec.Command(jobBinPath, "--id="+id).Run()
}
