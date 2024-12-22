package main

import (
	"context"
	"log"

	"github.com/nohattee/spidercat/src/scraper/internal/adapter/repository"
	"github.com/nohattee/spidercat/src/scraper/internal/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()
	dsn := "host=postgresql user=spidercat password=spidercat dbname=spidercat port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	itemRepo := repository.NewItemRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	authorRepo := repository.NewAuthorRepository(db)
	tagRepo := repository.NewTagRepository(db)

	uc := usecase.New(itemRepo, categoryRepo, authorRepo, tagRepo)
	err = uc.ScrapeURLs(ctx, []string{"https://truyenqqto.com/truyen-moi-cap-nhat.html"})
	if err != nil {
		log.Fatal(err)
	}
}
