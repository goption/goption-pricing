# golang-build-scripts

A bucket of resources for developing Golang projects

## A Short Disclaimer

I made these scripts to use on Golang projects hosted on Gitlab, not for public comsumption, although you're welcome to use them as long as your usage is in line with the GNU AGPLv3 license, which can be found in the LICENSE file in the same directory as this readme. If you're not sure, ask.

## Contents

1. [Getting Started](#Getting-Started)
1. [Submodule](#Submodule)
1. [Globals](#Globals)
1. [.gitlab-ci.yml](#.gitlab-ci.yml)
    1. [Usage](#Usage)
    1. [Extending](#Extending)
    1. [Conventions](#Conv)
1. [Makefiles](#Makefiles)
    1. [Inheritance](#Inheritance)
    1. [Templates](#Templates)
    1. [Conventions](#Conventions)
    1. [local](#local)

## Getting Started

When setting up a new repository, perform the following steps:

1. Create the repository in Gitlab CI
1. Determine whether to use this library as a [submodule](#Submodule-Usage)
1. Copy the boilerplate from the (#Globals) section

test coverage parsing in config ui
 add pipeline status/coverage report to readme
 read about "git strategy for pipelines" to see which is better and document
 Maybe need to change runners settings? Check defaults for analyst
 Mention predefined secret variables <- they are mentioned in th yml file
 Protect branches and tags v*
 Turn on goption-deploy deploy key under Settings->Repository
 Make sure JIRA integration is set up

## Submodule

Docker
Makefiles
Add yml for golint-ci and goreleaser https://goreleaser.com/ https://github.com/golangci/golangci-lint
Deployment of `go doc html` to s3
Upload container images to ECR https://medium.com/@stijnbe/using-gitlab-ci-with-aws-container-registry-ecaf4a37d791
Implement the things auto devops does where they make sense, which is probs not on these projects https://docs.gitlab.com/ce/ci/examples/README.html this has good links
Make a task to have the runner use the docker runner with ECS if possible, once that's all set up

## Globals

New repositories need a certain amount of boilerplace. This repository will contain the most up-to-date versions of all the boilerplate files you need. It contains three files that should be used with all projeccts:

- .editorconfig
- .gitignore
- LICENSE

What about `.gitlab-ci.yml`, you say? And what about a `Dockerfile`? This repository is meant to be a fount of continuous integration, and those files aren't always one-size-fits all. Read on!

## .gitlab-ci.yml

The base `gitlab-ci.sample.yml` (no preceding dot) provides a starting point for your continuous integration. Copy and paste it into your repository, and [alter it](https://docs.gitlab.com/ce/ci/) to include the things you need. If this repo is being used as a [submodule](#Submodule), you **cannot** configure GLCI to use the sample file as the CI configuration file under `Settings -> CI / CD -> General Pipeline Settings -> Custom CI config path` because the submodule will not exist until the repository has been cloned.

### Usage

Assuming this repository is checked out as a submodule of yours, you should be able to copy the `gitlab-ci.sample.yml` to `.gitlab-ci.yml` in your project root, update the paths of the includes to include the path to this repository (per the comment in the file), and be good to go. The sample is indeed just a sample, although it will get you started.

Inside the `gitlab-ci` directory are some yml files that define hidden jobs and snippets that can be implemented, extended, or composited in your CI configuration. Your CI configuration should consist of job definitions that consume the snippets defined in these files; if they do not fit your needs, define a job that does in its respective file and make sure to make it extensible. Use `$CI_PROJECT_NAME` instead of hardcoding the project name, for example, and define variables where necessary. Keep in mind that job-level variables will override global variables, but environment variables override anything defined in the yml file.

One last thing: for Golang projects, sometimes it's handy to parse out the module path such as `github.com/onwsk8r/golang-build-scripts`. You know from Bash school that you can find that with [parameter expansion](http://wiki.bash-hackers.org/syntax/pe) by typing `${CI_PROJECT_URL#*//}`. This **will not work** in the `variables` section; it must be done in one of the `script` sections.

### Extending

Use the yml content by _importing_ the correct file(s) into your project's `.gitlab-ci.yml` as shown in the sample. The imports use some [special YAML features](https://docs.gitlab.com/ce/ci/yaml/#hidden-keys-jobs) to create hidden jobs (that don't run) that you can include in your job, much like inheritance or composition in programming. There are also some snippets that can be composited into other jobs. If you are not familiar with YML syntax such as `<<:`, the documentation linked above takes about 100 seconds to read. Here is an example that shows the Vatik usage conventions:

```yaml
# The job is a hidden job; it must be merged into an existing job to run
# The name of the job follows the <build-stage>_<name> convention
.test_lint: &test_lint
  stage: test # The job includes the stage name
  before_script: &test_lint_before_script # Each major component is also aliased
    - go get -u gopkg.in/alecthomas/gometalinter.v2
    - gometalinter.v2 install
  script: &test_lint_script # Aliases for parts of a job are <job-name>_<part-name>
    - gometalinter.v2 ./...

# Use the job above in its entirety as follows:
lint:
  <<: *test_lint
  dependencies:
    - build

# Or like this in Gitlab (requires Libre >= 11.4 or less for nonfree versions)
lint:
  extends: .test_lint
  dependencies:
    - build

# If the entire job cannot be used as-is, use the necessary parts instead:
lint: # Gometalinter is already installed
  script: *test_lint_script
  dependencies:
    - build
```

### Conventions

All of the jobs and snippets in this repository follow some basic conventions.

- The job is a hidden job because it starts with a `.`. It can only run by being merged into another job.
- The name clearly identifies what it contains. Names like `lint` and `test` are fairly self-explanatory, but if you're creating an extensible cache to hold builds, name it `build_cache` instead of just cache.
- The job has a `stage` key. This is necessary for using it in its entirety.
- The major components of the job, such as `script`, `before_script`, and `variables` each have their own aliases so parts of the job can be used or extended. When in doubt, alias it.

The most important thing to remember is that everything **must be extensible**. This is a CI configuration file, and different projects can have different requirements. That's a major part of the motivation behind using snippets instead of just jobs.

## Makefiles

Everyone loves a good Makefile. Generally you should create a Makefile in your project root with the following contents:

```Makefile
include path/to/this/repo/Makefile
```

...so long as your repository doesn't need to customize the functionality that this repo provides:

- Testing
- Linting
- Vendoring/`go mod`
- Local setup (ie pre-commit hook)

You can get more information about what the Makefile provides by running `make help` or just `make`.

### Inheritance

You'll notice this repo's Makefile is suspiciously lacking in recipes. Specifically, it only contains generics such as the `help` target and several jobs without recipes. The reason for this is to provide an extensible, overridable system with a standard interface. Ideally the end user only needs to invoke the jobs defined in the main Makefile. In software development this is called the Adapter Pattern.

 The variables in the Makefile are all defined with `?=`, which allows you to override them in your project's Makefile. Of special interest are the `LINTER` and `TESTER` variables: note that Makefiles with those name are included later. The value of those variables (and consequently the name of the Makefile) indicate which test framework (eg `ginkgo`, `testify`, `goconvey`) and which linting framework (ie `glcilint` for golangci-lint or `gometalinter`) the application will use. The included Makefiles redefine those variables (also with `?=`) so they are functional Makefiles themselves.

This model of inheritance allows you, the end user, to create resuable job definitions and use them in your application by overriding a couple of variables. The Makefiles in this repository should _never_ assign a variable using `=`; they should always use `?=` to give users like you the option to override those variables. Each Makefile in the Makefiles directory should correspond to a piece of software. The only exception the variable rule is that the `SHELL` variable defaults to `/bin/sh`, which lacks some necessary functionality.

### Templates

The main Makefile defines a number of jobs that depend on other jobs, which are not defined. This templating system provides not only a standard interface to the user through encapsulation, but also ensures that the concrete implementation of each tool is both complete and standardized: what works with one tool will work with another. If a particular tool is not suited for the application, simply create a new Makefile for the correct tool, update the corresponding variable in your application's Makefile, and everything should just work.

Ignoring `local.mk`, there are Makefiles for linting and Makefiles for testing. if a tool were capable of doing both it would need to have Makefiles for each purpose to maintain the extensibility and customizability of the system. The system is set up so that each part can be used invidually or as a group. With both testers and linters, we have a few common goals: we need to install the linters, we need to run them on all files (for CI) or only changed files (for a pre-commit hook), and we need to clean up after them. Because of that there is a lot of crossover between the two:

**Test Tools** must contain the following jobs:

- **_install_tools**: After this job, the tool should be configured and in the $PATH.
- **_test**: This job should invoke the test framework.
- **_test_changed**: This job should use `git diff --cached`
- **_clean_test**: This potentially empty job should clean up after the tester
- **_install_coverage**: This job installs any coverage-specific prerequisites.
- **_coverage**: Generate a coverage report called `$(COVERAGE)`.
- **_clean_coverage**: This job should clean up any coverage files (such as `$(COVERAGE)`)

**Linting Tools** must contain the following jobs:

- **_install_lint**: After this job, the tool should be configured and in the $PATH.
- **_lint**: This job should invoke the linter.
- **_lint_changed**: This job should use `git diff --cached`
- **_clean_lint**: This potentially empty job should clean up after the linter.

**Dependency Management Tools** must contain the following jobs:

- **_pre_depend**: After this job, the dependency tool should be configured and in the $PATH.
- **_depend**: This job should install the dependencies (the equivalent of `dep ensure`).
- **_update_depend**: This job should update the vendor dependencies.
- **_clean_depend**: This job should remove the vendor dependencies, such as the `vendor/` folder.

### Conventions

Makefiles in this repository must conform to some simple rules.

- Global variables should be `IN CAPS`, while variables that are used by one job, for one job should be `lower_case`. They should be defined with `?=` so they can be overridden.
- Job names should be in `spinal_case`. Jobs that a user would not likely call directly (such as `_install_lint`) should be `_prefixed` with an underscore. Other jobs should have a `# Job Description` on the same line.
- Jobs such as `_lint` should not depend on `_install_lint`, but `lint` should depend on both `_install_lint` and `_lint`. This allows separation in CI where `_install_lint` may be called `before_script`, while `_lint` would be called in the `script`.
- All Makefiles should be functional by themselves and be as extensible as possible. That means they should define every variable they use, and they should define every variable with `?=`.

Other important rules:

- `install_*` functions should not do anything if the program is already installed. Check for the existence of a binary using, for example, `ifeq(,$(shell command $(binary_name)))`. This works on Powershell, OSX, and *nix (but not `cmd.exe`).
- Arguments to programs should come from a variable. It should be trivial to alter the behavior of any program without making changes to this repository.

### local

`make local` sets up a local development environment. It:

- Installs the chosen linter and tester
- Installs any necessary dependency management software
- Creates a pre-commit git hook to lint and test changed packages

It does **not** install dependencies. It is up to you to run `make depend`.
