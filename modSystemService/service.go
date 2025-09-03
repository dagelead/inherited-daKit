package modSystemService

import (
	"fmt"
	"os"
	"time"

	"github.com/kardianos/service"
)

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	p.exit = make(chan struct{})
	go p.run()
	return nil
}

func (p *program) run() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("Service tick...")
		case <-p.exit:
			fmt.Println("Service shutting down gracefully...")
			return
		}
	}
}

func (p *program) Stop(s service.Service) error {
	close(p.exit)
	return nil
}

func serviceStart() {
	svcConfig := &service.Config{
		Name:        "daKit",
		DisplayName: "daKit",
		Description: "DaGe Personal System Manager Kit",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println("Failed to create service:", err)
		return
	}

	// Run service (supports install/start/stop commands too)
	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			fmt.Printf("Valid actions: install, uninstall, start, stop\nError: %v\n", err)
		}
		return
	}

	err = s.Run()
	if err != nil {
		fmt.Println("Service failed:", err)
	}
}
