package notify

import (
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/parnurzeal/gorequest"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/workqueue"
	"time"
)

type NotifyInfo struct {
	CrdName      string
	CrdNamespace string
	Name         string
	Namespace    string
	Action       string
}

type HttpNotify struct {
	logger      log.Logger
	url         string
	notifyQueue workqueue.RateLimitingInterface
}

func New(logger log.Logger, url string, queue workqueue.RateLimitingInterface) *HttpNotify {
	return &HttpNotify{logger: logger, notifyQueue: queue, url: url}
}

func (n *HttpNotify) AddNotify(crdName, crdNameSpace, name, namespace, action string) {
	n.notifyQueue.AddRateLimited(NotifyInfo{
		CrdName:      crdName,
		CrdNamespace: crdNameSpace,
		Name:         name,
		Namespace:    namespace,
		Action:       action,
	})
}
func (n *HttpNotify) Run(stopCh <-chan struct{}) {
	defer runtime.HandleCrash()
	defer n.notifyQueue.ShutDown()

	n.logger.Log("info", "starting notifier")
	go wait.Until(n.runWorker, time.Second, stopCh)
	<-stopCh
	n.logger.Log("info", "stopping notifier")
}

func (n *HttpNotify) runWorker() {
	for n.processNextWorkItem() {
	}
}

func (n *HttpNotify) processNextWorkItem() bool {
	obj, shutdown := n.notifyQueue.Get()
	if shutdown {
		return false
	}

	// wrapping block in a func to defer c.workqueue.Done
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We must call Forget if we do not want
		// this work item being re-queued. If a transient error
		// occurs, we do not call Forget. Instead the item is put back
		// on the workqueue and attempted again after a back-off
		// period.
		defer n.notifyQueue.Done(obj)

		var key NotifyInfo
		var ok bool
		if key, ok = obj.(NotifyInfo); !ok {
			n.notifyQueue.Forget(obj)
			runtime.HandleError(fmt.Errorf("expected NotifyInfo in workqueue but got %#v", obj))
			return nil
		}
		// Run the syncHandler, passing it the namespace/name string of the
		// HelmRelease resource to sync the corresponding Chart release.
		// If the sync failed, then we return while the item will get requeued
		if err := n.notifyHandler(key); err != nil {
			return err
		}
		// If no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		n.notifyQueue.Forget(obj)
		return nil
	}(obj)

	if err != nil {
		runtime.HandleError(err)
		return true
	}
	return true
}

func (n *HttpNotify) notifyHandler(info NotifyInfo) error {
	resp, _, err := gorequest.New().Post(n.url).Send(info).End()
	if len(err) > 0 {
		return err[0]
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Status Code not 200, it is : %d", resp.StatusCode))
	}
	return nil
}
