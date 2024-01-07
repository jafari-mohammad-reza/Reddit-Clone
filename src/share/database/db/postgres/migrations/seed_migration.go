package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/reddit-clone/src/share/common"
	"github.com/reddit-clone/src/share/database/db/models"
	"github.com/reddit-clone/src/share/database/db/postgres"
)

func SeedData() error {
	client := postgres.GetPostgres()
	err := seedUsers(client)
	if err != nil {
		return err
	}
	return nil
}

func seedUsers(client *sql.DB) error {
	// Define users
	users := []models.User{
		{
			Email:    "admin@gmail.com",
			Password: common.HashPassword("Admin@123"),
		},
		{
			Email:    "subreddit-owner@gmail.com",
			Password: common.HashPassword("Owner@123"),
		},
		{
			Email:    "user@gmail.com",
			Password: common.HashPassword("User@123"),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Prepare the SQL statement with ON CONFLICT clause
	stmt, err := client.PrepareContext(ctx, "INSERT INTO \"user\" (email, password) VALUES ($1, $2) ON CONFLICT (email) DO NOTHING")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement for each user
	for _, u := range users {
		_, err := stmt.ExecContext(ctx, u.Email, u.Password)
		if err != nil {
			fmt.Print("SEED", err.Error())
			return err
		}
	}

	return nil
}
