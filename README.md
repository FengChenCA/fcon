# FCON
FCON is a partially implemented, simple HTTP server that mimics some of the basic operations of a container runtime.  This repository is intended used as an exercise to evaluate a developer with minimal Golang knowledge and working knowledge of REST APIs.  An experienced Golang developer should be able to complete this exercise in less than an hour.    

## Getting Started
Directions for building, running, testing, etc... are intentionally missing.  This exercise is designed to test your resourcefulness and ability to learn from documentation, in addition to evaluating your coding skills.  

### Coding
The sections that you need to implement are marked with a `//TODO:` comment.  Some sections are small, others might be relatively large.  Code is *not* restricted to the immediate location of the TODO, feel free to add code elsewhere as you see fit.

### Resources
  - https://tour.golang.org/welcome/1
  - https://golang.org/doc/
  - https://golang.org/pkg/
  - https://gobyexample.com/
  
## Requirements
The requirements are intentionally open ended.  There is no "one right answer"; that being said, not all answers are correct.  Your code will be judged on readability, correctness, reusability, etc...

### Basic
  - Compiles and does not crash (no panics!)
  - Code is properly formatted according to Golang standards
  - No warnings raised by vet

### Server
  - Port must be configurable
  - Server must have some form of debug output
  - Persistence is not required

### APIs
POST and PUT endpoints must support JSON bodies.
  - /containers
    - get: JSON response with all containers
    - post: create a new container, JSON response of the newly created container
  - /containers/{name}
    - get: JSON response of specific container
    - put: update an existing container, JSON response of updated container
    - delete: remove an existing container, JSON response of deleted container

### Containers
  - PIDs must not be mutable
  - Name must be unique at all times
  - Name cannot be changed while container is running
  - State must be valid at all times

### Responses
  - Appropriate status codes must be used
  - Non-error responses must have a JSON body

### Testing
  - Formal tests are not required, though they aren't frowned upon either!

