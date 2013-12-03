package Search

type Genome interface {
	Random()

	Represent() []float64

	Set(newGenome []float64) bool

	Get() []float64

	FitnessFunction() float64

	Lenght() int
}

type Population []Genome
