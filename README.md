# gowordle
[![Build Status](https://img.shields.io/github/workflow/status/benfrisbie/gowordle/ci?label=ci&logo=github&style=flat-square)](https://github.com/benfrisbie/gowordle/actions?workflow=ci)
[![Contributors](https://img.shields.io/github/contributors/benfrisbie/gowordle)](https://github.com/benfrisbie/gowordle/graphs/contributors)
[![Activity](https://img.shields.io/github/commit-activity/m/benfrisbie/gowordle)](https://github.com/benfrisbie/gowordle/pulse)

`gowordle` is a [Wordle](https://www.powerlanguage.co.uk/wordle/) clone written in Go for the terminal.

How to run:
```
docker run --rm -it ghcr.io/benfrisbie/gowordle
```

# Table of Contents
- [How to Play](#how-to-play)
    - [Hints](#hints)
    - [Example](#example)
- [Usage](#usage)
    - [Options](#options)

# How to Play
Wordle is a game where you guess words in hopes of determining the final solution. After guessing a word, hints are given based on the letters positions in your guess compared to the actual solution. These hints should be used to determine yor next guess.

## Hints
The hints are color coded:
| color | description |
| --- | --- |
| 🟢 | letter is in solution and in correct position |
| 🟡 | letter is in solution, but in wrong position |
| 🔴 | letter is not in solution |

## Example
![example.png](media/example.png?raw=true)
As you can see the final solution is `along`. Let's walk through the example one guess at a time:
- guess #1 = `rates`
    - 🟡 - `a` is in solution, but in wrong position
    - 🔴 - `r`,`t`,`e`, and `s` are not in solution
- guess #2 = `ample`
    - 🟢 - `a` is in correct position
    - 🟡 - `l` is in solution, but in wrong position
    - 🔴 - `m`,`p`, and `e` are not in solution
- guess #3 = `aloud`
    - 🟢 - `a`, `l`, `o` are in correct positions
    - 🔴 - `u` and `d` are not in solution
- guess #4 = `alone`
    - 🟢 - `a`, `l`, `o`, `n` are in correct positions
    - 🔴 - `e` is not in solution
- guess #5 = `along`
    - 🟢 - `a`, `l`, `o`, `n`, `g` are in correct positions


# Usage
`docker run --rm -it ghcr.io/benfrisbie/gowordle [OPTIONS]`

## Options
| Name | Description |
| ---- | --- |
| `-background_color` | use background color for hints instead of text color |
| `--debug` | enable debug logging |
| `--max_guesses` | max number of guesses |
| `--seed` | random seed to use |
