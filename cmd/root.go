/*
Copyright © 2020-2021 Hannes Hayashi

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

// Package cmd contains the commands available to the end user
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmauth"
	"github.com/hanneshayashi/gsm/gsmcalendar"
	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmcibeta"
	"github.com/hanneshayashi/gsm/gsmconfig"
	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmgmailpostmaster"
	"github.com/hanneshayashi/gsm/gsmgroupssettings"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmlicensing"
	"github.com/hanneshayashi/gsm/gsmpeople"
	"github.com/hanneshayashi/gsm/gsmreports"
	"github.com/hanneshayashi/gsm/gsmsheets"

	// crescengo "github.com/hanneshayashi/crescengo"
	// crescengo "github.com/hanneshayashi/gsm/crescengo"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile        string
	dwdSubject     string
	logFile        string
	home           string
	compressOutput bool
	streamOutput   bool
	batchFlags     map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
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
		"batchThreads": {
			AvailableFor: []string{"batch"},
			Type:         "int",
			Description:  "Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16)",
		},
	}
	recursiveFileFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
		"folderId": {
			AvailableFor: []string{"recursive"},
			Type:         "string",
			Description:  `File id of the folder.`,
			Required:     []string{"recursive"},
		},
		"batchThreads": {
			AvailableFor: []string{"recursive"},
			Type:         "int",
			Description:  "Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16)",
		},
		"excludeFolders": {
			AvailableFor: []string{"recursive"},
			Type:         "stringSlice",
			Description: `Ids of folders to exclude.
Note that due to the way permissions are automatically inherited in Drive, this may not have the desired result for permission commands!`,
		},
	}
	recursiveUserFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
		"orgUnit": {
			AvailableFor: []string{"recursive"},
			Type:         "stringSlice",
			Description:  `Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children!`,
		},
		"groupEmail": {
			AvailableFor: []string{"recursive"},
			Type:         "stringSlice",
			Description:  `An email address of a group. Can be used multiple times. Note that a group will include recursive memberships!`,
		},
		"batchThreads": {
			AvailableFor: []string{"recursive"},
			Type:         "int",
			Description:  "Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16)",
		},
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gsm",
	Short: "GoSpace Manager - Manage Google Workspace resources using a developer-friendly CLI written in Go",
	Long: `GSM is free software licensed under the GPLv3 (https://gsm.hayashi-ke.online/license).
Copyright © 2020-2021 Hannes Hayashi.
For documentation see https://gsm.hayashi-ke.online.`,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
	Version: "v0.5.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// gsmhelpers.CreateDocs(rootCmd)
	// crescengo.CreateCrescendoModuleDefs(rootCmd, "../gsm-powershell/json/", "--compressOutput", "--streamOutput")
}

func init() {
	cobra.OnInitialize(setHomeDir, initConfig, initLog, auth)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/gsm/.gsm.yaml)")
	rootCmd.PersistentFlags().StringVar(&dwdSubject, "dwdSubject", "", "Specify a subject used for DWD impersonation (overrides value in config file)")
	rootCmd.PersistentFlags().BoolVar(&compressOutput, "compressOutput", false, `By default, GSM outputs "pretty" (indented) objects. By setting this flag, GSM's output will be compressed. This may or may not improve performance in scripts.`)
	rootCmd.PersistentFlags().BoolVar(&streamOutput, "streamOutput", false, `Setting this flag will cause GSM to output slice values to stdout one by one, instead of one large object`)
	rootCmd.PersistentFlags().IntVar(&gsmhelpers.StandardDelay, "delay", 0, "This delay (plus a random jitter between 0 and 50) will be applied after every command to avoid reaching quota and rate limits. Set to 0 to disable.")
	rootCmd.PersistentFlags().StringVar(&logFile, "log", "", "Set the path of the log file. Default is either ~/gsm.log or defined in your config file")
	rootCmd.PersistentFlags().IntSliceVar(&gsmhelpers.RetryOn, "retryOn", nil, "Specify the HTTP error code(s) that GSM should retry on. Note that GSM will always retry on HTTP 403 errors that indicate a quota / rate limit error")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var err error
	gsmconfig.CfgDir = fmt.Sprintf("%s/.config/gsm", home)
	if _, err = os.Stat(gsmconfig.CfgDir); os.IsNotExist(err) {
		err = os.MkdirAll(gsmconfig.CfgDir, 0777)
		if err != nil {
			log.Fatalf("Config dir %s could not be found and could not be created: %v", gsmconfig.CfgDir, err)
		}
	}
	// Search config in home directory with name ".gsm" (without extension).
	viper.AddConfigPath(gsmconfig.CfgDir)
	if cfgFile == "" {
		cfgFile = ".gsm"
	}
	viper.SetConfigName(cfgFile)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err = viper.ReadInConfig(); err != nil && gsmhelpers.IsCommandOrChild(configsCmd, logCmd) {
		fmt.Println(`Error loading config file. Please run "gsm configs new" to create a new config and load it with "gsm configs load --name"`)
	}
	if rootCmd.Flags().Changed("delay") {
		gsmhelpers.StandardDelay, err = rootCmd.Flags().GetInt("delay")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		gsmhelpers.StandardDelay = viper.GetInt("standardDelay")
	}
	if streamOutput {
		compressOutput = true
	}
}

func setHomeDir() {
	var err error
	home, err = homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func auth() {
	if gsmhelpers.IsCommandOrChild(configsCmd, logCmd) {
		return
	}
	mode := viper.GetString("mode")
	var subject string
	var credentials []byte
	var err error
	var client *http.Client
	if mode == "dwd" || mode == "adc" {
		if dwdSubject == "" {
			subject = viper.GetString("subject")
		} else {
			subject = dwdSubject
		}
	}
	if mode == "dwd" || mode == "user" {
		credentials, err = ioutil.ReadFile(viper.GetString("credentialsFile"))
		if err != nil {
			fmt.Printf("Error reading service account credentials file: %v", err)
		}
	}
	switch mode {
	case "dwd":
		client = gsmauth.GetClient(subject, credentials, viper.GetStringSlice("scopes")...)
	case "user":
		client = gsmauth.GetClientUser(credentials, fmt.Sprintf("%s_token.json", viper.GetString("name")), viper.GetStringSlice("scopes")...)
	case "adc":
		client = gsmauth.GetClientADC(subject, viper.GetString("serviceAccount"), viper.GetStringSlice("scopes")...)
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
	gsmreports.SetClient(client)
	gsmgmailpostmaster.SetClient(client)
	gsmcibeta.SetClient(client)
}

func initLog() {
	if logFile == "" {
		logFile = viper.GetString("logFile")
		if logFile == "" {
			logFile = fmt.Sprintf("%s/gsm.log", home)
		}
	}
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(file)
}
