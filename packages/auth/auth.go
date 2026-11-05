package auth

import "errors"

// PUBLIC Struct: Capital 'U' so main can use it
type User struct {
	Username string // PUBLIC Field (Capital 'U')
	Email    string // PUBLIC Field (Capital 'E')
	password string // PRIVATE Field (lowercase 'p' -> Hidden from main!)
}

// PRIVATE Helper: lowercase 'h' -> Can ONLY be used inside auth package
func hashPassword(raw string) string {
	return "hashed_secret_" + raw
}

// PUBLIC Function: Capital 'R' -> main can call auth.RegisterUser()
func RegisterUser(username, email, rawPassword string) (*User, error) {
	if len(rawPassword) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	return &User{
		Username: username,
		Email:    email,
		password: hashPassword(rawPassword),
	}, nil
}

// PUBLIC Method: Allows verifying password safely without exposing the raw field
func (u *User) VerifyPassword(input string) bool {
	return u.password == hashPassword(input)
}