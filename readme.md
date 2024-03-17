![](https://i.imgur.com/9ucpYgY.png)

## Installation

`go install xudongz.com/code/gitstreak@latest`

## Usage

You can use gitsteak to display the graph for any number of repositories.

`gitstreak path/to/repo1 path/to/repo2`

You can limit it to a specific user by specifying their email address.

`gitstreak -author user@example.org path/to/repo`

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
