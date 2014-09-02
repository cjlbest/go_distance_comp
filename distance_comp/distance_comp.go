package main

import (
    "github.com/cjlbest/go_distance_comp/service"
)

//
// distance_comp.main() starts the distance_comp service on port 3000
//
func main() {
    service.CreateAndRun(3000)
}