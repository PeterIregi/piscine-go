package piscine

func PodiumPosition(podium [][]string) [][]string {
	len := len(podium)

	for i := 0; i < len/2; i++ {
		j := len - 1 - i
		podium[i], podium[j] = podium[j], podium[i]
	}
	return podium
}
