package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Mailbox struct {
	ID              uint `gorm:"primaryKey"`
	Address         string
	EmailsSentToday int
}

type Contact struct {
	ID           uint `gorm:"primaryKey"`
	EmailAddress string
	SequenceID   uint
}

type EmailSendingLog struct {
	ID           uint `gorm:"primaryKey"`
	SequenceID   uint
	EmailAddress string
	DateTime     time.Time
}

// Handler is the Lambda handler function
func Handler(ctx context.Context) error {
	db, err := connectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	var mailboxes []Mailbox
	if err := db.Where("emails_sent_today < ?", 30).Find(&mailboxes).Error; err != nil {
		return fmt.Errorf("failed to get mailboxes: %v", err)
	}

	for _, mailbox := range mailboxes {
		if mailbox.EmailsSentToday >= 30 {
			continue
		}

		var contacts []Contact
		if err := getContactsForMailbox(db, &mailbox, &contacts); err != nil {
			return fmt.Errorf("failed to get contacts for mailbox %s: %v", mailbox.Address, err)
		}

		for _, contact := range contacts {
			// Send email logic here
			err := sendEmail(mailbox, contact)
			if err != nil {
				log.Printf("Failed to send email to %s: %v", contact.EmailAddress, err)
				continue
			}

			err = updateEmailLog(db, mailbox, contact)
			if err != nil {
				log.Printf("Failed to update email log: %v", err)
			}

			// Update the count of sent emails
			mailbox.EmailsSentToday++
			db.Save(&mailbox)

			if mailbox.EmailsSentToday >= 30 {
				break
			}
		}
	}

	return nil
}

// Connect to the PostgreSQL database using GORM
func connectToDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return db, nil
}

// Get contacts associated with the mailbox that are eligible to receive emails
func getContactsForMailbox(db *gorm.DB, mailbox *Mailbox, contacts *[]Contact) error {
	return db.Raw(`
		SELECT c.email_address, c.sequence_id
		FROM contacts c
		JOIN sequence_mailboxes sm ON sm.mailbox_id = ?
		WHERE c.sequence_id = sm.sequence_id`, mailbox.ID).Scan(contacts).Error
}

// Simulate sending an email
func sendEmail(mailbox Mailbox, contact Contact) error {
	// Email sending logic here (e.g., using AWS SES or any other email service)
	log.Printf("Sending email to %s using mailbox %s", contact.EmailAddress, mailbox.Address)
	return nil
}

// Update the email sending log in the database
func updateEmailLog(db *gorm.DB, mailbox Mailbox, contact Contact) error {
	logEntry := EmailSendingLog{
		SequenceID:   contact.SequenceID,
		EmailAddress: contact.EmailAddress,
		DateTime:     time.Now(),
	}
	return db.Create(&logEntry).Error
}

func main() {
	lambda.Start(Handler)
}
