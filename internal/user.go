package entity

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	// For a matter of security we never save the password string to the db
	// We are using the bcrypt package to generate a hash, which later on
	// will be used to compare to the given one...
	// Never comparing the password itself, but the hash.
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       "0",
		Name:     name,
		Email:    email,
		Password: string(hash)}

	return user, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
