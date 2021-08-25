package languages

type Language string

var SupportedLanguages = []Language{
	"Italiano",
	"English",
}

func IsSupported(lang Language) bool {
	for _, l := range SupportedLanguages {
		if lang == l {
			return true
		}
	}
	return false
}

type Message string

var Messages = map[Message]map[Language][]string{
	"Welcoming": {
		"Italiano": {"Ciao! Io sono cardGames, un bot che raccoglie diversi giochi di carte. Sentiti libero di giocare con i tuoi amici o contro di me!\n\nUsa /help per una lista di tutti i comandi.\n\n\n- Sviluppatore: @MoraGames"},
		"English":  {"Hi! I'm cardGames, a bot that collects different card games. Feel free to play with your friends or against me!\n\nUse /help for a list of all commands.\n\n\n- Developer: @MoraGames"},
	},
	"Commands": {
		"Italiano": {"Questa è la lista di tutti i comandi del bot:\n\n"},
		"English":  {"This is the list of all the commands of the bot:\n\n"},
	},
	"Languages": {
		"Italiano": {
			"These are all languages supported by the bot:\n\n",
			"Se vuoi collaborare per aggiungere nuove lingue o perfezionarne alcune usa /collaborate.",
		},
		"English": {
			"Here are all the participants in the development of the bot:\n\n",
			"If you want to collaborate to add new languages or improve some of them use /collaborate.",
		},
	},
	"Credits": {
		"Italiano": {
			"Ecco tutti i partecipanti allo sviluppo del bot:\n\n",
			"Se vuoi dare il tuo contributo usa /collaborate.",
		},
		"English": {
			"Here are all the participants in the development of the bot:\n\n",
			"If you want to give your contribution use / collaborate.",
		},
	},
}
