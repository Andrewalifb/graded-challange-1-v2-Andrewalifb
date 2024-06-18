package utils // or cronjobs, schedulers, etc.

import (
	"context"

	"log"
	"time"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/dto"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/entity"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/repository"
	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson"
)

// Scheduler For Check Each transaction ID That Still Has Status Pending If After 24 Hours Has Not Completed Payment Yet. That Transaction Will Be Deleted
func StartCronJob(repo repository.TransactionRepository) {
	// Run on WIB Indonesia Time
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalf("%s: %v", dto.MESSAGE_FAILED_CRON_JOB, err)
	}
	s := gocron.NewScheduler(location)

	job := func() {
		ctx := context.Background()
		now := time.Now()
		oneDayAgo := now.Add(-24 * time.Hour)
		// Set Filter for each transaction has been more than 24 hours
		filter := bson.M{
			"createdAt": bson.M{
				"$lt": oneDayAgo,
			},
			// and status still pendong or not completed
			"status": "pending",
		}
		// Get All Transaction Data That Has Been Filtered
		cursor, err := repo.FindTransactions(ctx, filter)
		if err != nil {
			log.Printf("%s: %v", dto.MESSAGE_FAILED_CRON_JOB, err)
			return
		}

		for cursor.Next(ctx) {
			var transaction entity.Transaction
			if err := cursor.Decode(&transaction); err != nil {
				log.Printf("%s: %v", dto.MESSAGE_FAILED_CRON_JOB, err)
				continue
			}
			// Delete All Transaction That More Than 24 Hours and Status Still Pending
			err = repo.DeleteTransactionByID(ctx, transaction.ID)
			if err != nil {
				log.Printf("%s: %v", dto.MESSAGE_FAILED_CRON_JOB, err)
			}
		}
	}
	// Run Everyday on 00:00
	s.Every(1).Day().At("00:00").Do(job)

	s.StartAsync()
}
