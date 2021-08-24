package users

type user string

type settings struct {
	language language
}

var usersSettings = map[user]settings{
	"MoraGames": {
		language: "Italiano",
	},
}
