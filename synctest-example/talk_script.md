Lightning talk

---

Hello team,

Today, I am going to talk about a Go testing feature called "synctest."  (Or: "Today, I'll be talking about Go's 'synctest' feature.")

---

Before diving into the main part, let me share the motivation for this presentation.

The current latest version is 1.24, and it includes some cool features, such as Go tool directives and improved algorithms for data structures.

Go also added a new experimental testing feature that is not recommended for production use yet.

However, since this feature is for testing and doesn't affect production code, I believe it's a good time to explore it.

Effective testing boosts productivity and improves our systems and applications, so I hope you all gain some inspiration for enhancing your projects.

---

Here is the table of contents for my talk. (Or: "Here's what I'll cover today.")

First, I'll share the basics of Go testing for those who are not familiar with Go.

Apologies to the senior Gophers!

Then, I'll discuss the challenges of time-related testing in Go.

Finally, the highlight of this presentation: I'll show you how `synctest` solves these challenges.

Okay, let's begin.

---

First, I'll explain the basics of Go testing.

Unlike Python or PHP, Go doesn't require extra libraries for executing tests.

Go has a built-in `test` command and can run test code in files with the `_test.go` suffix.

Test files should contain test functions that have `Test` as a prefix (with a capital T).

The argument of the test function is `t` (lowercase), which is of type `*testing.T` from the `testing` package.  `testing` is an official package, so we don't need to import it explicitly.

`*testing.T` allows you to report errors when the actual output of a function or method doesn't match the expected output.

---

Let's look at an example in VS Code.

Here's a simple test case.

Consider this simple code:

I've created a `calculate` function that takes two `int` parameters and returns their sum.

The corresponding test code looks like this.

First, we import the necessary packages. Then, we call the `Add` function and compare the actual result with the expected result.

If these two values don't match, we report an error.

To test multiple cases, we can define the arguments and expected values in a slice of structs, like this.

The `tests` variable is a slice of structs, each containing three `int` values: `a`, `b`, and `want`.

We can then test these values using subtests.

This is very simple, isn't it?

---

Next, let's consider time-dependent tests.

To simplify things, let's say we want to test a function that waits for a specific duration.

The function might look like this:

Go provides the `time` package for handling time-related operations, and the `Sleep` function waits for the given duration.

To check if the wait time is accurate, we might write this test code:

Before calling the `sleepWrapper` function, we get the current time. Then, we calculate the elapsed time and compare it with the expected duration.

In theory, this test seems to work, but if we run it, the test often fails.

This is due to delays in code execution. To accurately check the function, the test code should be like this:

The expected time should be *less than or equal to* the actual time. This means we have to accept a range rather than comparing for an exact value. (Or: "This forces us to accept a range instead of comparing for an exact value.")

---

I've discussed the issue of time-dependent tests. The challenge is how to verify the precise passage of time.

The experimental `synctest` feature can solve this.

This package is primarily for testing concurrent processes, but it can also be used for simpler scenarios.

---

Let's see an example of testing the same function using `synctest`.

We can use `synctest` like this:

(Here, you *must* explain the code.  Something like: "Here, `synctest.StartTimer` creates a timer that we can control.  We then run the function, and `synctest.CheckTimer` allows us to verify that the expected duration has elapsed.  Crucially, `synctest` allows us to advance the virtual time, so the test doesn't have to wait for real time.")

Let's execute this test. The test passes, demonstrating that `synctest` enables accurate time testing.

Interestingly, the execution time is less than 1 second, meaning the test didn't wait for a full second of real time.

So, if we want to test this function with multiple durations, it takes much longer without `synctest`. Using `synctest` allows the tests to finish quickly.

This significantly improves both cost and developer experience.

In many cases, CI pipelines include a test execution step. Longer test times increase CI costs.

`synctest` minimizes test execution time, reducing CI costs. This is a major benefit.

Furthermore, reducing test time minimizes the time developers spend waiting for feedback.

---

To summarize my presentation:

Go provides built-in testing capabilities with its standard packages, eliminating the need for third-party libraries.

Time-dependent tests can be flaky, but `synctest` allows us to accurately test time durations and concurrent programs.

Additionally, reducing test execution time offers substantial benefits.

That's all for my presantation, thank you for listening.
