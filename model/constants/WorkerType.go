package constants

type WorkerType int

const (
	Barber      WorkerType = iota
	Plumber     WorkerType = iota
	Cleaner     WorkerType = iota
	Electrician WorkerType = iota
)

func String(w WorkerType) string {
	return []string{"Barber", "Plumber", "Cleaner", "Electrician"}[w]
}

func EnumIndex(worker WorkerType) int {
	return int(worker)
}
