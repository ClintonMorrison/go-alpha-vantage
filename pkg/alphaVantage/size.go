package alphaVantage

type Size string
const (
	SIZE_COMPACT Size = "compact"
	SIZE_FULL Size = "full"
)

func sizeFromString(s string) Size {
	switch s {
	case "Full size": return SIZE_FULL
	default: return SIZE_COMPACT

	}
}
