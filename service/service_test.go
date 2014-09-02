package service

import (
	"testing"
    "net/http"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "log"
    "os"
    "github.com/cjlbest/go_distance_comp/geometry"
)

const pkgTestFileLoc string  = "/src/github.com/cjlbest/go_distance_comp/"
const server         string  = "localhost"

//
// funct BenchmarkService : is the benchmark method for distance_comp/service.
//
// NOTE : A running Service is necessary, so distance_comp must be built and
//       running dring this benchmark test. I.E. build and run distance_comp
//       then run this benchmark test in a different terminal.
//
func BenchmarkService(bench *testing.B) {

    for n := 0; n < bench.N; n++ {

        // Pass GET request to Service - use first values from test file.

        request       := "http://" +
                          server   +
                         ":3000/distance/id/1/lat1/39.768434/lon1/-104.901872/lat2/44.779373/lon2/-63.673489"
        response, err := http.Get(request)

        if err != nil {
            bench.Error(err.Error())
            continue
        }

        defer response.Body.Close()

        // contents, err := ioutil.ReadAll(response.Body)

        if err != nil {
            bench.Error(err.Error())
        }

        // dist, err := parseJSON(contents)
        // // parseJSON(contents)
        // //
        // log.Printf("%s : %f\n", string(contents), dist)
    }
}

//
// func TestService : is the testing method for distance_comp/service.
//
// NOTE : A running Service is necessary, so distance_comp must be built and
//       running during this validity test. I.E. build and run distance_comp
//       then run this benchmark test in a different terminal.
//
func TestService(test *testing.T) {

    goPath := os.Getenv("GOPATH")

    if len(goPath) > 0 {

        goPath       = goPath + pkgTestFileLoc
        inFilePath  := goPath + "baselines.csv"
        outFilePath := goPath + "server_test.csv"

        fmt.Printf("\ndistance_comp test : %s\n\n", inFilePath)

        err := geometry.CalcAllWithCompOp(inFilePath, outFilePath, test, requestVals)

        if err != nil {
            log.Fatal (err)
            fmt.Printf("%s", err.Error())
        }

    } else {

        fmt.Printf("\nERROR : envar $GOPATH must be correctly set\n")

    }
}

//
// func requestVals : generates a request to the Service from the input id
// and location values.  The response JSON is parsed and the ID and distance
// values are returned.
//
func requestVals(idVal int,
                 lat1  float64,
                 lon1  float64,
                 lat2  float64,
                 lon2  float64,
                 test  *testing.T) (float64, error) {

    request       := fmt.Sprintf("http://"+server+":3000/distance/id/%d/lat1/%f/lon1/%f/lat2/%f/lon2/%f",
                                  idVal, lat1, lon1, lat2, lon2)
    response, err := http.Get(request)

    if err != nil {
        test.Error(err.Error())
    }

    defer response.Body.Close()

    contents, err := ioutil.ReadAll(response.Body)

    if err != nil {
        fmt.Printf("%s", err)
        test.Error(err.Error())
    }

    responseID, dist, err := parseJSON(contents, test)

    if responseID != idVal {
        test.Errorf("Expected ID : %d Received ID : %d", idVal, responseID)
    }

    return dist, err
}

//
// func parseJSON parses the input distance comp json response and returns the
// id and distance values. The expected format is:
//
// { id : (int), distance : (float) }
//
func parseJSON(jsonStr []byte, test *testing.T) (int, float64, error) {
    // Should Be : { id : (int), distance : (float) }

    var distSet map[string]interface{}

    err := json.Unmarshal(jsonStr, &distSet)

    if err != nil {
        test.Error(err)
    }

    idFlt := distSet["id"].(float64)
    dist  := distSet["distance"].(float64)
    idInt := int(idFlt)

    return idInt, dist, err
}
