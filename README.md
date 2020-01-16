# GoTestParallelization

This project illustrate the different usage of parallelization in Go

## Some concepts

### go test -p

The `go test` command has a `-p n` argument to define the number of test binaries that can be run in parallel. The default is the number of CPUs available.

Using `-p 1` set only one program and means that every (non-parallelized) tests will be executed consecutively.

Go will build a test binary for each package containing tests.

### t.Parallel()

The testing package exposes a Parallel() method that tells to Go that it can run the test in parallel with other parallel tests, in a same test binary.

## Concrete examples

### Using go test -p

As Go will build a test binary for each package and run test binaries in parallel if we set `-p` to 2 or more, it means it is useful only if tests are in different packages. And by package, we mean path (not the package declaration at the beginning of a file).

The commands `make run-serial-tests-in-1-program` and `make run-serial-tests-in-2-programs` shows the difference of time execution if we provide `-p 1` or `-p 2`.

In the first case, Go will execute all tests in `./tests/serial/` consecutively and it will take around 25 seconds.

In the second case, Go will execute the test binary for `./tests/serial/singletests` in parallel of the test binary for `./tests/serial/subtests` and it will take around 15 seconds.

### Using t.Parallel() at test level

Using `t.Parallel()` allows to run parallelized tests in parallel in a single test binary.

The commands `run-single-serial-tests` and `run-single-parallel-tests` shows the difference of time execution if we parallelize tests or not.

In the first case, Go will execute serial tests consecutively and it will take around 10 seconds.

In the second case, Go will execute parallized tests concurrently in the same test binary and it will take around 5 seconds.

### Extending t.Parallel() to subTests

Using Table Driven Tests is useful to run a same scenario with different inputs. It is also possible to parallelize the subtests by providing the `t.Parallel()` piece of code inside the `t.Run()` call. For more details, see https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721.

### Using got test -p and t.Parallel()

The commands `make run-all-tests-in-1-program` and `make run-all-tests-in-4-programs` shows the difference of time execution if we optimize the parallelization with all methodes.

In the first case, it will take 5+10 seconds for parallized tests + 5+5+5+10 seconds for serial tests, so a total of 40 seconds.

In the seconds cas, it will take 15 seconds.