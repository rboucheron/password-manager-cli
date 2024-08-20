/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pwm/colors"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		help()
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printBanner() {
	banner := `
          _____                    _____                    _____          
         /\    \                  /\    \                  /\    \         
        /::\    \                /::\____\                /::\____\        
       /::::\    \              /:::/    /               /::::|   |        
      /::::::\    \            /:::/   _/___            /:::::|   |        
     /:::/\:::\    \          /:::/   /\    \          /::::::|   |        
    /:::/__\:::\    \        /:::/   /::\____\        /:::/|::|   |        
   /::::\   \:::\    \      /:::/   /:::/    /       /:::/ |::|   |        
  /::::::\   \:::\    \    /:::/   /:::/   _/___    /:::/  |::|___|______  
 /:::/\:::\   \:::\____\  /:::/___/:::/   /\    \  /:::/   |::::::::\    \ 
/:::/  \:::\   \:::|    ||:::|   /:::/   /::\____\/:::/    |:::::::::\____\
\::/    \:::\  /:::|____||:::|__/:::/   /:::/    /\::/    / ~~~~~/:::/    /
 \/_____/\:::\/:::/    /  \:::\/:::/   /:::/    /  \/____/      /:::/    / 
          \::::::/    /    \::::::/   /:::/    /               /:::/    /  
           \::::/    /      \::::/___/:::/    /               /:::/    /   
            \::/____/        \:::\__/:::/    /               /:::/    /    
             ~~               \::::::::/    /               /:::/    /     
                               \::::::/    /               /:::/    /      
                                \::::/    /               /:::/    /       
                                 \::/____/                \::/    /        
                                  ~~                       \/____/         
                                                                         
`
	fmt.Println(colors.InfoColor(banner))
}

func help() {

	help := `Available Commands:
  add         Add a new password                 add [service] [username] [password]
  delete      remove a password save previous    delete [service]
  gen         generate password                  gen [size]
  get         Recover a password                 get [service]
  help        Help about any command
  list        list all passwords saved

Flags:
  -h, --help     help for password-manager
  -t, --toggle   Help message for toggle

Use "password-manager [command] --help" for more information about a command."`
	printBanner()
	fmt.Print(colors.WarningColor(help))
}
