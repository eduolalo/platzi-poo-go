package objects

type UberBlack struct {
	Car //<- Así indicamos que UberBlack hereda del struct Car
	TypeCarAccecpted map[string]map[string]int
	SeatMaterial     []string
}
