[![Build Status](https://travis-ci.org/AndreyBronin/component-manager.svg?branch=master)](https://travis-ci.org/AndreyBronin/component-manager)
[![codecov](https://codecov.io/gh/andreybronin/component-manager/branch/master/graph/badge.svg)](https://codecov.io/gh/andreybronin/component-manager)

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
- [Google Wire](https://blog.golang.org/wire) - Compile-time Dependency Injection based on code generation
- [jwells131313/dargo](https://github.com/jwells131313/dargo) - Dependency Injector for GO inspired by Java [HK2](https://javaee.github.io/hk2/)
- [sarulabs/di](https://github.com/sarulabs/di) - Dependency injection framework for go programs
                                                   
