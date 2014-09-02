package geometry

import (
    "testing"
    "os"
    "fmt"
    "log"
    "math"
    "bufio"
    "io"
    "strings"
    "strconv"
)

//
// This file provides testing functionality that is used by both the geometry
// and service packages, so the methods must be public.
//

const distErrorTol float64 = 10e-4

//
// func CalcAllWithCompOp : reads input baseline data file, calculates distance
// using the input compOp function, compares the calculate value with values in
// the baseline file, and writes any error to the output file.
//
// NOTE : this function is used by service_test and geometry_test.
//
func CalcAllWithCompOp(inFilePath  string,
                       outFilePath string,
                       test        *testing.T,
                       compOp func(int,float64,float64,float64,float64,*testing.T)(float64,error)) error {

    inFile, err := os.Open(inFilePath);

    if err != nil {
        test.Fatal(err.Error())
    }

    defer inFile.Close()

    reader   := bufio.NewReader(inFile)
    eof      := false
    lineNum  := 0
    numFails := 0

    // Incase we run into errors that we want to write to file
    // for an automated followup
    var outFile *os   .File   = nil
    var writer  *bufio.Writer = nil

    for !eof {

        //
        // Perform test operation for each line in the
        // input file, ignoring the header line.
        //

        line, err := reader.ReadString('\n')

        if err == io.EOF {
            err = nil // io.EOF isn't really an error
            eof = true // this will end the loop at the next iteration
        } else if err != nil {
            return err
        }

        if lineNum = lineNum+1; lineNum > 1 {

            //
            // Compute distances and write to csv results
            //

            line  = line[0:len(line)-2]         // Peel the carraige return off
            vals := strings.Split(line, ",")    // Split line by commas

            //
            // Parse values
            //

            idVal, err := strconv.Atoi(vals[0])

            if err != nil {
                test.Error(err.Error())
                continue
            }

            lat1, err := strconv.ParseFloat(vals[1], 64)

            if err != nil {
                test.Error(err.Error())
                continue
            }

            lon1, err := strconv.ParseFloat(vals[2], 64)

            if err != nil {
                test.Error(err.Error())
                continue
            }

            lat2, err := strconv.ParseFloat(vals[3], 64)

            if err != nil {
                test.Error(err.Error())
                continue
            }

            lon2, err := strconv.ParseFloat(vals[4], 64)

            if err != nil {
                test.Error(err.Error())
                continue
            }

            actualDist, err := strconv.ParseFloat(vals[5], 64)

            if err != nil {
                test.Error(err.Error())
                continue
            }

            //
            // Request distance computation from service and
            // compare response to baseline file values.
            //

            // responseID, dist, err := requestVals(idVal, lat1, lon1, lat2, lon2, test)
            dist, err := compOp(idVal, lat1, lon1, lat2, lon2, test)

            if err != nil {
                test.Fatal(err.Error())
            }

            // if responseID != idVal {
            //     test.Errorf("Expected ID : %d Received ID : %d", idVal, responseID)
            // }

            distErr := math.Abs(dist - actualDist)

            if distErr > distErrorTol {
                if outFile == nil {
                    // Create/Open outfile and write header.
                    outFile, err = os   .Create   (outFilePath); if err != nil { return err }
                    writer       = bufio.NewWriter(outFile)
                    errorTitle   := "Distance computation errors for threashold"
                    errorColumns := "ID, Location1, Location2, Distance, Computed, Error"
                    errorHeader  := fmt  .Sprintf  ("%s %e\n\n%s\n", errorTitle, distErrorTol, errorColumns)

                    log.SetOutput(writer)

                    // Write error file header.
                    writer.WriteString(errorHeader)

                    defer outFile.Close()
                }

                numFails   = numFails + 1
                resultStr := fmt.Sprintf("%d, (%f,%f), (%f,%f), %f, %f, %f\n",
                                         idVal, lat1, lon1, lat2, lon2,
                                         actualDist, dist, distErr)
                errorStr  := fmt.Sprintf("%d, Actual:%fkm, Computed:%fkm, Diff:%fkm\n",
                                         idVal, actualDist, dist, distErr)

                writer.WriteString(resultStr) // Write to csv file for followup
                test  .Error      (errorStr)
            }
        }
    }

    if numFails > 0 {
        numTested  := lineNum - 1
        errorStr   := fmt.Sprintf("%d failures out of %d tested with %e tolerance",
                                  numFails, numTested, distErrorTol)

        test.Error(errorStr)
    }

    return nil
}
