# socket-sysl

Inspired by Sysl and sysl-go, socket-sysl is a code generation framework for WebSocket-powered apps and services.


## Contents

socket-sysl contains a few components:

### codegen/

Code generation scripts that will populate a `gen/` folder.

### pkg/

A library that generated code uses.

### example/

An example of a socket-sysl app. This is not a part of socket-sysl.


## Mental Model

Building an event-driven application can be a disconcerting experience if approached with a resource-oriented mindset. The resource-oriented straitjacket can feel very comfortable, since any unit of work typically corresponds to either adding a new resource, or fiddling with an operation on a resource.

In contrast, the event-driven model is looser and gets straight to the point. Logic is grouped by "things that happen", which is a much larger space than that of "resources that exist". It's very easy to model disparate events that prevent the application logic from cohering around a consistent model.

To encourage some structure, start by thinking of events as use cases (a la [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)). When the user embarks on a use case, that action will produce an event. Like behaviour-driven development, one should keep the focus on the domain of things that are meaningful to the user, and not get sucked into implementation details.

### Role of WebSockets

In this mindset, WebSockets are one such implementation detail. It should be possible to implement the same logic behind a gRPC adapter, for example.

Ideally event handling should be stateless, with any necessary context passed along with the event.
