package internal

import (
	"encoding/csv"
	"fmt"
	"os"

	"gorm.io/gorm"

	"publicX/pkg/dao"
)

func AddRecipients(db *gorm.DB, filePath, campaignID string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("err msg:%s", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	// reader.Comma = ';'

	db.AutoMigrate(&dao.Recipient{})

	for {
		list := []dao.Recipient{}
		for i := 0; i < 100; i++ {
			record, e := reader.Read()
			if e != nil {
				// fmt.Println(e)
				break
			}
			fmt.Println(record)

			r := dao.Recipient{
				Name:        record[0],
				PhoneNumber: record[1],
				CampaignID:  campaignID,
			}
			list = append(list, r)

			//db.CreateInBatches(list, len(list))
		}
		if len(list) > 0 {
			db.Create(list)
		}
	}

}
