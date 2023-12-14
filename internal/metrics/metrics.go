package metrics

import "github.com/VictoriaMetrics/metrics"

var (
	MetricCreateTaskCounter = metrics.NewCounter(`todo_create_task_counter`)
	MetricFinishTaskCounter = metrics.NewCounter(`todo_finish_task_counter`)
	MetricDeleteTaskCounter = metrics.NewCounter(`todo_delete_task_counter`)
	MetricErrorCounter      = metrics.NewCounter(`todo_error_counter`)
)
