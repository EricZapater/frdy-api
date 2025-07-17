package users

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type UserRepository interface {	
	Create(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, id uuid.UUID) (error)
	ChangePassword(ctx context.Context, request ChangePasswordRequest) (User, error)
	FindByID(ctx context.Context, id uuid.UUID) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)	
	FindAll(ctx context.Context) ([]User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user User) (User, error) {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO users (id, email, username, password, is_active)
		VALUES ($1, $2, $3, $4, $5)`,
		user.ID,  user.Email, user.Username, user.Password, user.IsActive,
    )
    if err != nil {
        return User{}, fmt.Errorf("error inserting user: %w", err)
    }
    return user, nil
}

func(r *userRepository) Update(ctx context.Context, user User) (User, error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE users
		SET email = $1, username = $2, is_active = $3
		WHERE id = $4`,
		user.Email, user.Username, user.IsActive, user.ID,
	)
	if err != nil {
		return User{}, fmt.Errorf("error updating user: %w", err)
	}
	return user, nil
}

func(r *userRepository) Delete(ctx context.Context, id uuid.UUID) (error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE users
		SET is_active = false
		WHERE id = $1`,
		id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

func(r *userRepository) ChangePassword(ctx context.Context, request ChangePasswordRequest) (User, error){
	_, err := r.db.ExecContext(ctx, `
		UPDATE users
		SET password = $1, password_changed_at = now()
		WHERE id = $2`,
		request.Password, request.ID)
	if err != nil {
		return User{}, fmt.Errorf("error changing password: %w", err)
	}
	strId, err  := uuid.Parse(request.ID)
	if err != nil {
		return User{}, fmt.Errorf("error parsing id: %w", err)
	}
	user, err := r.FindByID(ctx, strId)
	if err != nil {
		return User{}, err
	}
	return user, nil
}



func(r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (User, error){
	var user User
	row := r.db.QueryRowContext(ctx, `SELECT id, email, username, password, is_active FROM users WHERE id = $1`, id)
	
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.IsActive)
	if err == sql.ErrNoRows {
		return User{}, ErrUserNotFound
	}else if err != nil {
		return User{}, fmt.Errorf("error getting user: %w", err)
	}
	if !user.IsActive {
		return User{}, ErrInactiveUser
	}
	return user, nil
}

func(r *userRepository) FindByUsername(ctx context.Context, username string) (User, error)	{
	var user User
	row := r.db.QueryRowContext(ctx, `SELECT id, email, username, password, is_active FROM users WHERE username = $1`, username)
	
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.IsActive)
	if err == sql.ErrNoRows {
		return User{}, ErrUserNotFound
	}else if err != nil {
		return User{}, fmt.Errorf("error getting user: %w", err)
	}
	if !user.IsActive {
		return User{}, ErrInactiveUser
	}
	return user, nil
}



func(r *userRepository) FindAll(ctx context.Context) ([]User, error){
	var users []User
	rows, err := r.db.QueryContext(ctx, `SELECT id, email, username, password, is_active FROM users`)
	if err != nil {
		return nil, fmt.Errorf("error getting users: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.IsActive)
		if err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}