PACKAGE DOCUMENTATION

package geometry
    import "distance_comp/geometry"

    package geometry : provides functionality for location operations. The
    package provides structs for converting between differnt types of
    location representations, like Location and Address, to facilitate
    computation of distance between any two location types.

CONSTANTS

const EarthRadius float64 = 6371.0

const PI float64 = 3.1415926535897932384626433832795 // 2.0 * math.Asin(1)


FUNCTIONS

func CalcAllWithCompOp(inFilePath string,
    outFilePath string,
    test *testing.T,
    compOp func(int, float64, float64, float64, float64, *testing.T) (float64, error)) error
    func CalcAllWithCompOp : reads input baseline data file, calculates
    distance using the input compOp function, compares the calculate value
    with values in the baseline file, and writes any error to the output
    file.

    NOTE : this function is used by service_test and geometry_test.

TYPES

type Address struct {
}
    struct Address : provides functionality to convert a generic "address"
    to a struct geometry.Location, i.e. a latitude,longitude pair.

func NewAddress(street string) *Address
    func NewAddress : is a struct Address factory method.

    NOTE : this function is currently a STUB and returns a default

	Address, regarless of input.

func (addr *Address) DiffAddr(compAddr *Address) float64
    func (loc *Address) DiffAddr(compAddr *Address) : calculates the
    distance between this Address and the input Address. The Haversine
    method is used to compute distance over the great arc between the two
    points.

func (addr *Address) DiffLoc(compLoc *Location) float64
    func (loc *Address) DiffLoc(compLoc *Location) : calculates the distance
    between this Address and the input Location. The Haversine method is
    used to compute distance over the great arc between the two points.

    NOTE : this formula assumes that lat and long values are in DEGREES.

func (address *Address) GetLocation() *Location
    func GetLocation : converts from struct Address to struct Location

    NOTE : this function is currently a STUB and returns a default

	Location at lat=0, lon=0, regarless of input.

func (address *Address) String() string
    func String : converts struct Address to a String.

    NOTE : this function is currently a STUB and returns a default

	empty string.

type Location struct {
    // contains filtered or unexported fields
}
    struct Location : represents a (latitude, longitude) point on a sphere.

func NewLocAddress(address *Address) *Location
    func NewLocAddress : is a struct Location factory method.

func NewLocLatLon(latitude float64, longitude float64) *Location
    func NewLocLatLon : is a struct Location factory method.

func (loc *Location) DiffAddr(compAddr *Address) float64
    func (loc *Location) DiffAddr(compAddr *Address) : calculates the
    distance between this Location and the input Address. The Haversine
    method is used to compute distance over the great arc between the two
    points.

func (loc *Location) DiffLoc(compLoc *Location) float64
    func (loc *Location) DiffLoc(compLoc *Location) : calculates the
    distance between this Location and the input Location. The Haversine
    method is used to compute distance over the great arc between the two
    points.

    NOTE : this formula assumes that lat and long values are in DEGREES.

func (loc *Location) Latitude() float64
    func Latitude : is a struct Location.latitude accessor

func (loc *Location) Longitude() float64
    func Longitude : is a struct Location.longitude accessor

func (loc *Location) String() string
    func String : converts struct Location to a String.

