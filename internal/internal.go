package internal

func GetMaturityIcon(maturityLevel int64) string {
	switch maturityLevel {
	case 3:
		return "\u2705"
	case 4:
		return "\u2705\u16ED"
	case 5:
		return "\u2705\u16ED\u16ED"
	default:
		return "\u274C"
	}
}
