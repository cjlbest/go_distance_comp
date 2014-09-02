//
// package geometry : provides functionality for location operations.
//
package geometry

import (
    "testing"
    "os"
    "log"
    "fmt"
)

const pkgTestFileLoc string  = "/src/github.com/cjlbest/go_distance_comp/"

//
// func BenchmarkGeometry : benchmarks the distHaversine calculation.
//
func BenchmarkGeometry(bench *testing.B){
    for n := 0; n < bench.N; n++ {
        // use first values from test file.
        distHaversine(39.768434, -104.901872, 44.779373, -63.673489)
    }
}

//
// funct TestGeometry : reads and parses input csv gis data test file and run
// syncronouss service for each entry pair.  Input file data layout must be :
//
// ID(int),lat1(float),lon1(float),lat2(float),lon2(float),distance(float)
//
// NOTE : GOPATH must be set correctly.
//
func TestGeometry(test *testing.T) {

    goPath := os.Getenv("GOPATH")

    if len(goPath) > 0 {

        goPath       = goPath + pkgTestFileLoc
        inFilePath  := goPath + "baselines.csv"
        outFilePath := goPath + "geometry_test.csv"

        fmt.Printf("\ndistance_comp test : %s\n\n", inFilePath)

        err := CalcAllWithCompOp(inFilePath, outFilePath, test, computeVals)

        if err != nil {
            log.Fatal (err)
            fmt.Printf("%s", err.Error())
        }

    } else {

        fmt.Printf("\nERROR : envar $GOPATH must be correctly set\n")

    }
}

//
// func computeVals : computes a distance value from the input location values.
//
func computeVals(idVal int,
                 lat1  float64,
                 lon1  float64,
                 lat2  float64,
                 lon2  float64,
                 test  *testing.T) (float64, error) {

     //
     // Compute distances
     //

     loc1 := NewLocLatLon(lat1, lon1)
     loc2 := NewLocLatLon(lat2, lon2)
     dist := loc1.DiffLoc(loc2)

     return dist, nil
}