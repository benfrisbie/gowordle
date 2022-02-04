# gowordle
[![Build Status](https://github.com/benfrisbie/gowordle/actions/workflows/release.yml/badge.svg)](https://github.com/benfrisbie/gowordle/actions/workflows/release.yml)
[![Contributors](https://img.shields.io/github/contributors/benfrisbie/gowordle)](https://github.com/benfrisbie/gowordle/graphs/contributors)
[![Activity](https://img.shields.io/github/commit-activity/m/benfrisbie/gowordle)](https://github.com/benfrisbie/gowordle/pulse)

`gowordle` is a [Wordle](https://www.powerlanguage.co.uk/wordle/) clone written in Go.

# Table of Contents
- [How to Play](#how-to-play)
    - [Example](#example)
- [Usage](#usage)
    - [Basic Usage](#basic-usage)
    - [Options](#options)

# How to Play
To win the game you must successfully guess a word within a given number of tries. After guessing a word, hints are given based on the letters positions in your guess compared to the actual solution.

The hints are:
- <span style="color:green">green</span> - letter is in correct position
- <span style="color:yellow">yellow</span> - letter is in the solution, but not in correct position
- <span style="color:red">red</span> - letter is not in word

## Example
Let's assume the correct solution is "cares". A guess of "rates" would look like this:

Guess: rates  
<span style="color:yellow">r</span><span style="color:green">a</span><span style="color:red">t</span><span style="color:green">e</span><span style="color:green">s</span> - a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z

This means:
- 'r' is in the solution, but in the wrong position (caRes)
- 'a' is in the correct position (cAres)
- 't' is not in the solution
- 'e' is in the correct position (carEs)
- 's' is in the correct position (careS)

# Usage
## Basic Usage
The easiest way to run `gowordle` is with docker.
`docker run --rm -it ghcr.io/benfrisbie/gowordle`

## Options
| Name | Description | Default |
| --- | --- | --- |
| `-max_guesses` | max number of guesses the user is allowed to make | 6 |
| `-seed` | seed to use for randomly selecting solution | current unix epoch time in nanoseconds (ex: 1643952789299435700) |
| `-debug` | enable debug logging |  |