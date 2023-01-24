package godo

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

type Resolver struct{}

func (r *Resolver) Task(ctx context.Context, args struct {
	ID graphql.ID
}) (*taskResolver, error) {
	for _, task := range tasks {
		if task.ID == args.ID {
			return &taskResolver{&task}, nil
		}
	}

	return nil, fmt.Errorf("task with id %q not found", args.ID)
}

func (r *Resolver) Tasks(ctx context.Context) []*taskResolver {
	var taskList []*taskResolver

	for _, task := range tasks {
		taskList = append(taskList, &taskResolver{&task})
	}
	return taskList
}

type taskResolver struct {
	t *Task
}

func (r *taskResolver) ID() graphql.ID {
	return r.t.ID
}

func (r *taskResolver) Name() *string {
	name := r.t.Name
	if name == "" {
		return nil
	}
	return &name
}

func (r *taskResolver) Description() *string {
	desc := r.t.Description
	if desc == "" {
		return nil
	}
	return &desc
}

func (r *taskResolver) StartDate() graphql.Time {
	return r.t.StartDate
}

func (r *taskResolver) EndDate() graphql.Time {
	return r.t.EndDate
}

func (r *taskResolver) Done() bool {
	return r.t.Done
}
