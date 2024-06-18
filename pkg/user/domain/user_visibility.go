package domain

type Visibility string

const (
	VisibilityUnknown            = ""
	VisibilityPrivate Visibility = "private"
	VisibilityPublic  Visibility = "public"
)

func (v Visibility) IsValid() bool {
	return v != VisibilityUnknown
}

func (v Visibility) String() string {
	return string(v)
}

func ToVisibility(value string) Visibility {
	switch value {
	case VisibilityPrivate.String():
		return VisibilityPrivate
	case VisibilityPublic.String():
		return VisibilityPublic
	default:
		return VisibilityUnknown
	}
}
