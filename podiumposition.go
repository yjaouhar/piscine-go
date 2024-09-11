package piscine

func PodiumPosition(podium [][]string) [][]string {
	sl := len(podium)
	for i := 0; i < sl/2; i++ {
		podium[i], podium[sl-1-i] = podium[sl-1-i], podium[i]
	}
	return podium
}
