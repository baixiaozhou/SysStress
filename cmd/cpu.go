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
CPUStress is a robust tool designed to simulate heavy workloads on specific CPU cores. It provides three key parameters for customized testing:
1. cpu-number: Specifies the number of CPU cores to be stressed. This parameter allows you to target a specific number of cores, enabling fine-tuned performance testing or benchmarking on multi-core systems.
2. duration: Defines the duration of the stress test. This parameter lets you determine how long the specified cores should be stressed, providing flexibility for short bursts or prolonged stress tests to evaluate system stability under load.
3. force: force is an optional parameter, primarily used to allow requesting a number of cores that exceeds the system’s core count.
Examples:
# To run stress test with 10 threads for 10 minutes
sysstress cpu --cpu-number 10 --duration 10m
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
			&cli.BoolFlag{
				Name:        "force",
				Aliases:     []string{"f"},
				Usage:       "Force stress test without confirmation",
				Required:    false,
				DefaultText: "false",
			},
		},
		Action: cpu,
	}
}

func cpu(c *cli.Context) error {
	cpuNum := c.Int("cpu-number")
	if cpuNum <= 0 || (!c.Bool("force") && cpuNum > runtime.NumCPU()) {
		return fmt.Errorf("cpu number must be between 0 and %d, "+
			"if you want to exceed the number of CPU cores, please use the --force option", runtime.NumCPU())
	}
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
