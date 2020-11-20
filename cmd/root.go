/*
Package cmd contains the commands available to the end user
Copyright © 2020 Hannes Hayashi

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"gsm/gsmadmin"
	"gsm/gsmauth"
	"gsm/gsmcalendar"
	"gsm/gsmci"
	"gsm/gsmconfig"
	"gsm/gsmdrive"
	"gsm/gsmgmail"
	"gsm/gsmgroupssettings"
	"gsm/gsmhelpers"
	"gsm/gsmlicensing"
	"gsm/gsmpeople"
	"gsm/gsmsheets"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgDir     string
	cfgFile    string
	dwdSubject string
	batchFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
		"path": {
			AvailableFor: []string{"batch"},
			Type:         "string",
			Description:  "Path of the import file (CSV)",
			Required:     []string{"batch"},
		},
		"delimiter": {
			AvailableFor: []string{"batch"},
			Type:         "string",
			Description:  "Delimiter to use for CSV columns. Must be exactly one character.",
		},
		"skipHeader": {
			AvailableFor: []string{"batch"},
			Type:         "bool",
			Description:  "Whether to skip the first row (header)",
		},
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gsm",
	Short: "GoSpace Manager - Manage Google Workspace resources using a simple CLI written in Golang",
	Long: `GSM is free software licenced under the GPLv3 (https://gsm.hayashi-ke.online/license).
Copyright © 2020 Hannes Hayashi.
For documentation see https://gsm.hayashi-ke.online.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// gsmhelpers.CreateDocs(rootCmd)
}

func init() {
	cobra.OnInitialize(initConfig, initLog, auth)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/gsm/.gsm.yaml)")
	rootCmd.PersistentFlags().StringVar(&dwdSubject, "dwdSubject", "", "Specify a subject used for DWD impersonation (overrides value in config file)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		cfgFile = gsmconfig.GetConfigPath(cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		gsmconfig.CfgDir = fmt.Sprintf("%s/.config/gsm", home)
		if _, err := os.Stat(gsmconfig.CfgDir); os.IsNotExist(err) {
			err = os.MkdirAll(gsmconfig.CfgDir, 0777)
			if err != nil {
				log.Fatalf("Config dir %s could not be found and could not be created: %v", gsmconfig.CfgDir, err)
			}
		}
		// Search config in home directory with name ".gsm" (without extension).
		viper.AddConfigPath(gsmconfig.CfgDir)
		viper.SetConfigName(".gsm")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// log.Printf("Error reading config file: %v", err)
		// _, err = gsmconfig.CreateConfig(&gsmconfig.GSMConfig{Name: ".gsm"})
		// if err != nil {
		// 	log.Fatalf("Error creating default empty config file: %v", err)
		// }
		log.Println(`Error loading config file. Please run "gsm configs new" to create a new config and load it with "gsm configs load --name"`)
	}
}

func auth() {
	credentials, err := ioutil.ReadFile(viper.GetString("credentialsFile"))
	if err != nil {
		log.Printf("Error reading service account credentials file: %v", err)
	} else {
		var client *http.Client
		switch viper.GetString("mode") {
		case "dwd":
			var subject string
			if dwdSubject == "" {
				subject = viper.GetString("subject")
			} else {
				subject = dwdSubject
			}
			client = gsmauth.GetClient(subject, credentials, viper.GetStringSlice("scopes")...)
		case "user":
			client = gsmauth.GetClientUser(credentials, fmt.Sprintf("%s_token.json", viper.GetString("name")), viper.GetStringSlice("scopes")...)
		}
		gsmadmin.SetClient(client)
		gsmgmail.SetClient(client)
		gsmci.SetClient(client)
		gsmdrive.SetClient(client)
		gsmgroupssettings.SetClient(client)
		gsmcalendar.SetClient(client)
		gsmlicensing.SetClient(client)
		gsmpeople.SetClient(client)
		gsmsheets.SetClient(client)
	}
}

func initLog() {
	file, err := os.OpenFile("gsm.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}
