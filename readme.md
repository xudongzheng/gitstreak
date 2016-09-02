![](https://i.imgur.com/9ucpYgY.png)

## Installation

If you have [Go](https://golang.org/) installed, simply use `go get` to install
gitstreak.

`go get github.com/xudongzheng/gitstreak`

## Usage

You can use gitsteak to display the graph for any number of repositories.

`gitstreak path/to/repo1 path/to/repo2`

You can limit it to a specific user by specifying their email address.

`gitstreak -author user@example.org path/to/repo`

You can name multiple email addresses by separating them by a comma.

`gitstreak -author user@example.org,another@example.org path/to/repo`

If you have a single directory with all your repositories, you can easily
include each repository.

`gitstreak -author user@example.org path/to/parent/*`

You can define an alias within your `.bashrc` file if you do not want to specify
a list of repository every time.

## Example Graphs

`gitstreak go`

![](https://i.imgur.com/dzouqZQ.png)

`gitstreak linux`

![](https://i.imgur.com/zcSxFr3.png)

`gitstreak -author torvalds@linux-foundation.org linux`

![](https://i.imgur.com/m8TtLAy.png)

`gitstreak node`

![](https://i.imgur.com/FW0cp6M.png)

## Author

[Xudong Zheng](https://www.xudongz.com/) 

## License

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation and/or
other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
may be used to endorse or promote products derived from this software without
specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

