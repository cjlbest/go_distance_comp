package geometry

import (
    "fmt"
)

//
// struct Location : represents a (latitude, longitude) point on a sphere.
//
type Location struct {
    latitude  float64
    longitude float64
}

//
// func NewLocLatLon : is a struct Location factory method.
//
func NewLocLatLon(latitude float64, longitude float64) *Location {
    return &Location{latitude, longitude}
}

//
// func NewLocAddress : is a struct Location factory method.
//
func NewLocAddress(address *Address) *Location {
    return &Location{0,0}
}

//
// func Latitude : is a struct Location.latitude accessor
//
func (loc *Location) Latitude() float64 {
    return loc.latitude
}

//
// func Longitude : is a struct Location.longitude accessor
//
func (loc *Location) Longitude() float64 {
    return loc.longitude
}

//
// func String : converts struct Location to a String.
//
func (loc *Location) String() string {
    return fmt.Sprintf("(%f, %f)", loc.latitude, loc.longitude)
}

//
// func (loc *Location) DiffLoc(compLoc *Location) : calculates the distance
// between this Location and the input Location. The Haversine method is used
// to compute distance over the great arc between the two points.
//
// NOTE : this formula assumes that lat and long values are in DEGREES.
//
func (loc *Location) DiffLoc(compLoc *Location) float64 {

    return distHaversine(loc    .Latitude (),
                         loc    .Longitude(),
                         compLoc.Latitude (),
                         compLoc.Longitude())
}

//
// func (loc *Location) DiffAddr(compAddr *Address) : calculates the distance
// between this Location and the input Address. The Haversine method is used
// to compute distance over the great arc between the two points.
//
func (loc *Location) DiffAddr(compAddr *Address) float64 {

    // Convert the input Address to a Location for the calculation.

    return loc.DiffLoc(compAddr.GetLocation())
}
