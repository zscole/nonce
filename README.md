# Nonce

A quick and dirty network QA utility developed for blockchain and distributed systems.

Currently, the warning in `containers.go` comes from Docker vendoring issues. This can be addressed by renaming or removing the vendor folder in your Docker src directory in your Go path. For example, the path to my Docker vendor file is `/Users/zak/go/src/github.com/docker/docker/vendor`. If I run the command `mv Users/zak/go/src/github.com/docker/docker/vendor /Users/zak/go/src/github.com/docker/docker/vendor2`, the `vendor` folder is renamed to `vendor2` and this error no longer arises. 

This has been a [common issue](https://github.com/moby/moby/issues/29362) with the Docker API since it made the switch from Moby. I don't really get it, but I think it's just the result of sloppy development. Working on a more efficient way of addressing this and there are a few potential options available. If you have a better idea or think I'm wrong in any way, please feel free to contribute. 

## Basic Usage Example
`nonce -i <docker image you wish to build> -n <number of nodes> -I <name of your primary interface>` 

Contact zcole@linux.com for questions, comments, or concerns. 
