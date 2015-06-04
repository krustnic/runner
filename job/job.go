package job

import (        
    "github.com/krustnic/runner/common"
)

type Job struct {
    Message map[string]interface{}
    Log *common.LogWrap
}

var log *common.LogWrap

func NewJob(message map[string]interface{}, log *common.LogWrap) Job {
    return Job{ message, log }
}

func (job *Job) Start() {
    job.Log.Printf("Job. Processing message: %v", job.Message)        
}