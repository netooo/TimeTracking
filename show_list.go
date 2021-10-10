package main

import (
	"fmt"
	tracking "github.com/netooo/TimeTracking/lib"
	"github.com/urfave/cli"
)

func List(c *cli.Context) error {
	taskList, err := tracking.TaskRead()
	if err != nil {
		return err
	}

	for _, task := range taskList {
		fmt.Printf("%-9d", task.ID)
		fmt.Printf("%-4d", task.ContentLine)
		fmt.Println(task.Name)
	}

	return nil
}
