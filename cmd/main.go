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
	"github.com/matesu777/Mattix/internal/config"
	mattixhttp "github.com/matesu777/Mattix/internal/http"

	"github.com/matesu777/Mattix/internal/collector"
)

const version = "0.1.5"

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error load configs: %s", err)
	}

	cpuColl := cpu.NewCollector()
	diskColl := disk.NewCollector(&cfg.Disk)
	hostnameColl := hostname.NewCollector()
	memoryColl := memory.NewCollector()
	networkColl := network.NewCollector(&cfg.Network)
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

	manager := collector.NewManager(fastComponents, slowComponents, &cfg.Interval)
	ctx := context.Background()

	manager.Start(ctx)
	handler := mattixhttp.ServerNew(&cfg.Server, manager)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /metrics", handler.GetMetrics)
	hostname, err := hostnameColl.Collect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	printBanner(version, hostname, cfg)

	port := fmt.Sprintf(":%d", cfg.Server.Port)

	log.Fatal(http.ListenAndServe(port, mattixhttp.Cors(mux)))
}

func printBanner(version string, hostname any, cfg *config.Config) {
	fmt.Printf(
		`Mattix Agent %s
--------------------------
%-12s %s
%-12s :%d
--------------------------
		`,
		version,
		"Hostname", hostname,
		"Port", cfg.Server.Port,
	)
}
