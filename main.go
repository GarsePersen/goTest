package main
import(	
    "jobs/handlers"
    "net/http"
	"github.com/go-zoo/bone"
)

func main(){
     
     router := bone.New()

     router.Get("/api/jobs",http.HandlerFunc(handlers.GetJobs))
     //run it on port 8080
     http.ListenAndServe("0.0.0.0:8080", router)
 
}