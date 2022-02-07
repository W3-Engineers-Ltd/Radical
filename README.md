
radical
===

radical is a command-line tool facilitating development of radiant-based application.

[![Build Status](https://img.shields.io/travis/radiant/radical.svg?branch=master&label=master)](https://travis-ci.org/radiant/radical)
[![Build Status](https://img.shields.io/travis/radiant/radical.svg?branch=develop&label=develop)](https://travis-ci.org/radiant/radical)

## Requirements

- Go version >= 1.12

## Installation

To install `radical` use the `go get` command:

```bash
go get github.com/W3-Engineers-Ltd/Radical
```

Then you can add `radical` binary to PATH environment variable in your `~/.bashrc` or `~/.bash_profile` file:

```bash
export PATH=$PATH:<your_main_gopath>/bin
```

> If you already have `radical` installed, updating `radical` is simple:

```bash
go get -u github.com/W3-Engineers-Ltd/Radical
```

## Basic commands

radical provides a variety of commands which can be helpful at various stages of development. The top level commands include:

```
    version     Prints the current radical version
    migrate     Runs database migrations
    api         Creates a radiant API application
    bale        Transforms non-Go files to Go source files
    fix         Fixes your application by making it compatible with newer versions of radiant
    dlv         Start a debugging session using Delve
    dockerize   Generates a Dockerfile for your radiant application
    generate    Source code generator
    hprose      Creates an RPC application based on Hprose and radiant frameworks
    new         Creates a radiant application
    pack        Compresses a radiant application into a single file
    rs          Run customized scripts
    run         Run the application by starting a local development server

```

### radical version

To display the current version of `radical`, `radiant` and `go` installed on your machine:

```bash
$ radical version
██████   █████  ██████  ██  █████  ███    ██ ████████ 
██   ██ ██   ██ ██   ██ ██ ██   ██ ████   ██    ██    
██████  ███████ ██   ██ ██ ███████ ██ ██  ██    ██    
██   ██ ██   ██ ██   ██ ██ ██   ██ ██  ██ ██    ██    
██   ██ ██   ██ ██████  ██ ██   ██ ██   ████    ██  v1.0.0

├── radiant     : 1.7.2
├── GoVersion : go1.7.4
├── GOOS      : linux
├── GOARCH    : amd64
├── NumCPU    : 2
├── GOPATH    : /home/radicaluser/.go
├── GOROOT    : /usr/lib/go
├── Compiler  : gc
└── Date      : Monday, 26 Dec 2016
```

You can also change the output format using `-o` flag:

```bash
$ radical version -o json
{
    "GoVersion": "go1.7.4",
    "GOOS": "linux",
    "GOARCH": "amd64",
    "NumCPU": 2,
    "GOPATH": "/home/radicaluser/.go",
    "GOROOT": "/usr/lib/go",
    "Compiler": "gc",
    "radicalVersion": "1.6.2",
    "radiantVersion": "1.7.2"
}
```

For more information on the usage, run `radical help version`.

### radical new

To create a new radiant web application:

```bash
$ radical new my-web-app
██████   █████  ██████  ██  █████  ███    ██ ████████ 
██   ██ ██   ██ ██   ██ ██ ██   ██ ████   ██    ██    
██████  ███████ ██   ██ ██ ███████ ██ ██  ██    ██    
██   ██ ██   ██ ██   ██ ██ ██   ██ ██  ██ ██    ██    
██   ██ ██   ██ ██████  ██ ██   ██ ██   ████    ██  v1.0.0
2016/12/26 22:28:11 INFO     ▶ 0001 Creating application...
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/conf/
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/controllers/
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/models/
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/routers/
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/tests/
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/static/
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/static/js/
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/static/css/
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/static/img/
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/views/
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/conf/app.conf
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/controllers/default.go
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/views/index.tpl
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/routers/router.go
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/tests/default_test.go
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/main.go
2016/12/26 22:28:11 SUCCESS  ▶ 0002 New application successfully created!
```

For more information on the usage, run `radical help new`.

### radical run

To run the application we just created, you can navigate to the application folder and execute:

```bash
$ cd my-web-app && radical run
```

Or from anywhere in your machine:

```
$ radical run github.com/user/my-web-app
```

For more information on the usage, run `radical help run`.

### radical pack

To compress a radiant application into a single deployable file:

```bash
$ radical pack

██████   █████  ██████  ██  █████  ███    ██ ████████ 
██   ██ ██   ██ ██   ██ ██ ██   ██ ████   ██    ██    
██████  ███████ ██   ██ ██ ███████ ██ ██  ██    ██    
██   ██ ██   ██ ██   ██ ██ ██   ██ ██  ██ ██    ██    
██   ██ ██   ██ ██████  ██ ██   ██ ██   ████    ██  v1.0.0
2016/12/26 22:29:29 INFO     ▶ 0001 Packaging application on '/home/radicaluser/.go/src/github.com/user/my-web-app'...
2016/12/26 22:29:29 INFO     ▶ 0002 Building application...
2016/12/26 22:29:29 INFO     ▶ 0003 Using: GOOS=linux GOARCH=amd64
2016/12/26 22:29:31 SUCCESS  ▶ 0004 Build Successful!
2016/12/26 22:29:31 INFO     ▶ 0005 Writing to output: /home/radicaluser/.go/src/github.com/user/my-web-app/my-web-app.tar.gz
2016/12/26 22:29:31 INFO     ▶ 0006 Excluding relpath prefix: .
2016/12/26 22:29:31 INFO     ▶ 0007 Excluding relpath suffix: .go:.DS_Store:.tmp
2016/12/26 22:29:32 SUCCESS  ▶ 0008 Application packed!
```

For more information on the usage, run `radical help pack`.

### radical rs 
Inspired by makefile / npm scripts.
  Run script allows you to run arbitrary commands using radical.
  Custom commands are provided from the "scripts" object inside radical.json or radicalfile.

  To run a custom command, use: $ radical rs mycmd ARGS

```bash
$ radical help rs

USAGE
  radical rs

DESCRIPTION
  Run script allows you to run arbitrary commands using radical.
  Custom commands are provided from the "scripts" object inside radical.json or radicalfile.

  To run a custom command, use: $ radical rs mycmd ARGS
  
AVAILABLE SCRIPTS
  gtest
      APP_ENV=test APP_CONF_PATH=$(pwd)/conf go test -v -cover
  gtestall
      APP_ENV=test APP_CONF_PATH=$(pwd)/conf go test -v -cover $(go list ./... | grep -v /vendor/)

```

*Run your scripts with:*
```$ radical rs gtest tests/*.go```
```$ radical rs gtestall```


### radical api

To create a radiant API application:

```bash
$ radical api my-api

██████   █████  ██████  ██  █████  ███    ██ ████████ 
██   ██ ██   ██ ██   ██ ██ ██   ██ ████   ██    ██    
██████  ███████ ██   ██ ██ ███████ ██ ██  ██    ██    
██   ██ ██   ██ ██   ██ ██ ██   ██ ██  ██ ██    ██    
██   ██ ██   ██ ██████  ██ ██   ██ ██   ████    ██  v1.0.0
2016/12/26 22:30:12 INFO     ▶ 0001 Creating API...
    create   /home/radicaluser/.go/src/github.com/user/my-api
    create   /home/radicaluser/.go/src/github.com/user/my-api/conf
    create   /home/radicaluser/.go/src/github.com/user/my-api/controllers
    create   /home/radicaluser/.go/src/github.com/user/my-api/tests
    create   /home/radicaluser/.go/src/github.com/user/my-api/conf/app.conf
    create   /home/radicaluser/.go/src/github.com/user/my-api/models
    create   /home/radicaluser/.go/src/github.com/user/my-api/routers/
    create   /home/radicaluser/.go/src/github.com/user/my-api/controllers/object.go
    create   /home/radicaluser/.go/src/github.com/user/my-api/controllers/user.go
    create   /home/radicaluser/.go/src/github.com/user/my-api/tests/default_test.go
    create   /home/radicaluser/.go/src/github.com/user/my-api/routers/router.go
    create   /home/radicaluser/.go/src/github.com/user/my-api/models/object.go
    create   /home/radicaluser/.go/src/github.com/user/my-api/models/user.go
    create   /home/radicaluser/.go/src/github.com/user/my-api/main.go
2016/12/26 22:30:12 SUCCESS  ▶ 0002 New API successfully created!
```

For more information on the usage, run `radical help api`.

### radical hprose

To create an Hprose RPC application based on radiant:

```bash
$ radical hprose my-rpc-app

██████   █████  ██████  ██  █████  ███    ██ ████████ 
██   ██ ██   ██ ██   ██ ██ ██   ██ ████   ██    ██    
██████  ███████ ██   ██ ██ ███████ ██ ██  ██    ██    
██   ██ ██   ██ ██   ██ ██ ██   ██ ██  ██ ██    ██    
██   ██ ██   ██ ██████  ██ ██   ██ ██   ████    ██  v1.0.0
2016/12/26 22:30:58 INFO     ▶ 0001 Creating application...
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/conf/
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/controllers/
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/models/
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/routers/
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/tests/
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/static/
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/static/js/
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/static/css/
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/static/img/
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/views/
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/conf/app.conf
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/controllers/default.go
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/views/index.tpl
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/routers/router.go
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/tests/default_test.go
    create   /home/radicaluser/.go/src/github.com/user/my-rpc-app/main.go
2016/12/26 22:30:58 SUCCESS  ▶ 0002 New application successfully created!
```

For more information on the usage, run `radical help hprose`.

### radical bale

To pack all the static files into Go source files:

```bash
$ radical bale

██████   █████  ██████  ██  █████  ███    ██ ████████ 
██   ██ ██   ██ ██   ██ ██ ██   ██ ████   ██    ██    
██████  ███████ ██   ██ ██ ███████ ██ ██  ██    ██    
██   ██ ██   ██ ██   ██ ██ ██   ██ ██  ██ ██    ██    
██   ██ ██   ██ ██████  ██ ██   ██ ██   ████    ██  v1.0.0
2016/12/26 22:32:41 INFO     ▶ 0001 Loading configuration from 'radical.json'...
2016/12/26 22:32:41 SUCCESS  ▶ 0002 Baled resources successfully!
```

For more information on the usage, run `radical help bale`.

### radical migrate

For database migrations, use `radical migrate`.

For more information on the usage, run `radical help migrate`.

### radical generate

radical also comes with a source code generator which speeds up the development.

For example, to generate a new controller named `hello`:

```bash
$ radical generate controller hello

██████   █████  ██████  ██  █████  ███    ██ ████████ 
██   ██ ██   ██ ██   ██ ██ ██   ██ ████   ██    ██    
██████  ███████ ██   ██ ██ ███████ ██ ██  ██    ██    
██   ██ ██   ██ ██   ██ ██ ██   ██ ██  ██ ██    ██    
██   ██ ██   ██ ██████  ██ ██   ██ ██   ████    ██  v1.0.0
2016/12/26 22:33:58 INFO     ▶ 0001 Using 'Hello' as controller name
2016/12/26 22:33:58 INFO     ▶ 0002 Using 'controllers' as package name
    create   /home/radicaluser/.go/src/github.com/user/my-web-app/controllers/hello.go
2016/12/26 22:33:58 SUCCESS  ▶ 0003 Controller successfully generated!
```

For more information on the usage, run `radical help generate`.

### radical dockerize

radical also helps you dockerize your radiant application by generating a Dockerfile.

For example, to generate a Dockerfile with `Go version 1.6.4` and exposing port `9000`:

```bash
$ radical dockerize -image="library/golang:1.6.4" -expose=9000

██████   █████  ██████  ██  █████  ███    ██ ████████ 
██   ██ ██   ██ ██   ██ ██ ██   ██ ████   ██    ██    
██████  ███████ ██   ██ ██ ███████ ██ ██  ██    ██    
██   ██ ██   ██ ██   ██ ██ ██   ██ ██  ██ ██    ██    
██   ██ ██   ██ ██████  ██ ██   ██ ██   ████    ██  v1.0.0
2016/12/26 22:34:54 INFO     ▶ 0001 Generating Dockerfile...
2016/12/26 22:34:54 SUCCESS  ▶ 0002 Dockerfile generated.
```

For more information on the usage, run `radical help dockerize`.

### radical dlv

radical can also help with debugging your application. To start a debugging session:

```bash
██████   █████  ██████  ██  █████  ███    ██ ████████ 
██   ██ ██   ██ ██   ██ ██ ██   ██ ████   ██    ██    
██████  ███████ ██   ██ ██ ███████ ██ ██  ██    ██    
██   ██ ██   ██ ██   ██ ██ ██   ██ ██  ██ ██    ██    
██   ██ ██   ██ ██████  ██ ██   ██ ██   ████    ██  v1.0.0
2017/03/22 11:17:05 INFO     ▶ 0001 Starting Delve Debugger...
Type 'help' for list of commands.
(dlv) break main.main
Breakpoint 1 set at 0x40100f for main.main() ./main.go:8

(dlv) continue
> main.main() ./main.go:8 (hits goroutine(1):1 total:1) (PC: 0x40100f)
     3:	import (
     4:		_ "github.com/user/myapp/routers"
     5:		radiant "github.com/W3-Engineers-Ltd/Radiant/server/web"
     6:	)
     7:	
=>   8:	func main() {
     9:		radiant.Run()
    10:	}
    11:
```

For more information on the usage, run `radical help dlv`.

## Shortcuts

Because you'll likely type these generator commands over and over, it makes sense to create aliases:

```bash
# Generator Stuff
alias g:a="radical generate appcode"
alias g:m="radical generate model"
alias g:c="radical generate controller"
alias g:v="radical generate view"
alias g:mi="radical generate migration"
```

These can be stored , for example, in your `~/.bash_profile` or `~/.bashrc` files.

## Help

To print more information on the usage of a particular command, use `radical help <command>`.

For instance, to get more information about the `run` command:

```bash
$ radical help run
USAGE
  radical run [appname] [watchall] [-main=*.go] [-downdoc=true]  [-gendoc=true] [-vendor=true] [-e=folderToExclude]  [-tags=goBuildTags] [-runmode=RADIANT_RUNMODE]

OPTIONS
  -downdoc
      Enable auto-download of the swagger file if it does not exist.

  -e=[]
      List of paths to exclude.

  -gendoc
      Enable auto-generate the docs.

  -main=[]
      Specify main go files.

  -runmode
      Set the radiant run mode.

  -tags
      Set the build tags. See: https://golang.org/pkg/go/build/

  -vendor=false
      Enable watch vendor folder.

DESCRIPTION
  Run command will supervise the filesystem of the application for any changes, and recompile/restart it.
```

## Contributing
Bug reports, feature requests and pull requests are always welcome.

We work on two branches: `master` for stable, released code and `develop`, a development branch.
It might be important to distinguish them when you are reading the commit history searching for a feature or a bugfix,
or when you are unsure of where to base your work from when contributing.

### Found a bug?

Please [submit an issue][new-issue] on GitHub and we will follow up.
Even better, we would appreciate a [Pull Request][new-pr] with a fix for it!

- If the bug was found in a release, it is best to base your work on `master` and submit your PR against it.
- If the bug was found on `develop` (the development branch), base your work on `develop` and submit your PR against it.

Please follow the [Pull Request Guidelines][new-pr].

### Want a feature?

Feel free to request a feature by [submitting an issue][new-issue] on GitHub and open the discussion.

If you'd like to implement a new feature, please consider opening an issue first to talk about it.
It may be that somebody is already working on it, or that there are particular issues that you should be aware of
before implementing the change. If you are about to open a Pull Request, please make sure to follow the [submissions guidelines][new-pr].

## Submission Guidelines

### Submitting an issue

Before you submit an issue, search the archive, maybe you will find that a similar one already exists.

If you are submitting an issue for a bug, please include the following:

- An overview of the issue
- Your use case (why is this a bug for you?)
- The version of `radical` you are running (include the output of `radical version`)
- Steps to reproduce the issue
- Eventually, logs from your application.
- Ideally, a suggested fix

The more information you give us, the more able to help we will be!

### Submitting a Pull Request

- First of all, make sure to base your work on the `develop` branch (the development branch):

```
  # a bugfix branch for develop would be prefixed by fix/
  # a bugfix branch for master would be prefixed by hotfix/
  $ git checkout -b feature/my-feature develop
```

- Please create commits containing **related changes**. For example, two different bugfixes should produce two separate commits.
A feature should be made of commits splitted by **logical chunks** (no half-done changes). Use your best judgement as to
how many commits your changes require.

- Write insightful and descriptive commit messages. It lets us and future contributors quickly understand your changes
without having to read your changes. Please provide a summary in the first line (50-72 characters) and eventually,
go to greater lengths in your message's body. A good example can be found in [Angular commit message format](https://github.com/angular/angular.js/blob/master/CONTRIBUTING.md#commit-message-format).

- Please **include the appropriate test cases** for your patch.

- Make sure all tests pass before submitting your changes.

- Rebase your commits. It may be that new commits have radicaln introduced on `develop`.
Rebasing will update your branch with the most recent code and make your changes easier to review:

  ```
  $ git fetch
  $ git rebase origin/develop
  ```

- Push your changes:

  ```
  $ git push origin -u feature/my-feature
  ```

- Open a pull request against the `develop` branch.

- If we suggest changes:
  - Please make the required updates (after discussion if any)
  - Only create new commits if it makes sense. Generally, you will want to amend your latest commit or rebase your branch after the new changes:

    ```
    $ git rebase -i develop
    # choose which commits to edit and perform the updates
    ```

  - Re-run the tests
  - Force push to your branch:

    ```
    $ git push origin feature/my-feature -f
    ```

[new-issue]: #submitting-an-issue
[new-pr]: #submitting-a-pull-request

## Licence

```text
Copyright 2016 radical authors

                   GNU LESSER GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.


  This version of the GNU Lesser General Public License incorporates
the terms and conditions of version 3 of the GNU General Public
License, supplemented by the additional permissions listed below.

  0. Additional Definitions.

  As used herein, "this License" refers to version 3 of the GNU Lesser
General Public License, and the "GNU GPL" refers to version 3 of the GNU
General Public License.

  "The Library" refers to a covered work governed by this License,
other than an Application or a Combined Work as defined below.

  An "Application" is any work that makes use of an interface provided
by the Library, but which is not otherwise based on the Library.
Defining a subclass of a class defined by the Library is deemed a mode
of using an interface provided by the Library.

  A "Combined Work" is a work produced by combining or linking an
Application with the Library.  The particular version of the Library
with which the Combined Work was made is also called the "Linked
Version".

  The "Minimal Corresponding Source" for a Combined Work means the
Corresponding Source for the Combined Work, excluding any source code
for portions of the Combined Work that, considered in isolation, are
based on the Application, and not on the Linked Version.

  The "Corresponding Application Code" for a Combined Work means the
object code and/or source code for the Application, including any data
and utility programs needed for reproducing the Combined Work from the
Application, but excluding the System Libraries of the Combined Work.

  1. Exception to Section 3 of the GNU GPL.

  You may convey a covered work under sections 3 and 4 of this License
without being bound by section 3 of the GNU GPL.

  2. Conveying Modified Versions.

  If you modify a copy of the Library, and, in your modifications, a
facility refers to a function or data to be supplied by an Application
that uses the facility (other than as an argument passed when the
facility is invoked), then you may convey a copy of the modified
version:

   a) under this License, provided that you make a good faith effort to
   ensure that, in the event an Application does not supply the
   function or data, the facility still operates, and performs
   whatever part of its purpose remains meaningful, or

   b) under the GNU GPL, with none of the additional permissions of
   this License applicable to that copy.

  3. Object Code Incorporating Material from Library Header Files.

  The object code form of an Application may incorporate material from
a header file that is part of the Library.  You may convey such object
code under terms of your choice, provided that, if the incorporated
material is not limited to numerical parameters, data structure
layouts and accessors, or small macros, inline functions and templates
(ten or fewer lines in length), you do both of the following:

   a) Give prominent notice with each copy of the object code that the
   Library is used in it and that the Library and its use are
   covered by this License.

   b) Accompany the object code with a copy of the GNU GPL and this license
   document.

  4. Combined Works.

  You may convey a Combined Work under terms of your choice that,
taken together, effectively do not restrict modification of the
portions of the Library contained in the Combined Work and reverse
engineering for debugging such modifications, if you also do each of
the following:

   a) Give prominent notice with each copy of the Combined Work that
   the Library is used in it and that the Library and its use are
   covered by this License.

   b) Accompany the Combined Work with a copy of the GNU GPL and this license
   document.

   c) For a Combined Work that displays copyright notices during
   execution, include the copyright notice for the Library among
   these notices, as well as a reference directing the user to the
   copies of the GNU GPL and this license document.

   d) Do one of the following:

       0) Convey the Minimal Corresponding Source under the terms of this
       License, and the Corresponding Application Code in a form
       suitable for, and under terms that permit, the user to
       recombine or relink the Application with a modified version of
       the Linked Version to produce a modified Combined Work, in the
       manner specified by section 6 of the GNU GPL for conveying
       Corresponding Source.

       1) Use a suitable shared library mechanism for linking with the
       Library.  A suitable mechanism is one that (a) uses at run time
       a copy of the Library already present on the user's computer
       system, and (b) will operate properly with a modified version
       of the Library that is interface-compatible with the Linked
       Version.

   e) Provide Installation Information, but only if you would otherwise
   be required to provide such information under section 6 of the
   GNU GPL, and only to the extent that such information is
   necessary to install and execute a modified version of the
   Combined Work produced by recombining or relinking the
   Application with a modified version of the Linked Version. (If
   you use option 4d0, the Installation Information must accompany
   the Minimal Corresponding Source and Corresponding Application
   Code. If you use option 4d1, you must provide the Installation
   Information in the manner specified by section 6 of the GNU GPL
   for conveying Corresponding Source.)

  5. Combined Libraries.

  You may place library facilities that are a work based on the
Library side by side in a single library together with other library
facilities that are not Applications and are not covered by this
License, and convey such a combined library under terms of your
choice, if you do both of the following:

   a) Accompany the combined library with a copy of the same work based
   on the Library, uncombined with any other library facilities,
   conveyed under the terms of this License.

   b) Give prominent notice with the combined library that part of it
   is a work based on the Library, and explaining where to find the
   accompanying uncombined form of the same work.

  6. Revised Versions of the GNU Lesser General Public License.

  The Free Software Foundation may publish revised and/or new versions
of the GNU Lesser General Public License from time to time. Such new
versions will be similar in spirit to the present version, but may
differ in detail to address new problems or concerns.

  Each version is given a distinguishing version number. If the
Library as you received it specifies that a certain numbered version
of the GNU Lesser General Public License "or any later version"
applies to it, you have the option of following the terms and
conditions either of that published version or of any later version
published by the Free Software Foundation. If the Library as you
received it does not specify a version number of the GNU Lesser
General Public License, you may choose any version of the GNU Lesser
General Public License ever published by the Free Software Foundation.

  If the Library as you received it specifies that a proxy can decide
whether future versions of the GNU Lesser General Public License shall
apply, that proxy's public statement of acceptance of any version is
permanent authorization for you to choose that version for the
Library.

```
