package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/AbdelilahOu/GoferQl/config"
	db "github.com/AbdelilahOu/GoferQl/internal/db/sqlc"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func generateUsers(count int) []db.CreateUserParams {
	users := make([]db.CreateUserParams, count)
	for i := 0; i < count; i++ {
		users[i] = db.CreateUserParams{
			Username: faker.Username(),
			Email:    faker.Email(),
			Password: hashPassword("password123"),
			Bio: pgtype.Text{
				String: faker.Sentence(),
				Valid:  true,
			},
		}
	}
	return users
}

func generateCategories(count int) []db.CreateCategoryParams {
	categories := make([]db.CreateCategoryParams, count)
	categoryTypes := []string{
		"Technology", "Lifestyle", "Travel", "Food", "Fashion",
		"Sports", "Health", "Business", "Entertainment", "Science",
		"Art", "Music", "Photography", "Gaming", "Education",
	}

	for i := 0; i < count; i++ {
		catName := categoryTypes[i%len(categoryTypes)]
		if i >= len(categoryTypes) {
			catName = fmt.Sprintf("%s %d", categoryTypes[i%len(categoryTypes)], i/len(categoryTypes))
		}

		categories[i] = db.CreateCategoryParams{
			Name: catName,
			Description: pgtype.Text{
				String: faker.Sentence(),
				Valid:  true,
			},
		}
	}
	return categories
}

func generateTags(count int) []db.Tag {
	tags := make([]db.Tag, count)
	techWords := []string{
		"programming", "development", "coding", "software", "tech",
		"web", "mobile", "data", "cloud", "security", "ai", "ml",
		"database", "frontend", "backend", "devops", "architecture",
	}

	for i := 0; i < count; i++ {
		tagName := techWords[i%len(techWords)]
		if i >= len(techWords) {
			tagName = fmt.Sprintf("%s-%d", techWords[i%len(techWords)], i/len(techWords))
		}

		tags[i] = db.Tag{
			ID:   uuid.New(),
			Name: tagName,
		}
	}
	return tags
}

func generatePosts(count int, users []db.User, categories []db.Category) []db.CreatePostParams {
	posts := make([]db.CreatePostParams, count)
	statuses := []string{"draft", "published", "archived"}

	for i := 0; i < count; i++ {
		posts[i] = db.CreatePostParams{
			Title:   faker.Sentence(),
			Content: generateArticleContent(),
			UserID: pgtype.UUID{
				Bytes: users[rand.Intn(len(users))].ID,
				Valid: true,
			},
			CategoryID: pgtype.UUID{
				Bytes: categories[rand.Intn(len(categories))].ID,
				Valid: true,
			},
			Status: pgtype.Text{
				String: statuses[rand.Intn(len(statuses))],
				Valid:  true,
			},
		}
	}
	return posts
}

func generateComments(
	ctx context.Context,
	queries *db.Queries,
	posts []db.Post,
	users []db.User,
	count int,
) {
	fmt.Println("Generating comments...")

	for _, post := range posts {
		for i := 0; i < count; i++ {
			commentParams := db.CreateCommentParams{
				Content: faker.Sentence(),
				PostID: pgtype.UUID{
					Bytes: post.ID,
					Valid: true,
				},
				UserID: pgtype.UUID{
					Bytes: users[rand.Intn(len(users))].ID,
					Valid: true,
				},
				ParentID: pgtype.UUID{
					Valid: false,
				},
			}

			comment, err := queries.CreateComment(ctx, commentParams)
			if err != nil {
				log.Printf("Error creating comment for post %s: %v", post.ID, err)
				continue
			}

			if rand.Float64() < rand.Float64() {
				numReplies := rand.Intn(3) + 1
				for j := 0; j < numReplies; j++ {
					replyParams := db.CreateCommentParams{
						Content: faker.Sentence(),
						PostID: pgtype.UUID{
							Bytes: post.ID,
							Valid: true,
						},
						UserID: pgtype.UUID{
							Bytes: users[rand.Intn(len(users))].ID,
							Valid: true,
						},
						ParentID: pgtype.UUID{
							Bytes: comment.ID,
							Valid: true,
						},
					}

					_, err := queries.CreateComment(ctx, replyParams)
					if err != nil {
						log.Printf("Error creating reply for comment %s: %v", comment.ID, err)
						continue
					}
				}
			}
		}
	}
}

