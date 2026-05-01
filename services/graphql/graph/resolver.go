package graph

import "tip2_pr12/services/graphql/internal/client/tasksclient"

type Resolver struct {
	TasksClient *tasksclient.Client
}
