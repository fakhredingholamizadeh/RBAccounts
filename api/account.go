package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/flfreecode/rbaccounts/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Username          string    `json:"username" binding:"required"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name" binding:"required"`
	Email             string    `json:"email" binding:"required"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	// One Of=UDST BTC ETH
	//	Email             string    `json:"email" binding:"required",One Of=USDT BTC ETH`

}

func (Server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil { // Should ctx have a JSON
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Username:          req.Username,
		HashedPassword:    "",
		FullName:          req.FullName,
		Email:             req.Email,
		PasswordChangedAt: time.Now(),
	}
	user, err := Server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Binging URI EXAMPLE ========= 	ID int64 `uri:"id" binding:"required,min=1"`

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (Server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := Server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// Binding QUERY EXAMPLE =====

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
