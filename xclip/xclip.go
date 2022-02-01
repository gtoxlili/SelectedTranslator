package xclip

func GetSelection() string {
	return RunShell(xclipPath, "-selection primary", "-out")
}
