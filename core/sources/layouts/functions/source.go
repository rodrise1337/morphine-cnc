package functions

import (
	"Morphine/core/clients/sessions"
	deployment "Morphine/core/configs"
	"Morphine/core/sources/language/evaluator"
	"Morphine/core/sources/language/lexer"
	"Morphine/core/sources/layouts/toml"
	"io"
)

func init() {

	RegisterFunction(&evaluator.Function{
		FunctionName: "cnc",
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
			return evaluator.ArrayObject(evaluator.Object{Literal: toml.ConfigurationToml.AppSettings.AppName, Type: lexer.String}), nil
		},
	})

	RegisterFunction(&evaluator.Function{
		FunctionName: "version",
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
			return evaluator.ArrayObject(evaluator.Object{Literal: deployment.Version, Type: lexer.String}), nil
		},
	})
	RegisterFunction(&evaluator.Function{
		FunctionName: "attackprefix",
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
			return evaluator.ArrayObject(evaluator.Object{Literal: toml.AttacksToml.Attacks.Prefix, Type: lexer.String}), nil
		},
	})
}
