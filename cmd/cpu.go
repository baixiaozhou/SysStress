package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"runtime"
	"sync"
	"time"
)

func stressCpu() *cli.Command {
	return &cli.Command{
		Name:  "cpu",
		Usage: "CPU Stress",
		Description: `
CPUStress is a robust tool designed to simulate heavy workloads on specific CPU cores. It provides two key parameters for customized testing:
· cpu-number: Specifies the number of CPU cores to be stressed. This parameter allows you to target a specific number of cores, enabling fine-tuned performance testing or benchmarking on multi-core systems.
· duration: Defines the duration of the stress test. This parameter lets you determine how long the specified cores should be stressed, providing flexibility for short bursts or prolonged stress tests to evaluate system stability under load.

Examples:
# To run stress test with 10 threads for 10 minutes
sysstrain cpu --cpu-number 10 --duration 10m
`,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "cpu-number",
				Aliases:  []string{"n"},
				Usage:    "CPU number",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "duration",
				Aliases:  []string{"d"},
				Usage:    "Duration in seconds",
				Required: false,
			},
		},
		Action: cpu,
	}
}

func cpu(c *cli.Context) error {
	cpuNum := c.Int("cpu-number")
	if cpuNum <= 0 || cpuNum > runtime.NumCPU() {
		return fmt.Errorf("cpu number must be between 0 and %d, %d", runtime.NumCPU(), cpuNum)
	}
	fmt.Println("duration:", c.Duration("duration"))
	duration, err := time.ParseDuration(c.String("duration"))
	if err != nil {
		return err
	}
	durSec := int64(duration.Seconds())
	startTime := time.Now()
	// Set Max cpu cores
	runtime.GOMAXPROCS(cpuNum)
	var wg sync.WaitGroup
	wg.Add(cpuNum)

	for i := 0; i < cpuNum; i++ {
		go burnCpu(&wg, startTime, durSec)
	}
	wg.Wait()
	return nil
}

func burnCpu(wg *sync.WaitGroup, start time.Time, durSec int64) {
	defer wg.Done()
	for {
		_ = 1 * 1
		now := time.Now()
		if now.Sub(start) > time.Duration(durSec)*time.Second {
			break
		}
	}
}
