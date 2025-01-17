# Pathfinder algorithm

Simple CLI pathfinder algorithm, written in GO

## Goal

Make a pathfinder algorithm capable of finding the shortest path between two points using a 2D array as a map.

It must :

- Travel **only up/down** and **left/right**, _no diagonals allowed_
- Travel through 1s and avoid 0s
- Return the distance of the shortest path
- Return the map with the path displayed
- Return an error message if there are no paths, invalid data is provided, etc...

## Resources

- https://pkg.go.dev/flag
- https://github.com/alexflint/go-arg
- https://github.com/spf13/cobra

## Usage

Just use `make` to know what you can do.

## TODO

### MVP

- [ ] Write a few basic matrices & display them
- [ ] Solve manual maps
- [ ] Generate random maps
- [ ] Solve 'em all!
  - [ ] Handle _valid_ maps
  - [ ] Handle _invalid_ maps
- [ ] Display solved map
  - [ ] Display stats (path length, time...)
  - [ ] Display solved map with colors

### v1

- [ ] Parse JSON input
- [ ] Simple command line argument to get input matrix

### v1.x

- [ ] Display map (fancier)
- [ ] Better CLI
