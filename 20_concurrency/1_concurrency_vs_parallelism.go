package main

/**

Some important historical context re: when GoLang came into development.

- In 2006, Intel released the first dual core CPU (multiple core processors).
- In 2007, Google started to create Go to natively take advantage of multiple cores. It was the first software programming language to do this.
- Big significance. Means that concurrency was built into Go's thinking from day 1.
- In other programming languages, it can be much harder to do concurrency.
- Go makes concurrency easy (according to Todd). It is at the heart of the design of the programming language.
- March 2012 - first version of GoLang was released

Concurrency VS parallelism:

- If you write code in Go and you run that code on a machine that only has 1 CPU, your code is not going to run in parallel.
- You're not going to have multiple threads of code in parallel running at the same time.
- It's going to run sequentially, one instruction after the other, on the 1 CPU.
- If you have more than 1 CPU, you're going to have the opportunity to run code in parallel.
- This is parallelism.

- So what's concurrency then?
- Concurrency is a design pattern. It's a way you can write your code.
- You can write concurrent code.
- You can have code which has the possibility and potential to run in parallel
- If you have multiple cores (CPUs), that concurrent code can run in parallel.
- Writing concurrent code does NOT guarantee that your code is going to run in parallel (because it depends on the number of CPUs)

//he recommends a talk by Rob Pike - concurrency is not parallelism - you tube

*/
