
package colors

import "github.com/fatih/color"

var (
    ErrorColor   = color.New(color.FgRed).SprintFunc()
    SuccessColor = color.New(color.FgGreen).SprintFunc()
    InfoColor    = color.New(color.FgBlue).SprintFunc()
    WarningColor = color.New(color.FgYellow).SprintFunc()
)
