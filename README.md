<div align="center">
   <img src="banner.png" width="67%"> <!--67 haha-->
</div>

<p align="center">
   The hourly journal.
   <br/>
   <br/>
   <a href="#contribution">
      <img src="https://img.shields.io/badge/status-WIP-orange">
   </a>
   <a href="https://goreportcard.com/report/github.com/sharkmu/jhournal">
      <img src="https://goreportcard.com/badge/github.com/sharkmu/jhournal" />
   </a>
   <a href="https://github.com/sharkmu/jhournal/blob/master/LICENSE">
      <img alt="GitHub License" src="https://img.shields.io/github/license/sharkmu/jhournal?cacheSeconds=1">
   </a>
   <br/> 
   <a href="https://github.com/sharkmu/jhournal/issues">
      <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="CodeFactor" />
   </a>
   <a href="https://github.com/sharkmu/jhournal/tags" rel="nofollow">
      <img src="https://img.shields.io/github/v/tag/sharkmu/jhournal?include_prereleases&label=version"/>
   </a>
</p>

**Jhournal** is a kind of journal that only lets you write a short (77 character) entry every hour.
It is written entirely in [Go](https://go.dev/) with [Fyne](https://fyne.io/), a cross-platform GUI toolkit.

## Features
> Note: This project is still a ***WIP***, so you might come across some errors. In such case please [create an issue](https://github.com/sharkmu/jhournal/issues/new) about it

There are 3 tabs on the left side of the window.
1. **New entry** - Create a new journal entry.
2. **View entries** - Browse and read your previous entries.
3. **Settings** - Change your configurations.


## Contribution
Feel free to contribute to the project. Be aware that this is a beginner project, so the code may sometimes be slightly harder to read. You can also collaborate by making an issue or giving feedback. Or even just by editing this README.md and making a PR.

## Installation
### Manually
- Clone the repository with [git](https://git-scm.com/): `git clone https://github.com/sharkmu/jhournal.git`
- Make sure that you have [Go](https://go.dev/) installed
- Go to the repository's folder: `cd jhournal`
- Install the necessary packages: `go mod tidy`
- To run the repository: `go run .`
- To build the repository: `go build`, if this doesn't work on Windows, then instead try running: `go build -ldflags="-s -w -H=windowsgui" -o jhournal.exe .`

### Download from [Releases](https://github.com/sharkmu/jhournal/releases)
You can download the binary file for Jhournal from the latest release in the [Releases](https://github.com/sharkmu/jhournal/releases) tab


## About the name
Below you can see an [Excalidraw](https://excalidraw.com/) that I made, when I was trying to name the project:
<p align="center">
   <img src="name.png" width="50%">
</p>