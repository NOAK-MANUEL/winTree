package globarvar

import "embed"

var JustFile bool
var JustDir bool

var ShellScripts embed.FS
var HasScript string = "# wintree shell integration"
var ON_CD_COMMAND bool
