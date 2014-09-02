package geometry

//
// struct Address : provides functionality to convert a generic "address" to a
// struct geometry.Location, i.e. a latitude,longitude pair.
//
type Address struct {
    // STUB.  TBD
}

//
// func NewAddress : is a struct Address factory method.
//
// NOTE : this function is currently a STUB and returns a default
//        Address, regarless of input.
//
func NewAddress(street string) *Address {
    // Just return a dummy for STUB
    return &Address{}
}

//
// func GetLocation : converts from struct Address to struct Location
//
// NOTE : this function is currently a STUB and returns a default
//        Location at lat=0, lon=0, regarless of input.
//
func (address *Address) GetLocation() *Location {

    // Just return a dummy for STUB

    return &Location{0,0}
}

//
// func String : converts struct Address to a String.
//
// NOTE : this function is currently a STUB and returns a default
//        empty string.
//
func (address *Address) String() string {
    // Just return a dummy for STUB
    return ""
}

//
// func (loc *Address) DiffAddr(compAddr *Address) : calculates the distance
// between this Address and the input Address. The Haversine method is used
// to compute distance over the great arc between the two points.
//
func (addr *Address) DiffAddr(compAddr *Address) float64 {

    // Convert both Addresses to locations for the calculation

    distance := addr.GetLocation().DiffLoc(compAddr.GetLocation())

    return distance
}

//
// func (loc *Address) DiffLoc(compLoc *Location) : calculates the distance
// between this Address and the input Location. The Haversine method is used
// to compute distance over the great arc between the two points.
//
// NOTE : this formula assumes that lat and long values are in DEGREES.
//
func (addr *Address) DiffLoc(compLoc *Location) float64 {

    // Convert this Address to a Location for the calculation

    distance := addr.GetLocation().DiffLoc(compLoc)

    return distance
}
