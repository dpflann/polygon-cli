
package cmd

import (
	"os"
	"fmt"
	cobra "github.com/spf13/cobra"
	viper "github.com/spf13/viper"
)


var cfgFile string
var natsEndpoints string
var apiKey string


var RootCmd = &cobra.Command{
	Use:   "polygon-cli",
	Short: "polygon-cli utility",
	Long: "polygon-cli is a console util tool to access Polygon APIs\n",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Running Command..")
	},
}


func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}


func init(){
	cobra.OnInitialize( initConfig )
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.polygon-cli.yaml)")
	RootCmd.PersistentFlags().StringVar(&natsEndpoints, "natsEndpoints", "nats://nats1.polygon.io:30401,nats://nats2.polygon.io:30402,nats://nats3.polygon.io:30403", "NATS url(s), delimited by comma")
	RootCmd.PersistentFlags().StringVar(&apiKey, "apiKey", "", "Polygon API Key")
}


func initConfig(){
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigName(".polygon-cli")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

