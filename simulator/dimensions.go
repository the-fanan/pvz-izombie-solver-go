package simulator

import "math"

const CELL_LENGTH_MICRO_METER = float64(1180122)
const MAX_HEIGHT_FOR_POLE_VAULT_ZOMBIE_MICRO_METER = float64(810000)
const MAX_HEIGHT_FOR_CATAPULT_ITEMS_IN_MICRO_METER = float64(2000000)
const TIME_TO_TRAVERSE_CELL_FOR_FASTEST_OBJECT_MILLI_SECOND = float64(250)

func SubCellsPerMicroMeter(resolution int) float64 {
	return float64(resolution)/ CELL_LENGTH_MICRO_METER
}

func FramePerMilliSecond(resolution int) float64 {
	return float64(resolution)/ TIME_TO_TRAVERSE_CELL_FOR_FASTEST_OBJECT_MILLI_SECOND
}

func MaxProjectileHeightInSubCells(maxHeightInMicrometer float64, resolution int) float64 {
	return (maxHeightInMicrometer / CELL_LENGTH_MICRO_METER) * float64(resolution)
}

func AccelerationDueToGravityInSubcellsPerSquareFrames(resolution int) float64 {
	fpm := FramePerMilliSecond(resolution)
	spm := SubCellsPerMicroMeter(resolution)

	return (9.8 * spm) / math.Pow(fpm, 2)
}

func StartingAngleForProjectileMotion(maxHeightInSubCells, distanceInSubCells float64) float64 {
	return math.Atan((4 * maxHeightInSubCells) / distanceInSubCells)
}

func ProjectileVelocity(resolution int, maxHeightInSubCells, angle float64) float64 {
	adtg := AccelerationDueToGravityInSubcellsPerSquareFrames(resolution)
	return math.Sqrt((2 * maxHeightInSubCells * adtg) / math.Pow(math.Sin(angle), 2))
}

func ProjectileVerticalDisplacement(resolution, frame int, velocity, angle float64) float64 {
	adtg := AccelerationDueToGravityInSubcellsPerSquareFrames(resolution)
	return (velocity * float64(frame) * math.Sin(angle)) - (0.5 * adtg * math.Pow(float64(frame), 2))
}

func ProjectileHorizontalDisplacement(frame int, velocity, angle float64) float64 {
	return velocity * float64(frame) * math.Cos(angle)
}

// positions hold just co-ordinate values but we need to convert them to actual linear
// distances relative to start of board to perform calculations
// we don't bother about row since once and object is in motion that is not expected to change
// except for starfish attacks which will be dealt with later
func GetHorizontalDistanceFromCellColAndSubCellCol(cellCol, subCellCol, resolution int) int {
	return (cellCol * resolution) + subCellCol
}

func GetCellColAndSubCellColFromHorizontalDistance(distance, resolution int) (int, int) {
	newCol := distance / resolution
	newSubCol := distance % resolution
	// we have gone out of bounds and integer division will give us 0 instead of -1
	// which would lead to erroneous results
	if distance < 0 {
		newCol = -1
	}
	return newCol, newSubCol
}
