package domain

type Action string

const (
	ActionUnknown                      Action = ""
	ActionAnswerExercise               Action = "answerExercise"
	ActionContributeVocabularySentence Action = "contributeVocabularySentence"
)

func (a Action) IsValid() bool {
	return a != ActionUnknown
}

func (a Action) String() string {
	return string(a)
}

func ToAction(value string) Action {
	switch value {
	case ActionAnswerExercise.String():
		return ActionAnswerExercise
	case ActionContributeVocabularySentence.String():
		return ActionContributeVocabularySentence
	default:
		return ActionUnknown
	}
}

func (a Action) IsAnswerExercise() bool {
	return a == ActionAnswerExercise
}

func (a Action) IsContributeVocabularySentence() bool {
	return a == ActionContributeVocabularySentence
}
