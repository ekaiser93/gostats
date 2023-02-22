package calculations

import "math"

func GetAverages(b *Batter) {
	getBattingAvg(b)
	getOnBasePercentage(b)
	getSluggingPercentage(b)
	getOnBasePlusSlugging(b)
}

func toFixed(stat float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(int(stat*output)) / output
}

func getBattingAvg(b *Batter) {
	b.BattingAverage = toFixed(float64(b.Hits)/float64(b.AtBats), 3)
}

func getOnBasePercentage(b *Batter) {
	b.OnBasePercentage = toFixed(float64(b.Hits+b.Walks+b.IntentionalWalks+b.HitByPitch)/float64(b.AtBats+b.Walks+b.IntentionalWalks+b.SacrificeFly), 3)
}

func getSluggingPercentage(b *Batter) {
	b.SluggingPercentage = toFixed(float64(b.Hits+b.Doubles+(2*b.Triples)+(3*b.Homeruns))/float64(b.AtBats), 3)
}

func getOnBasePlusSlugging(b *Batter) {
	b.OnBasePlusSlugging = toFixed(b.OnBasePercentage+b.SluggingPercentage, 3)
}
