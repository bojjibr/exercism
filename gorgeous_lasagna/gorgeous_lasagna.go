package gorgeouslasagna

const Oventime = 40 // 40 minutes

func RemainingOvenTime(actual int) int {
	return Oventime - actual
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
