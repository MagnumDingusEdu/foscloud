package api

import (
	"database/sql"
	"fmt"
	db "foscloud/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type registerAccountRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) registerAccount(ctx *gin.Context) {
	var req registerAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := db.RegisterTxParams{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	_, err := server.store.GetAccountByUsername(ctx, args.Username)
	if err != sql.ErrNoRows {
		ctx.JSON(http.StatusConflict, errorResponse(fmt.Errorf("error : account with the same username already exists")))
		return
	}
	_, err = server.store.RegisterAccount(ctx, args)
	if err != nil {
		// check if account already exists (fallback)
		if err.(*pq.Error).Code == ("23505") {
			ctx.JSON(http.StatusConflict, errorResponse(fmt.Errorf("account with the same username already exists")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "Account created successfully",
	})
}

type loginAccountRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) loginAccount(ctx *gin.Context) {
	var res loginAccountRequest
	if err := ctx.ShouldBindJSON(&res); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	loginArgs := db.LoginAccountTxParams{
		LoginID:  res.Username,
		Password: res.Password,
	}

	loginResult, err := server.store.LoginAccountTx(ctx, loginArgs)
	if err != nil {
		if _, ok := err.(*db.IncorrectCredentialsError); ok {
			ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("Invalid credentials.")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, loginResult.Authtoken)
	return
}
