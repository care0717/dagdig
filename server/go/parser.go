package openapi

import (
	"bufio"
	"io"
	"strconv"
)

func Parse(input io.Reader) (Job, error) {
	scanner := bufio.NewScanner(input)
	job := Job{}
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "[JobID]":
			scanner.Scan()
			line = scanner.Text()
			id, err := strconv.Atoi(line)
			if err != nil {
				return Job{}, err
			}
			job.Id = int32(id)
		case "[Created]":
			scanner.Scan()
			line = scanner.Text()
			job.Created = line
		case "[Priority]":
			scanner.Scan()
			line = scanner.Text()
			job.Priority = line
		case "[Tasks]":
			for scanner.Scan() {
				line = scanner.Text()
				if line == "" {
					continue
				}
				point, err := strconv.Atoi(line)
				if err != nil {
					return Job{}, err
				}
				job.Tasks = append(job.Tasks, int32(point))
			}
		}
	}
	return job, nil
}
