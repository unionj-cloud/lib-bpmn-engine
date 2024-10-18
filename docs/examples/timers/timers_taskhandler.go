package main

import (
	"github.com/unionj-cloud/lib-bpmn-engine/pkg/bpmn_engine"
)

func registerDummyTaskHandlers(bpmnEngine *bpmn_engine.BpmnEngineState) {
	var justCompleteHandler = func(job bpmn_engine.ActivatedJob) {
		job.Complete()
	}
	bpmnEngine.NewTaskHandler().Id("ask").Handler(justCompleteHandler)
	bpmnEngine.NewTaskHandler().Id("win").Handler(justCompleteHandler)
	bpmnEngine.NewTaskHandler().Id("lose").Handler(justCompleteHandler)
}
