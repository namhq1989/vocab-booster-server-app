package domain

type Action string

const (
	ActionUnknown              Action = ""
	ActionReviewSentence       Action = "reviewSentence"
	ActionDoVocabularyExercise Action = "doVocabularyExercise"
	ActionDoListeningExercise  Action = "doListeningExercise"
	ActionDoSpeakingExercise   Action = "doSpeakingExercise"
	ActionDoWritingExercise    Action = "doWritingExercise"
)

func (a Action) String() string {
	return string(a)
}

func (a Action) IsValid() bool {
	return a != ActionUnknown
}

func ToAction(value string) Action {
	switch value {
	case ActionReviewSentence.String():
		return ActionReviewSentence
	case ActionDoVocabularyExercise.String():
		return ActionReviewSentence
	case ActionDoListeningExercise.String():
		return ActionDoListeningExercise
	case ActionDoSpeakingExercise.String():
		return ActionDoSpeakingExercise
	case ActionDoWritingExercise.String():
		return ActionDoWritingExercise
	default:
		return ActionUnknown
	}
}

func (a Action) HasExceededLimit(plan string, value int) bool {
	dPlan := ToPlan(plan)
	if !dPlan.IsValid() {
		dPlan = PlanFree
	}

	var (
		isPremium        = dPlan.IsPremium()
		hasExceededLimit bool
	)

	switch a {
	case ActionReviewSentence:
		if isPremium {
			hasExceededLimit = value > 20
		} else {
			hasExceededLimit = value > 5
		}
	case ActionDoVocabularyExercise:
		if isPremium {
			hasExceededLimit = value > 10
		} else {
			hasExceededLimit = value > 2
		}
	case ActionDoListeningExercise:
		if isPremium {
			hasExceededLimit = value > 10
		} else {
			hasExceededLimit = value > 2
		}
	case ActionDoSpeakingExercise:
		if isPremium {
			hasExceededLimit = value > 10
		} else {
			hasExceededLimit = value > 2
		}
	case ActionDoWritingExercise:
		if isPremium {
			hasExceededLimit = value > 10
		} else {
			hasExceededLimit = value > 2
		}
	default:
		hasExceededLimit = true
	}

	return hasExceededLimit
}
