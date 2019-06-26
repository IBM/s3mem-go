## Contributing In General
Our project welcomes external contributions. If you have an itch, please feel
free to scratch it.

To contribute code or documentation, please submit a [pull request](https://github.com/ibm/s3mem-go/pulls).

A good way to familiarize yourself with the codebase and contribution process is
to look for and tackle low-hanging fruit in the [issue tracker](https://github.com/ibm/s3mem-go/issues).
Before embarking on a more ambitious contribution, please quickly [get in touch](#communication) with us.

**Note: We appreciate your effort, and want to avoid a situation where a contribution
requires extensive rework (by you or by us), sits in backlog for a long time, or
cannot be accepted at all!**

### Proposing new features

If you would like to implement a new feature, please [raise an issue](https://github.com/ibm/s3mem-go/issues)
before sending a pull request so the feature can be discussed. This is to avoid
you wasting your valuable time working on a feature that the project developers
are not interested in accepting into the code base.

### Fixing bugs

If you would like to fix a bug, please [raise an issue](https://github.com/ibm/s3mem-go/issues) before sending a
pull request so it can be tracked.

### Merge approval

The project maintainers use LGTM (Looks Good To Me) in comments on the code
review to indicate acceptance. A change requires LGTMs from two of the
maintainers of each component affected.

For a list of the maintainers, see the [MAINTAINERS.md](MAINTAINERS.md) page.

## Legal

Each source file must include a license header for the Apache
Software License 2.0. Using the SPDX format is the simplest approach.
e.g.

```
/*
################################################################################
# Copyright 2019 IBM Corp. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
################################################################################
*/
```

A githook is in place to verify if Copyright is present at commit.

We have tried to make it as easy as possible to make contributions. This
applies to how we handle the legal aspects of contribution. We use the
same approach - the [Developer's Certificate of Origin 1.1 (DCO)](https://github.com/hyperledger/fabric/blob/master/docs/source/DCO1.1.txt) - that the Linux® Kernel [community](https://elinux.org/Developer_Certificate_Of_Origin)
uses to manage code contributions.

We simply ask that when submitting a patch for review, the developer
must include a sign-off statement in the commit message.

Here is an example Signed-off-by line, which indicates that the
submitter accepts the DCO:

```Text
Signed-off-by: John Doe <john.doe@example.com>
```

You can include this automatically when you commit a change to your
local git repository using the following command:

```bash
git commit -s
```

## Communication

**FIXME** Please feel free to connect with us on our [Slack channel](link).

## Setup

If you made changes and want to test, just run `make` in the root directory will install `dep` if not yet present and then run `dep ensure -v` to update the vendor directory. Next, it will run the `go test` and calculate the test coverage. It will also check that all files contain the copyright.

## Testing

The project is tested through unit-tests.

## Coding style guidelines

- For each new S3 API implementation a new go file must be created like [get_object_request.go](https://github.com/IBM/s3mem-go/blob/4c1bd8e44612744d7772d52ca1d3070b400c24bc/s3mem/get_object_request.go#L48) along with a test file. The Non-Implemented method must be removed from the [s3mem.go](https://github.com/IBM/s3mem-go/blob/4c1bd8e44612744d7772d52ca1d3070b400c24bc/s3mem/s3mem.go#L55) file.
- The [handler.go](https://github.com/IBM/s3mem-go/blob/5395cbaf07722d69baa21a54e3462c2bc0876aa6/s3mem/handlers.go#L38) must be updated to redirect the `send` to the correct method.
- Errors must be raised using [s3memerr](https://github.com/IBM/s3mem-go/blob/5395cbaf07722d69baa21a54e3462c2bc0876aa6/s3mem/s3memerr/errors.go#L19), it implements awserr.Error.
  