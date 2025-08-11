package repository

import (
	"context"
	"fmt"
	"gameplatform/internal/database"
	"gameplatform/internal/dbconn"
	"log/slog"

	// "gameplatform/internal/models"
	// "slices"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CreateUser struct {
	Name             string
	Email            string
	Password         []byte
	IsAdmin          bool
	VerificationCode string
	Verified         bool
	Birthday         *time.Time
	Gender           *string
}

type GetUser struct {
	ID               uuid.UUID
	Name             string
	Email            string
	Password         string
	IsAdmin          bool
	VerificationCode string
	Verified         bool
	Birthday         *time.Time
	Gender           *string
	CreatedAt        time.Time
}

type UpdateUser struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
	IsAdmin  bool
	Verified bool
	Birthday *time.Time
	Gender   *string
}

// func PlatformUserToGetUser(u *database.PlatformUser) GetUser {
// 	return GetUser{
// 		ID:               u.ID,
// 		Name:             u.Name,
// 		Email:            u.Email,
// 		Password:         string(u.Password),
// 		IsAdmin:          u.IsAdmin,
// 		VerificationCode: u.VerificationCode,
// 		Verified:         u.Verified,
// 		Birthday:         u.Birthday,
// 		Gender:           u.Gender,
// 		CreatedAt:        u.CreatedAt,
// 	}
// }
//
// func PlatformUsersToGetUsers(platfom_users []database.PlatformUser) []GetUser {
// 	length := len(platfom_users)
// 	users := make([]GetUser, length)
// 	for i := 0; i < length; i++ {
// 		users[i] = PlatformUserToGetUser(&platfom_users[i])
// 	}
//
// 	return users
// }

type UserRepository struct {
	db   *dbconn.DatabaseConnection
	conv UserConverter
}

func NewUserRepository(db *dbconn.DatabaseConnection, conv UserConverter) *UserRepository {
	return &UserRepository{db: db, conv: conv}
}

func (r *UserRepository) Create(user CreateUser) (*GetUser, error) {
	ctx := context.Background()
	user_platform, err := r.db.Queries.CreateUser(ctx, r.conv.CreateUserToCreateUserParams(&user))
	if err != nil {
		return nil, SqlcErrToRepositoryErr(err)
	}
	get_user := r.conv.PlatformUserToGetUser(&user_platform)

	return &get_user, SqlcErrToRepositoryErr(err)
}

func (r *UserRepository) Update(user *UpdateUser) error {
	ctx := context.Background()
	params := database.UpdateUserByIdParams{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: string(user.Password),
		IsAdmin:  user.IsAdmin,
		Verified: user.Verified,
		Birthday: user.Birthday,
		Gender:   user.Gender,
	}

	return SqlcErrToRepositoryErr(r.db.Queries.UpdateUserById(ctx, params))
}

func (r *UserRepository) UpdateUserVerification(id uuid.UUID, verification_code string, verified bool) error {
	ctx := context.Background()
	err := r.db.Queries.UpdateUserVerification(ctx, database.UpdateUserVerificationParams{
		ID:               id,
		VerificationCode: verification_code,
		Verified:         verified,
	})

	return SqlcErrToRepositoryErr(err)
}

func (r *UserRepository) UpdatePassword(id uuid.UUID, password string) error {
	ctx := context.Background()
	err := r.db.Queries.UpdateUserPassword(ctx, database.UpdateUserPasswordParams{
		ID:       id,
		Password: password,
	})

	return SqlcErrToRepositoryErr(err)
}

func (r *UserRepository) GetByVerificationCode(verificationCode string) (*GetUser, error) {
	ctx := context.Background()
	platform_user, err := r.db.Queries.GetUserByVerificationCode(ctx, verificationCode)
	err = SqlcErrToRepositoryErr(err)

	if err != nil {
		return nil, err
	}
	user := r.conv.PlatformUserToGetUser(&platform_user)
	return &user, nil
}

func (r *UserRepository) CreateUserWithTransaction(user CreateUser, callback func(CreateUser) error) (*GetUser, error) {
	ctx := context.Background()
	tx, err := r.db.Pool.BeginTx(ctx, pgx.TxOptions{})
	tx_queries := r.db.Queries.WithTx(tx)

	user_platform, err := tx_queries.CreateUser(ctx, r.conv.CreateUserToCreateUserParams(&user))
	if err != nil {
		return nil, SqlcErrToRepositoryErr(err)
	}
	err = callback(user)
	if err == nil {
		tx.Commit(ctx)
	} else {
		tx.Rollback(ctx)
	}

	get_user := r.conv.PlatformUserToGetUser(&user_platform)
	return &get_user, err
}

func (r *UserRepository) GetByEmail(email string) (*GetUser, error) {
	ctx := context.Background()
	user, err := r.db.Queries.GetUserByEmail(ctx, email)
	err = SqlcErrToRepositoryErr(err)
	if err != nil {
		return nil, err
	}
	result := r.conv.PlatformUserToGetUser(&user)
	return &result, nil
}

func (r *UserRepository) GetUserById(id string) (user *GetUser, err error) {
	ctx := context.Background()
	user_id, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't parse user uuid. UUID: %v, Error: %v", user_id, err.Error()))
		return nil, Error
	}

	p_user, err := r.db.Queries.GetUserByID(ctx, user_id)
	db_err := err
	err = SqlcErrToRepositoryErr(err)
	if err != nil {
		slog.Error(fmt.Sprintf("Error: %v, ID: %v", db_err.Error(), user_id))
		return nil, err
	}
	result := r.conv.PlatformUserToGetUser(&p_user)
	return &result, nil
}

func (r *UserRepository) GetAll() ([]GetUser, error) {
	ctx := context.Background()
	platform_users, err := r.db.Queries.GetAllUsers(ctx)
	err = SqlcErrToRepositoryErr(err)
	if err != nil {
		return nil, err
	}

	return r.conv.PlatformUsersToGetUsers(platform_users), nil
}

func (r *UserRepository) DeleteUser(id string) error {
	ctx := context.Background()
	user_id, err := uuid.Parse(id)
	if err != nil {
		return Error
	}

	return SqlcErrToRepositoryErr(r.db.Queries.DeleteUserById(ctx, user_id))
}
