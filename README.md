# tilt-trigger-on-ready

A Tilt extension that allows you to specify a command any time a given target becomes ready.

# Examples

```
load('path/to/tilt-trigger-on-ready/Tiltfile', 'trigger_on_ready')

# kick off the frontend test resource when the frontend server finishes deploying
trigger_on_ready('frontend:serve', ['tilt', 'trigger', 'frontend-tests'])

# pop up a notification when an update finishes
# (on OS X, after `brew install terminal-notifier`)
trigger_on_ready('*', ['sh', '-c', 'terminal-notifier -message "tilt: $TILT_TARGET_NAME is ready"'])
