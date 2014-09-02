//
// package service : encapsulates functionality for the distance computing
// service.
//
package service

import (
    "github.com/go-martini/martini"
    "fmt"
    "strconv"
    "github.com/cjlbest/go_distance_comp/geometry"
    "errors"
    "log"
    "net/http"
)

//
// struct Service : is the functional representation of a distance computation
// service. The expected REST request format is:
//
// http://localhost:3000/distance/id/<int>/lat1/<float>/lon1/<float>/lat2/<float>/lon2/<float>
//
// Future addtions (currently stubbed functionality) for the following formats will be populated:
//
// http://localhost:3000/distance/id/<int>/lat/<float>/lon/<float>/address/<string>
// http://localhost:3000/distance/id/<int>/address/<string>/address/<string>
//
type Service struct {
    *martini.ClassicMartini
}

//
// Public Methods
//

//
// func CreateAndRun : is a convenience method for starting a Service.
//
func CreateAndRun(port int64) {
    s := New  ()
    s .  Start(port)
}

//
// func New() : is a factory method for struct Service.
//
func New() *Service {

    s := Service{martini.Classic()}

    return &s
}

//
// func RunService : sets up REST parameter parsing and starts
// the microservice instance.
//
func (s *Service) Start(port int64) {

    //
    // Set GET handlers
    //

    // Check for lat1,lon1,lat2,lon2 request
    s.Get("/distance/id/:id/lat1/:lat1/lon1/:lon1/lat2/:lat2/lon2/:lon2",
            func(params martini.Params) string {
                return distPairToPair(params)
            })

    // Check for address1,address2 request
    s.Get("/distance/id/:id/address1/:address1/address2/:address2",
            func(params martini.Params) string {
                return distAddressToAddress(params)
            })

    // Check for lat,lon,address or address,lat,lon request
    s.Get("/distance/id/:id/lat/:lat/lon/:lon/address/:address",
            func(params martini.Params) string {
                return distPairToAddress(params)
            })
    s.Get("/distance/id/:id/address/:address/lat/:lat/lon/:lon",
            func(params martini.Params) string {
                return distPairToAddress(params)
            })

    //
    // Start the service
    //
    if port == 3000 {
        // Default port
        s.Run()
    } else {
        // Custom port
        portStr := ":" + strconv.FormatInt(port, 10)

        log.Fatal(http.ListenAndServe(portStr, s))
    }
}

//
// Private handlers
//

//
// func distPairToPair : GETS passed http params for id, lat1, lon1, lat2, lon2,
// and caculates the distance between the two input points on the earth.
//
func distPairToPair(params martini.Params) string {

    // Pull Link Parameter Values

    idVal, err := readIntParam  (params, "id");   if err != nil { return err.Error() }
    lat1,  err := readFloatParam(params, "lat1"); if err != nil { return err.Error() }
    lon1,  err := readFloatParam(params, "lon1"); if err != nil { return err.Error() }
    lat2,  err := readFloatParam(params, "lat2"); if err != nil { return err.Error() }
    lon2,  err := readFloatParam(params, "lon2"); if err != nil { return err.Error() }

    location1  := geometry .NewLocLatLon(lat1, lon1)
    location2  := geometry .NewLocLatLon(lat2, lon2)
    distance   := location1.DiffLoc     (location2)

    return encodeJSON(idVal, distance)
}

//
// Private Methods
//

//
// func distPairToAddress : GETS passed http params for id, lat, lon, address,
// and caculates the distance between the two input points on the earth.
// 
func distPairToAddress(params martini.Params) string {

    // Pull Link Parameter Values

    idVal,   err := readIntParam  (params, "id");      if err != nil { return err.Error() }
    lat,     err := readFloatParam(params, "lat");     if err != nil { return err.Error() }
    lon,     err := readFloatParam(params, "lon");     if err != nil { return err.Error() }
    addrStr, err := readStrParam  (params, "address"); if err != nil { return err.Error() }

    location     := geometry.NewLocLatLon(lat, lon)
    address      := geometry.NewAddress  (addrStr)
    distance     := location.DiffAddr    (address)

    return encodeJSON(idVal, distance)
}

//
// func distAddressToAddress : GETS passed http params for address1, address2
// and caculates the distance between the two input points on the earth.
// 
func distAddressToAddress(params martini.Params) string {

    // Pull Link Parameter Values

    idVal,    err := readIntParam(params, "id");       if err != nil { return err.Error() }
    addrStr1, err := readStrParam(params, "address1"); if err != nil { return err.Error() }
    addrStr2, err := readStrParam(params, "address2"); if err != nil { return err.Error() }

    address1      := geometry.NewAddress (addrStr1)
    address2      := geometry.NewAddress (addrStr2)
    distance      := address1.DiffAddr   (address2)

    return encodeJSON(idVal, distance)
}

//
// func readIntParam : reads the REST parameter with the input name and
// parses the result string for an integer value. An error is returned
// if if the read fails or if the conversion from string to int fails.
//
func readIntParam(params martini.Params, name string) (int, error) {

    strVal      := params[name]
    intVal, err := strconv.Atoi(strVal)

    if err != nil {
        // Push error info on.
        // err = errors.New(err.Error() + " : " + name + " parameter is not an int")
        err = errors.New("404 page not found")
    }

    return intVal, err
}

//
// func readFloatParam : reads the REST parameter with the input name and
// parses the result string for a float64 value. An error is returned
// if if the read fails or if the conversion from string to float64 fails.
//
func readFloatParam(params martini.Params, name string) (float64, error) {

    strVal        := params[name]
    floatVal, err := strconv.ParseFloat(strVal, 64)

    if err != nil {
        // Push error info on.
        // err = errors.New(err.Error() + " : " + name + " parameter is not a float")
        err = errors.New("404 page not found")
    }

    return floatVal, err
}

//
// func readFloatParam : reads the REST parameter with the input name. An error
// is returned if the read fails.
//
func readStrParam(params martini.Params, name string) (string, error) {
    strVal := params[name]

    if len(strVal) == 0 {
        err := errors.New("404 page not found")

        return strVal, err
    }

    return strVal, nil
}

//
// func encodeJSON : formats a JSON return value for input id and distance data.
//
func encodeJSON(idVal int, distance float64) string {
    jsonStr := fmt.Sprintf("{\"id\": %d, \"distance\": %f}", idVal, distance)

    return jsonStr
}
