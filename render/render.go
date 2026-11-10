package render

func IdentifySymbol(symbol int) string {
	switch symbol {
	case 0:
		return "🏰"
	case 1:
		return "⚔️ "
	case 2:
		return "🔱"
	case 3:
		return "🏹"
	case 4:
		return "🎯"
	case 5:
		return "🐎"
	case 6:
		return "🏇"
	}
	return "⬛"
}
