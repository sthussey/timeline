![Testing Badge](https://github.com/sthussey/timeline/actions/workflows/test.yaml/badge.svg) ![Linting Badge](https://github.com/sthussey/timeline/actions/workflows/linter.yaml/badge.svg)
# timeline

Library for executing a serialized event sequence

## Summary

This is a simple library that can generate a described
sequence of events to allow for testing of code response.

The library comes packaged with a few event types and
the user can map additional actions and synchronizers
for timeline to execute.

## Sequence

A timeline sequence is composed of events that
are either synchornizers or actions. Synchronizers indicate
timeline should synchornize the event flow with something such as
an OS signal or a timer expiration. Actions simply
call functions in your code base.

## Example

```yaml
  timeline:
    - sync: SignalRecv
      inputs:
        signal: SIGUSR1
    - action: NetworkBlock
      block: true
      inputs:
        dest: 10.0.0.1/32
    - sync: TimerWait
      inputs:
        duration: 30s
    - action: NetworkAllow
      block: false
      inputs:
        dest: 10.0.0.1/32 
```

The above serialization indicates a timeline
that will wait until the executable receives
the SIGUSR1 signal and then execute
a function the user has mapped to `NetworkBlock`
passing the `inputs` object as a parameter. It will
block until that function returns and then the timeline
will wait idly for 30s before calling the function
for `NetworkAllow`.

## Packaged Event Types

Technically there is little difference between an action and
a synchronizer. The serialization only differentiates them
to reason about the intent of the timeline.

### Synchronizers

#### SignalRecv
#### TimerWait

### Actions

#### LogMessage
