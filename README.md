# timeline

Library for synchronizing event sequences

## Summary

This is a simple library that can generate a described
sequence of events to allow for testing of code response.

## Sequence

A timeline sequence is composed of synchronizers and
actions. Synchronizers indicate timeline should
synchornize the event flow with something such as
an OS signal or a timer expiration. Actions simply
call functions in your code base.

## Example

```yaml
  timeline:
    - sync: Timer
      inputs:
        duration: 10s
    - action: NetworkBlock
      block: true 
      inputs:
        dest: 10.0.0.1/32
```

The above serialization indicates a timeline
that will wait 10 seconds and then execute
a function passing the `inputs` object
as a parameter. It will block until that function
returns and then the timeline is complete.

