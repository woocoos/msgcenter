package service

import (
	"errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/test"
	"testing"
)

type fakeRegisterer struct {
	registeredCollectors []prometheus.Collector
}

func (r *fakeRegisterer) Register(prometheus.Collector) error {
	return nil
}

func (r *fakeRegisterer) MustRegister(c ...prometheus.Collector) {
	r.registeredCollectors = append(r.registeredCollectors, c...)
}

func (r *fakeRegisterer) Unregister(prometheus.Collector) bool {
	return false
}

func TestCoordinatorRegistersMetrics(t *testing.T) {
	fr := fakeRegisterer{}
	NewCoordinator(conf.New(conf.WithLocalPath(test.Path("testdata/alertmanager/conf.good.yml"))))

	if len(fr.registeredCollectors) == 0 {
		t.Error("expected NewCoordinator to register metrics on the given registerer")
	}
}

func TestCoordinatorNotifiesSubscribers(t *testing.T) {
	callBackCalled := false
	c := NewCoordinator(conf.New(conf.WithLocalPath(test.Path("testdata/alertmanager/conf.good.yml"))))
	c.Subscribe(func(*profile.Config) error {
		callBackCalled = true
		return nil
	})

	err := c.Reload()
	if err != nil {
		t.Fatal(err)
	}

	if !callBackCalled {
		t.Fatal("expected coordinator.Reload() to call subscribers")
	}
}

func TestCoordinatorFailReloadWhenSubscriberFails(t *testing.T) {
	errMessage := "something happened"
	c := NewCoordinator(conf.New(conf.WithLocalPath(test.Path("testdata/alertmanager/conf.good.yml"))))

	c.Subscribe(func(*profile.Config) error {
		return errors.New(errMessage)
	})

	err := c.Reload()
	if err == nil {
		t.Fatal("expected reload to throw an error")
	}

	if err.Error() != errMessage {
		t.Fatalf("expected error message %q but got %q", errMessage, err)
	}
}
