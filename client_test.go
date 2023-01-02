package keepaapi

import (
	"testing"
	"log"
	"miromiro/keepaapi/api"
)


func Test(t *testing.T) {
	Client := api.New("YOUR_API_KEY")
	graph,err := Client.GetGraph("ASIN", api.DomainUS)
	if err != nil {
		log.Println(err)
	}
	log.Println(graph)
}