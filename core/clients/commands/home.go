package commands

import (
	"Morphine/core/clients/sessions"
	deployment "Morphine/core/configs"
	"Morphine/core/sources/language"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "home",
		Aliases:            []string{"welcome"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "redirect your self back home",
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			//renders the clear splash information properly and safely without issues happening on request
			return language.ExecuteLanguage([]string{"welcome.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
		},
	})
}
