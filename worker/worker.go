package worker

import (            
    "github.com/krustnic/runner/common"
    "github.com/krustnic/runner/job"
)

type Worker struct {
    Id int
    WorkingQueue chan map[string]interface{}    
    Log *common.LogWrap
}

func CreateWorker( id int, workingQueue chan map[string]interface{} ) Worker {    
    // Init wrapped log        
    log := &common.LogWrap{WorkerId : id}    
    log.Printf("created")    
       
    return Worker{ id, workingQueue, log }
}

func (worker *Worker) Run() {
    for {
        message := <- worker.WorkingQueue              
        
        job := job.NewJob( message, worker.Log )
        job.Start()
    }
}