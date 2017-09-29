-- +goose Up
-- +goose StatementBegin

ALTER TABLE "Users"
	RENAME COLUMN "createdAt" TO created_at;
ALTER TABLE "Users" 
	RENAME COLUMN "deletedAt" TO deleted_at;
ALTER TABLE "Users" 
	RENAME COLUMN "updatedAt" TO updated_at;
ALTER TABLE "Users" RENAME TO users;

ALTER TABLE "Books"
	RENAME COLUMN "createdAt" TO created_at;
ALTER TABLE "Books" 
	RENAME COLUMN "deletedAt" TO deleted_at;
ALTER TABLE "Books" 
	RENAME COLUMN "updatedAt" TO updated_at;
ALTER TABLE "Books" RENAME TO books;

ALTER TABLE "Chapter_Paths"
	RENAME COLUMN "createdAt" TO created_at;
ALTER TABLE "Chapter_Paths" 
	RENAME COLUMN "updatedAt" TO updated_at;   
ALTER TABLE "Chapter_Paths" RENAME TO chapter_paths;

ALTER TABLE "Chapters"
	RENAME COLUMN "createdAt" TO created_at;
ALTER TABLE "Chapters" 
	RENAME COLUMN "deletedAt" TO deleted_at;
ALTER TABLE "Chapters" 
	RENAME COLUMN "updatedAt" TO updated_at;
ALTER TABLE "Chapters" RENAME TO chapters;


-- +goose StatementEnd
