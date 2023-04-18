package geo

import (
	"math"
)

/*
*
boundingBox calculates and returns the four corners of a bounding box determined by the given distance
latitude := 37.7749
longitude := -122.4194
distance := 10.0 // miles

minLat, minLon, minLat2, maxLon, maxLat, minLon2, maxLat2, maxLon2 := boundingBox(latitude, longitude, distance)
fmt.Printf("Bounding Box Coordinates:\n")
fmt.Printf("(%.6f, %.6f), (%.6f, %.6f), (%.6f, %.6f), (%.6f, %.6f)\n", minLat, minLon, minLat2, maxLon, maxLat, minLon2, maxLat2, maxLon2)
*/
func boundingBox(latitude, longitude, distance float64) (float64, float64, float64, float64, float64, float64, float64, float64) {
	const earthRadius = 3960.0 // Radius of Earth in miles
	const radianConversion = math.Pi / 180.0

	// Convert latitude and longitude to radians
	latRad := latitude * radianConversion
	lonRad := longitude * radianConversion

	// Calculate the distance in radians
	distanceRad := distance / earthRadius

	// Calculate the bounding box coordinates
	minLat := latRad - distanceRad
	maxLat := latRad + distanceRad
	minLon := lonRad - math.Asin(math.Sin(distanceRad)/math.Cos(latRad))
	maxLon := lonRad + math.Asin(math.Sin(distanceRad)/math.Cos(latRad))

	// Convert the bounding box coordinates back to degrees
	minLat = minLat / radianConversion
	maxLat = maxLat / radianConversion
	minLon = minLon / radianConversion
	maxLon = maxLon / radianConversion

	return minLat, minLon, minLat, maxLon, maxLat, minLon, maxLat, maxLon
}

func haversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 3960.0 // Radius of Earth in miles
	const radianConversion = math.Pi / 180.0

	// Convert latitude and longitude to radians
	lat1Rad := lat1 * radianConversion
	lon1Rad := lon1 * radianConversion
	lat2Rad := lat2 * radianConversion
	lon2Rad := lon2 * radianConversion

	// Calculate the differences between the coordinates
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	// Calculate the Haversine distance
	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func midpoint(lat1, lon1, lat2, lon2 float64) (float64, float64) {
	const radianConversion = math.Pi / 180.0

	// Convert latitude and longitude to radians
	lat1Rad := lat1 * radianConversion
	lon1Rad := lon1 * radianConversion
	lat2Rad := lat2 * radianConversion
	lon2Rad := lon2 * radianConversion

	// Calculate the midpoint
	Bx := math.Cos(lat2Rad) * math.Cos(lon2Rad-lon1Rad)
	By := math.Cos(lat2Rad) * math.Sin(lon2Rad-lon1Rad)

	midLatRad := math.Atan2(math.Sin(lat1Rad)+math.Sin(lat2Rad), math.Sqrt((math.Cos(lat1Rad)+Bx)*(math.Cos(lat1Rad)+Bx)+By*By))
	midLonRad := lon1Rad + math.Atan2(By, math.Cos(lat1Rad)+Bx)

	// Convert the midpoint coordinates back to degrees
	midLat := midLatRad / radianConversion
	midLon := midLonRad / radianConversion

	return midLat, midLon
}

func initialBearing(lat1, lon1, lat2, lon2 float64) float64 {
	const radianConversion = math.Pi / 180.0

	// Convert latitude and longitude to radians
	lat1Rad := lat1 * radianConversion
	lon1Rad := lon1 * radianConversion
	lat2Rad := lat2 * radianConversion
	lon2Rad := lon2 * radianConversion

	// Calculate the initial bearing
	y := math.Sin(lon2Rad-lon1Rad) * math.Cos(lat2Rad)
	x := math.Cos(lat1Rad)*math.Sin(lat2Rad) - math.Sin(lat1Rad)*math.Cos(lat2Rad)*math.Cos(lon2Rad-lon1Rad)
	initialBearingRad := math.Atan2(y, x)

	// Convert the initial bearing back to degrees
	initialBearing := initialBearingRad / radianConversion

	// Normalize the bearing to a value between 0 and 360
	return math.Mod(initialBearing+360, 360)
}
