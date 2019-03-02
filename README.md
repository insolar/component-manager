[![Build Status](https://travis-ci.org/insolar/component-manager.svg?branch=master)](https://travis-ci.org/insolar/component-manager)
[![GolangCI](https://golangci.com/badges/github.com/insolar/component-manager.svg)](https://golangci.com/r/github.com/insolar/component-manager/)
[![Go Report Card](https://goreportcard.com/badge/github.com/insolar/component-manager)](https://goreportcard.com/report/github.com/insolar/component-manager)
[![GoDoc](https://godoc.org/github.com/insolar/component-manager?status.svg)](https://godoc.org/github.com/insolar/component-manager)
[![codecov](https://codecov.io/gh/insolar/component-manager/branch/master/graph/badge.svg)](https://codecov.io/gh/insolar/component-manager)

# Component Manager

For monolith component based architecture.

### Features 

- two step initialization
- reflect based dependency injection for interfaces
- resolving circular dependency 
- components lifecycle support
- ordered start, gracefully stop with reverse order
- easy component and integration tests with mock
- subcomponents support
- reduce boilerplate code

### Component lifecycle:

- new(just created) 
- link(inject dependency before start)
- init
- start(component can call their dependency interfaces, run goroutines)
- prepare stop(optional)
- stop (gracefully stop goroutines, close descriptors)


### Similar projects

- [facebookgo/inject](https://github.com/facebookgo/inject) - reflect based dependency injector
- [Uber FX](https://github.com/uber-go/fx) - A dependency injection based application framework
- [Google Wire](https://github.com/google/wire) - Compile-time Dependency Injection based on code generation
- [jwells131313/dargo](https://github.com/jwells131313/dargo) - Dependency Injector for GO inspired by Java [HK2](https://javaee.github.io/hk2/)
- [sarulabs/di](https://github.com/sarulabs/di) - Dependency injection framework for go programs
                                                   
