package service

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func InitOAuthProviders() {
	goth.UseProviders(google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), os.Getenv("GOOGLE_CALLBACK_URL"), "email", "profile"))
}
