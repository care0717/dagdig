package openapi

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func Parse(input io.Reader) (Job, error) {
	body, err := ioutil.ReadAll(input)
	if err != nil {
		return Job{}, err
	}
	lines := strings.Split(string(body), "\n")
	job := Job{}
	for i := 0; i < len(lines); i++ {
		switch lines[i] {
		case "[JobID]":
			i++
			id, err := strconv.Atoi(lines[i])
			if err != nil {
				return Job{}, err
			}
			job.Id = int32(id)
		case "[Created]":
			i++
			job.Created = lines[i]
		case "[Priority]":
			i++
			job.Priority = lines[i]
		case "[Tasks]":
			i++
			for ; i < len(lines); i++ {
				if lines[i] == "" {
					continue
				}
				point, err := strconv.Atoi(lines[i])
				if err != nil {
					return Job{}, err
				}
				job.Tasks = append(job.Tasks, int32(point))
			}
		}
	}
	return job, nil
}
