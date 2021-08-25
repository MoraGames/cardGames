package users

import "github.com/MoraGames/cardGames/pkg/internal/languages"

type User string

type Settings struct {
	language languages.Language
}

var UsersSettings = map[User]Settings{
	"MoraGames": {
		language: "Italiano",
	},
}
