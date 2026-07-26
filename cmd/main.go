package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/matesu777/Mattix/internal/components/cpu"
	"github.com/matesu777/Mattix/internal/components/disk"
	"github.com/matesu777/Mattix/internal/components/hostname"
	"github.com/matesu777/Mattix/internal/components/memory"
	"github.com/matesu777/Mattix/internal/components/network"
	"github.com/matesu777/Mattix/internal/components/temperature"
	"github.com/matesu777/Mattix/internal/components/uptime"
	mattixhttp "github.com/matesu777/Mattix/internal/http"

	"github.com/matesu777/Mattix/internal/collector"
)

func main() {
	cpuColl := cpu.NewCollector()
	diskColl := disk.NewCollector()
	hostnameColl := hostname.NewCollector()
	memoryColl := memory.NewCollector()
	networkColl, _ := network.NewCollector()
	temperatureColl := temperature.NewCollector()
	uptimeColl := uptime.NewCollector()

	fastComponents := []collector.Component{
		cpuColl,
		networkColl,
		uptimeColl,
	}

	slowComponents := []collector.Component{
		diskColl,
		memoryColl,
		temperatureColl,
		hostnameColl,
	}

	manager := collector.NewManager(fastComponents, slowComponents)
	ctx := context.Background()

	manager.Start(ctx)
	handler := mattixhttp.ServerNew(manager)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /metrics", handler.GetMetrics)
	hostname, _ := hostnameColl.Collect(ctx)

	fmt.Printf("Mattix agent v0.1.0 \n\nhostname: %s \nlistening on: 8080\n", hostname)

	log.Fatal(http.ListenAndServe(":8080", mattixhttp.Cors(mux)))
}
