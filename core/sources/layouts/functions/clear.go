package functions

import (
	"Morphine/core/clients/sessions"
	"Morphine/core/sources/language/evaluator"
	"Morphine/core/sources/language/lexer"
	"io"
)

func init() {

	RegisterFunction(&evaluator.Function{
		FunctionName: "clear",
		//clears the screen properly and safer
		//this will make sure its done correctly without issues
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {

			//correctly tries to write without issues
			//this will make sure its done without issues making it safer
			return make([]evaluator.Object, 0), s.Write("\033c")
		},
	})
}
