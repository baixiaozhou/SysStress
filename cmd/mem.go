package cmd

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/mem"
	"github.com/urfave/cli/v2"
	"time"
)

const (
	KB_1   = 1 << 10
	FORMAT = "2006-01-02 15:04:05"
)

func stressMem() *cli.Command {
	return &cli.Command{
		Name:  "memory",
		Usage: "memory stress",
		Description: `Memory stress testing tool primarily used to consume system memory. It provides three parameters:
1.	size: The amount of memory to allocate. If the requested size exceeds the total system memory, the operation is denied. If it exceeds the available memory, the operation is also denied. However, if the --force parameter is set to true, the restriction can be lifted.
2.	duration: The duration for which the memory allocation should run.
3.	force: An optional parameter that references the size parameter, allowing the request to exceed system memory limits as described in point 1.”
`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "size",
				Aliases:  []string{"s"},
				Usage:    "memory size",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "duration",
				Aliases:  []string{"d"},
				Usage:    "memory duration",
				Required: true,
			},
			&cli.BoolFlag{
				Name:        "force",
				Aliases:     []string{"f"},
				Usage:       "force overwrite",
				Required:    false,
				DefaultText: "false",
			},
		},
		Action: memory,
	}
}

func memory(c *cli.Context) error {
	duration, err := time.ParseDuration(c.String("duration"))
	if err != nil {
		return err
	}
	force := c.Bool("force")
	size, err := humanize.ParseBytes(c.String("size"))
	if err != nil {
		return fmt.Errorf("memory stress error: %s", err)
	}
	memInfo, _ := mem.VirtualMemory()
	if size >= memInfo.Total {
		return fmt.Errorf("memory stress error: memory stress is more than total size")
	}
	if !force && size >= memInfo.Available {
		return fmt.Errorf("memory stress error: memory stress is more than avaible memory, you can use --force true")
	}
	allocSize := make([]byte, size)

	fmt.Println("Start to alloc memory size: ", size, " duration: ", duration, " start time: ", time.Now().Format(FORMAT))
	for i := range allocSize {
		// It has no actual purpose, just to prevent optimization.
		allocSize[i] = byte(i % KB_1)
	}
	time.Sleep(time.Duration(int64(duration.Seconds())) * time.Second)
	fmt.Println("memory stress end", time.Now().Format(FORMAT))
	return nil
}
