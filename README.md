# Sugarscape in Go

This project aims to reimplement the classic **Sugarscape** model as described by **Joshua M. Epstein and Robert Axtell** in their book _Growing Artificial Societies: Social Science from the Bottom Up_ (MIT Press, 1996).

Sugarscape is a well-known Agent-Based Model (ABM) used to explore complex social phenomena, such as trade, migration, and wealth distribution, by simulating autonomous agents in an environment with limited resources.

## Why Go?

The project is developed in **Go** to take advantage of:

- The **performance** of a compiled language.
- The **simplicity** and **readability** of Go’s syntax.
- The built-in **concurrency** model, which allows for efficient multi-agent simulations.

## Installation

To get started, make sure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).

Clone the repository and initialize the module:

```sh
git clone https://github.com/nicodigos/sugarscape.git
cd sugarscape-go
go mod init github.com/nicodigos/sugarscape

```