func generateArticleContent() string {
	paragraphs := rand.Intn(5) + 2
	content := ""
	for i := 0; i < paragraphs; i++ {
		content += faker.Paragraph() + "\n\n"
	}
	return content
}

func hashPassword(password string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedBytes)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	pgPool, err := pgxpool.New(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer pgPool.Close()

	queries := db.New(pgPool)
	ctx := context.Background()

	fmt.Println("Generating users...")
	users := generateUsers(50)
	createdUsers := make([]db.User, 0, 50)
	for _, user := range users {
		createdUser, err := queries.CreateUser(ctx, user)
		if err != nil {
			log.Printf("Error creating user %s: %v", user.Username, err)
			continue
		}
		createdUsers = append(createdUsers, createdUser)
	}

	fmt.Println("Generating categories...")
	categories := generateCategories(10)
	createdCategories := make([]db.Category, 0, 10)
	for _, category := range categories {
		createdCategory, err := queries.CreateCategory(ctx, category)
		if err != nil {
			log.Printf("Error creating category %s: %v", category.Name, err)
			continue
		}
		createdCategories = append(createdCategories, createdCategory)
	}

	fmt.Println("Generating tags...")
	tags := generateTags(20)
	createdTags := make([]db.Tag, 0, 20)
	for _, tag := range tags {
		createdTag, err := queries.CreateTag(ctx, tag.Name)
		if err != nil {
			log.Printf("Error creating tag %s: %v", tag.Name, err)
			continue
		}
		createdTags = append(createdTags, createdTag)
	}

	if len(createdUsers) == 0 || len(createdCategories) == 0 || len(createdTags) == 0 {
		log.Fatal("Missing required data: ensure users, categories, and tags were created successfully")
	}

	fmt.Println("Generating posts and assigning tags...")
	posts := generatePosts(100, createdUsers, createdCategories)
	createdPosts := make([]db.Post, 0, 100)
	for _, post := range posts {
		createdPost, err := queries.CreatePost(ctx, post)
		if err != nil {
			log.Printf("Error creating post %s: %v", post.Title, err)
			continue
		}
		createdPosts = append(createdPosts, createdPost)

		numTags := rand.Intn(5) + 1
		usedTagIndices := make(map[int]bool)

		maxTags := min(numTags, len(createdTags))
		for i := 0; i < maxTags; i++ {
			var tagIndex int
			maxAttempts := len(createdTags) * 2
			attempts := 0

			for {
				if attempts >= maxAttempts {
					log.Printf("Could not find unused tag for post %s after %d attempts", createdPost.ID, attempts)
					break
				}

				tagIndex = rand.Intn(len(createdTags))
				if !usedTagIndices[tagIndex] {
					usedTagIndices[tagIndex] = true
					break
				}
				attempts++
			}

			if tagIndex >= len(createdTags) {
				log.Printf("Invalid tag index %d for post %s", tagIndex, createdPost.ID)
				continue
			}

			err := queries.AddPostTag(ctx, db.AddPostTagParams{
				PostID: createdPost.ID,
				TagID:  createdTags[tagIndex].ID,
			})
			if err != nil {
				log.Printf("Error creating post-tag relation for post %s and tag %s: %v",
					createdPost.ID, createdTags[tagIndex].ID, err)
				continue
			}
		}
	}
	generateComments(ctx, queries, createdPosts, createdUsers, 30)

	fmt.Println("Seeding completed successfully!")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
