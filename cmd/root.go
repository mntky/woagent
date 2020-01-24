/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "fmt"
	"woagent/pkg"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"

)


var cfgFile string


func NewWoagent() *cobra.Command {
	var rootCmd = &cobra.Command{
	  Use:   "woagent",
	  Short: "woyendetsa cluster worker agent",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run()
		},
	}

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("listen", "l", "", "receive notifer address")
	rootCmd.PersistentFlags().StringP("join", "j", "", "join woyendetsa cluster")
	viper.BindPFlag("listen", rootCmd.PersistentFlags().Lookup("listen"))
	viper.BindPFlag("join", rootCmd.PersistentFlags().Lookup("join"))
	return rootCmd
}

func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  }
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
  viper.AddConfigPath("$HOME/.woagent/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("%s", err))
  }
}

func Run() error {
	fmt.Println("start woagent")

	//url := viper.GetString("listen")
	if masterurl := viper.GetString("join"); masterurl == "" {
		panic("require -j [master address]")
	}
	err := pkg.Join(viper.GetString("join"), viper.GetString("listen"))
	if err != nil {
		return err
	}
	err = startAgent(viper.GetString("listen"))
	if err != nil {
		return err
	}

	return nil
}

