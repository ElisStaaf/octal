# Octal - LOC counter
[![Version](https://img.shields.io/badge/Version-1.0.0-2dd245?style=for-the-badge)](https://github.com/ElisStaaf/octal)
[![Build](https://img.shields.io/badge/Build%20(Fedora)-passing-2a7fd5?logo=fedora&logoColor=2a7fd5&style=for-the-badge)](https://github.com/ElisStaaf/octal)
[![Language](https://img.shields.io/badge/Language-Go-20c9df?logo=Go&style=for-the-badge)](https://github.com/ElisStaaf/octal)  
OCTOL: A subpar LOC counter, written in Go. But I mean, it works! And you, yes, even *you*, can 
use it. Enjoy the project and contributions are welcome!

# Requirements
* The "Go" programming language
* Make
* Git or Github CLI

# Install
Firstly, you'd want to install to your computer:
```bash
# git
git clone https://github.com/ElisStaaf/octal ~/octal

# gh
gh repo clone ElisStaaf/octal ~/octal
```
Then, you'd want to build the binary to your `/usr/bin`:
```bash
sudo make install
```
Then, you're completely free to use it.

# Usage
```bash
octal <file>
```
