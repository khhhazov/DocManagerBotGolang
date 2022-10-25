package calculation

func CalcTaskTypeWeight(buttonData string) int {
	switch buttonData {
	case "Учебная":
		return 2
	case "Научная":
		return 2
	case "Административная":
		return 3
	case "Культурно массовая":
		return 1
	default:
		return 0
	}
}
