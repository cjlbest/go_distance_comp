PACKAGE DOCUMENTATION

package service
    import "distance_comp/service"

    package service : encapsulates functionality for the distance computing
    service.

FUNCTIONS

func CreateAndRun(port int64)
    func CreateAndRun : is a convenience method for starting a Service.

TYPES

type Service struct {
    *martini.ClassicMartini
}
    struct Service : is the functional representation of a distance
    computation service. The expected REST request format is:

    http://localhost:3000/distance/id/<int>/lat1/<float>/lon1/<float>/lat2/<float>/lon2/<float>

    Future addtions (currently stubbed functionality) for the following
    formats will be populated:

    http://localhost:3000/distance/id/<int>/lat/<float>/lon/<float>/address/<string>
    http://localhost:3000/distance/id/<int>/address/<string>/address/<string>

func New() *Service
    func New() : is a factory method for struct Service.

func (s *Service) Start(port int64)
    func RunService : sets up REST parameter parsing and starts the
    microservice instance.

