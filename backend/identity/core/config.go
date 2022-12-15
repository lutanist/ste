package core

type IdentityOptions struct {
	Lockout LockoutOptions
	User    UserOptions
}

type LockoutOptions struct {
	AllowedForNewUsers bool
}

type UserOptions struct {
	AllowedUserNameCharacters string
	RequireUniqueEmail        bool
}

func NewUserOptions() UserOptions {
	return UserOptions{
		AllowedUserNameCharacters: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._@+",
		RequireUniqueEmail:        true,
	}
}
