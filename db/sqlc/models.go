// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// hashed password
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	LastLogin sql.NullTime `json:"last_login"`
}

type Link struct {
	ID        int64          `json:"id"`
	Node      int64          `json:"node"`
	Link      string         `json:"link"`
	Clicks    sql.NullInt32  `json:"clicks"`
	Password  sql.NullString `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
}

type Node struct {
	ID int64 `json:"id"`
	// parent_id is null for root node
	ParentID sql.NullInt64 `json:"parent_id"`
	Name     string        `json:"name"`
	IsDir    bool          `json:"is_dir"`
	Filesize sql.NullInt64 `json:"filesize"`
	// depth starting from parent node (0)
	Depth sql.NullInt32 `json:"depth"`
	// used for breadcrumbs
	Lineage   sql.NullString `json:"lineage"`
	Owner     sql.NullInt64  `json:"owner"`
	CreatedAt time.Time      `json:"created_at"`
}