package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	clientset "github.com/tilt-dev/tilt-trigger-on-ready/pkg/clientset/versioned"
	"github.com/tilt-dev/tilt-trigger-on-ready/pkg/config"
	"github.com/tilt-dev/tilt/pkg/apis/core/v1alpha1"
	"io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"os/exec"
	"time"
)

type Trigger struct {
	Target string     `json:"target"`
	Command []string  `json:"command"`
}

type Config []Trigger

func readConfig(r io.Reader) (Config, error) {
	var ret Config
	d := json.NewDecoder(r)
	d.DisallowUnknownFields()
	err := d.Decode(&ret)
	return ret, err
}

func watchSessions() (chan *v1alpha1.Session, context.CancelFunc, error) {
	tiltAPIConfig, err := config.NewConfig()
	if err != nil {
		return nil, nil, err
	}

	cli := clientset.NewForConfigOrDie(tiltAPIConfig)

	ctx, cancel := context.WithCancel(context.Background())

	w, err := cli.TiltV1alpha1().Sessions().Watch(ctx, v1.ListOptions{})
	if err != nil {
		return nil, cancel, err
	}

	ch := make(chan *v1alpha1.Session)
	go func() {
		for {
			e, ok := <-w.ResultChan()
			if !ok {
				close(ch)
				return
			}
			ch <- e.Object.(*v1alpha1.Session)
		}
	}()

	return ch, cancel, nil
}

// calculates the time at which the given target became ready, or zero if it's not ready
func targetReadyTime(t v1alpha1.Target) time.Time {
	switch t.Type {
	case v1alpha1.TargetTypeServer:
		if t.State.Active != nil && t.State.Active.Ready {
			return t.State.Active.StartTime.Time
		} else {
			return time.Time{}
		}
	case v1alpha1.TargetTypeJob:
		if t.State.Terminated != nil && t.State.Terminated.Error == "" {
			return t.State.Terminated.FinishTime.Time
		} else {
			return time.Time{}
		}
	default:
		_, _ = fmt.Fprintf(os.Stderr, "target %s has unknown type %s\n", t.Name, t.Type)
		return time.Time{}
	}
}

// returns a list of any targets that have completed between last and cur
func newlyReadyTargets(last, cur *v1alpha1.Session) []string {
	lastReadyTimes := make(map[string]time.Time)
	for _, t := range last.Status.Targets {
		lastReadyTimes[t.Name] = targetReadyTime(t)
	}

	var ret []string
	for _, t := range cur.Status.Targets {
		if targetReadyTime(t).After(lastReadyTimes[t.Name]) {
			ret = append(ret, t.Name)
		}
	}
	return ret
}

// turns a chan of Session updates into a chan of newly ready targets
func completedTargets(ch chan *v1alpha1.Session) chan string {
	ret := make(chan string)
	var lastSession *v1alpha1.Session
	go func() {
		for {
			sess, ok := <-ch
			if !ok {
				close(ret)
				return
			}
			if lastSession != nil {
				for _, t := range newlyReadyTargets(lastSession, sess) {
					ret <- t
				}
			}
			lastSession = sess
		}
	}()
	return ret
}

func usage() {
	_, _ = fmt.Fprint(os.Stderr, `Usage: echo "$MYCONFIG" | tilt-session-trigger
Where $MYCONFIG is JSON of the form:
[
  {"target": "myresource:update", "command": ["tilt", "trigger", "myotherresource"]}
]

You can also specify target: "*" to run on all targets. For these triggers, the $TILT_TARGET_NAME variable will be set.

To see target names of a running tilt instance, run: tilt get session -ojsonpath='{.items[*].status.targets[*].name}'
`)
}

func maybeTrigger(cfg Config, target string) {
	for _, trigger := range cfg {
		if trigger.Target == target || trigger.Target == "*" {
			fmt.Printf("target %s completed, running %v\n", target, trigger.Command)
			cmd := exec.Command(trigger.Command[0], trigger.Command[1:]...)
			if trigger.Target == "*" {
				cmd.Env = append(cmd.Env, fmt.Sprintf("TILT_TARGET_NAME=%s", target))
			}

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Printf("command %v failed: %v\n", trigger.Command, err)
			}
		}
	}
}

func run() error {
	cfg, err := readConfig(os.Stdin)
	if err != nil {
		usage()
		return errors.Wrap(err, "error reading config from stdin")
	}

	sessions, cancel, err := watchSessions()
	if err != nil {
		return err
	}
	defer cancel()

	for t := range completedTargets(sessions) {
		maybeTrigger(cfg, t)
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
