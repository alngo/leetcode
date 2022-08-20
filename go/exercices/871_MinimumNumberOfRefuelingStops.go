package exercices

const (
    POSITION = iota
    FUEL
)

func MinRefuelStops(target int, startFuel int, stations [][]int) int {
    fuelPaths := make([]int, len(stations)+1)

    for i := range fuelPaths {
        fuelPaths[i] = startFuel
    }

    for i := 0; i < len(stations); i++ {
        station := stations[i]
        for j := i; j >= 0 && isWithinRange(fuelPaths[j], station[POSITION]); j-- {
            newFuelPath := fuelPaths[j] + station[FUEL]
            oldFuelPath := fuelPaths[j + 1]
            if newFuelPath > oldFuelPath {
                fuelPaths[j + 1] = newFuelPath
            } 
        }
    }

    for i := range fuelPaths {
        if fuelPaths[i] >= target {
            return i
        }
    }

    return -1
}

func isWithinRange(fuel int, position int) bool {
    return fuel >= position
}
