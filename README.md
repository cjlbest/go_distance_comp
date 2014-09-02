go_distance_comp
================

Microservice for calculating the distance between two lon/lat points on the globe.

This package requires golang to be installed and the $GOPATH envar to be set.

To install package: 
go get github.com/cjlbest/go_distance_comp/distance_comp

Building the package will create a binary executable, "distance_comp", to run the service.

To build a binary in the local package :
go build github.com/cjlbest/go_distance_comp/distance_comp

To build and install a binary in $GOPATH/bin :
go install github.com/cjlbest/go_distance_comp/distance_comp
