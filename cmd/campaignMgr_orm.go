/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"publicX/internal"
	"publicX/pkg/dao"
)

func CreateCampaign(path string, campaignID string, template string, schedualTime int64) {
	db := dao.GetDBConn()
	if db != nil {
		db.AutoMigrate(&dao.Campaign{})

		c := &dao.Campaign{
			CampaignID:   campaignID,
			Template:     template,
			SchedualTime: schedualTime,
		}
		db.Create(c)

		internal.AddRecipients(db, path, campaignID)
	}
}

// campaignMgrCmd represents the campaignMgr command
var campaignMgrOrmCmd = &cobra.Command{
	Use:   "campaignMgrOrm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		CreateCampaign("./data.csv", "c3", "welcome %s", time.Now().Unix()+30)
		fmt.Println("campaignMgrOrm called")
	},
}

func init() {
	rootCmd.AddCommand(campaignMgrOrmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// campaignMgrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// campaignMgrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
