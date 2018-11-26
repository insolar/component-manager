[![Build Status](https://travis-ci.org/AndreyBronin/component-manager.svg?branch=master)](https://travis-ci.org/AndreyBronin/component-manager)
[![codecov](https://codecov.io/gh/andreybronin/component-manager/branch/master/graph/badge.svg)](https://codecov.io/gh/andreybronin/component-manager)

# Component Manager

For monolith component based architecture.

Project is inspired by [facebookgo/inject](https://github.com/facebookgo/inject)

### Features 

- two step initialization
- reflect based dependency injection for interfaces
- resolving circular dependency 
- components lifecycle support
- ordered start, gracefully stop with reverse order
- easy component and integration tests with mock
- subcomponents support

### Component lifecycle:

- new(just created) 
- link(inject dependency before start)
- init
- start(component can call their dependency interfaces, run goroutines)
- prepare stop(optional)
- stop (gracefully stop goroutines, close descriptors)
