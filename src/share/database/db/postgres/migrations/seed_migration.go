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
	err = seedCategories(client)
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

func seedCategories(client *sql.DB) error {
	categories := []models.Category{
		{
			Name: "Programming",
		},
		{
			Name: "Game",
		},
		{
			Name:           "Golang",
			ParentCategory: &models.Category{Name: "Programming"},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start a transaction
	tx, err := client.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Map to track category names and their IDs
	categoryIDs := make(map[string]int)

	// Insert categories
	for _, category := range categories {
		var parentID *int
		if category.ParentCategory != nil {
			pid, exists := categoryIDs[category.ParentCategory.Name]
			if !exists {
				return fmt.Errorf("parent category not found: %s", category.ParentCategory.Name)
			}
			parentID = &pid
		}

		var categoryID int
		query := "INSERT INTO category (name, parent_category_id) VALUES ($1, $2) RETURNING id"
		err := tx.QueryRowContext(ctx, query, category.Name, parentID).Scan(&categoryID)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return err
			}
			return err
		}

		categoryIDs[category.Name] = categoryID
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
