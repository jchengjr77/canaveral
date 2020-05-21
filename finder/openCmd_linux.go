package finder

import "os/exec"

// build for linux
var openCmd = exec.Command("xdg-open .", ".")
