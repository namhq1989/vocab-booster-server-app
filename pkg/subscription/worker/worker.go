package worker

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/namhq1989/vocab-booster-server-app/internal/queue"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type (
	Handlers interface {
		DowngradeUserSubscription(ctx *appcontext.AppContext, payload domain.QueueDowngradeUserSubscriptionPayload) error
	}
	Cronjob interface {
		ScanExpiredUserSubscription(ctx *appcontext.AppContext, _ domain.QueueScanExpiredUserSubscription) error
	}
	Instance interface {
		Handlers
		Cronjob
	}

	workerHandlers struct {
		DowngradeUserSubscriptionHandler
	}
	workerCronjob struct {
		ScanExpiredUserSubscriptionHandler
	}
	Worker struct {
		queue queue.Operations
		workerHandlers
		workerCronjob
	}
)

var _ Instance = (*Worker)(nil)

func New(
	queue queue.Operations,
	userSubscriptionRepository domain.UserSubscriptionRepository,
	queueRepository domain.QueueRepository,
) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			DowngradeUserSubscriptionHandler: NewDowngradeUserSubscriptionHandler(userSubscriptionRepository),
		},
		workerCronjob: workerCronjob{
			ScanExpiredUserSubscriptionHandler: NewScanExpiredUserSubscriptionHandler(userSubscriptionRepository, queueRepository),
		},
	}
}

func (w Worker) Start() {
	w.addCronjob()

	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.DowngradeUserSubscription), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueDowngradeUserSubscriptionPayload](bgCtx, t, queue.ParsePayload[domain.QueueDowngradeUserSubscriptionPayload], w.DowngradeUserSubscription)
	})
}

type cronjobData struct {
	Task       string      `json:"task"`
	CronSpec   string      `json:"cronSpec"`
	Payload    interface{} `json:"payload"`
	RetryTimes int         `json:"retryTimes"`
}

func (w Worker) addCronjob() {
	var (
		ctx  = appcontext.NewWorker(context.Background())
		jobs = []cronjobData{
			{
				Task:       w.queue.GenerateTypename(queue.TypeNames.ScanExpiredUserSubscription),
				CronSpec:   "1 0 * * *", // at 00:01am every day
				Payload:    domain.QueueScanExpiredUserSubscription{},
				RetryTimes: 3,
			},
		}
	)

	for _, job := range jobs {
		entryID, err := w.queue.ScheduleTask(job.Task, job.Payload, job.CronSpec, job.RetryTimes)
		if err != nil {
			ctx.Logger().Error("error when initializing cronjob", err, appcontext.Fields{"job": job})
			panic(err)
		}

		ctx.Logger().Info(fmt.Sprintf("[cronjob] cronjob '%s' initialize successfully with cronSpec '%s' and retryTimes '%d'", job.Task, job.CronSpec, job.RetryTimes), appcontext.Fields{
			"entryId": entryID,
		})
	}
}
