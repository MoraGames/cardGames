module github.com/MoraGames/cardGames

go 1.16

require pkg/internal/languages v0.0.0
replace pkg/internal/languages => ./pkg/internal/languages
require pkg/internal/users v0.0.0
replace pkg/internal/users => ./pkg/internal/users
require pkg/internal/games v0.0.0
replace pkg/internal/games => ./pkg/internal/games
