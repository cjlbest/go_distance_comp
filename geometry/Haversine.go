package geometry

import (
    "math"
)

const PI          float64 = 3.1415926535897932384626433832795 // 2.0 * math.Asin(1)
const EarthRadius float64 = 6371.0

//
// func Diff : calculates the distance between this Location and the input
// Location. The Haversine method is used to compute distance over the great
// arc between the two points.
//
// NOTE : this formula assumes that lat and long values are in DEGREES.
//
func distHaversine(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {

    sinLat    := math.Sin(degreesToRadians(lat2 - lat1)/2.0)
    sinLon    := math.Sin(degreesToRadians(lon2 - lon1)/2.0)
    haversine := sinLat * sinLat +
                 sinLon * sinLon *
                 math.Cos(degreesToRadians(lat2)) * 
                 math.Cos(degreesToRadians(lat1))
    distance  := 2.0 * EarthRadius * math.Atan2(math.Sqrt(haversine),
                                                math.Sqrt(1.0-haversine))

    return distance
}

//
// func degressToRadians : converts the input decimal degree value into radians.
//
func degreesToRadians(degrees float64) float64 {
    radians := 2.0 * PI  * degrees / 360.0

    return radians
}
