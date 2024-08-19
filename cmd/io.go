package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"sync"
	"time"
)

const (
	OPER_READ       = "read"
	OPER_WRITE      = "write"
	OPER_READ_WRITE = "read_write"
	OPER_RANDREAD   = "randread"
	OPER_RANDWRITE  = "randwrite"
)

var VALID_OPERATION = []string{OPER_READ, OPER_WRITE, OPER_READ_WRITE, OPER_RANDREAD, OPER_RANDWRITE}

func stressIO() *cli.Command {
	return &cli.Command{
		Name:  "io",
		Usage: "IO Stress",
		Description: `
The I/O stress test tool performs specified I/O operations on files or directories to simulate load and test the performance of the system.
· The "--threads"" option specifies the number of concurrent threads to use.
· The "--filepath" option specifies the path to the file or directory where operations will be performed.
· The "-operation" option specifies the type of I/O operation. The available operations are:
    - "read": Read data from files.
    - "write": Write data to files.
    - "readwrite": Perform both read and write operations.
    - "randread": Perform random read operations.
    - "randwrite": Perform random write operations.
· The "-size" option specifies the total size of data to be written to the files, in human-readable format (e.g., 100G for 100GB, 1T for 1TB).
· The "-filecount" option specifies the number of files to create and use during the test.
· The "-blocksize" option specifies the size of each block used during read/write operations, in bytes.
· The "-duration" option specifies the total duration of the test in seconds.
`,
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "threads",
				Usage:    "Number of threads",
				Aliases:  []string{"p"},
				Value:    1,
				Required: false,
			},
			&cli.StringFlag{
				Name:     "filepath",
				Usage:    "File path",
				Aliases:  []string{"f"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "operation",
				Usage:    "Operation, include read, write, readwrite, randread, randwrite",
				Aliases:  []string{"o"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "size",
				Usage:    "Size of file",
				Aliases:  []string{"s"},
				Required: false,
			},
			&cli.IntFlag{
				Name:     "filecount",
				Usage:    "Number of files to write",
				Aliases:  []string{"c"},
				Required: false,
				Value:    1,
			},
			&cli.IntFlag{
				Name:     "blocksize",
				Usage:    "Block size",
				Aliases:  []string{"b"},
				Required: false,
				Value:    KB_1,
			},
			&cli.StringFlag{
				Name:     "duration",
				Usage:    "Duration",
				Aliases:  []string{"d"},
				Required: false,
				Value:    "1m",
			},
		},
		Action: iostress,
	}
}

func isValidOperation(operation string) bool {
	for _, v := range VALID_OPERATION {
		if v == operation {
			return true
		}
	}
	return false
}

func checkOperationParam(ctx *cli.Context) bool {
	operation := ctx.String("operation")
	if !isValidOperation(operation) {
		return false
	}
	switch operation {
	case OPER_READ:
		fileInfo, err := os.Stat(ctx.String("filepath"))
		if err != nil {
			fmt.Println("File does not exist")
			return false
		}
		if fileInfo.IsDir() {
			fmt.Println("File is a directory")
			return false
		}
		return true
	case OPER_WRITE:
		_, err := os.Stat(ctx.String("filepath"))
		if err != nil {
			fmt.Println("File does not exist")
		}
		return true
	case OPER_READ_WRITE:
		return true
	case OPER_RANDREAD:
		return true
	case OPER_RANDWRITE:
		return true
	default:
		return false
	}
}

func performRead(c *cli.Context) error {
	file, err := os.Open(c.String("filepath"))
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, c.Int("blocksize"))
	duration, err := time.ParseDuration(c.String("duration"))
	if err != nil {
		return err
	}
	durSec := int64(duration.Seconds())
	start := time.Now()
	for {
		now := time.Now()
		if now.Sub(start) > time.Duration(durSec)*time.Second {
			break
		}
		_, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				if _, err := file.Seek(0, io.SeekStart); err != nil {
					break
				}
				continue
			}
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func iostress(c *cli.Context) error {
	if !checkOperationParam(c) {
		return fmt.Errorf("Invalid operation")
	}
	operation := c.String("operation")
	var wg sync.WaitGroup
	for i := 0; i < c.Int("threads"); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			switch operation {
			case OPER_READ:
				performRead(c)
			default:
				fmt.Printf("Unknown operation '%s'\n", operation)
			}
		}()
	}
	wg.Wait()
	return nil
}
