package finder

import "os/exec"

// build for darwin
var openCmd = exec.Command("open", ".")
