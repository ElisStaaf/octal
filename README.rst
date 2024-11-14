Octal - LOC counter
-------------------
.. image:: https://img.shields.io/badge/Build%20(Fedora)-passing-2a7fd5?logo=fedora&logoColor=2a7fd5&style=for-the-badge
   :alt: Build = Passing
   :target: https://github.com/ElisStaaf/octal
.. image:: https://img.shields.io/badge/Version-1.0.1-2dd245?style=for-the-badge
   :alt: Version = 1.0.1
   :target: https://github.com/ElisStaaf/octal
.. image:: https://img.shields.io/badge/Language-Go-20c9df?logo=Go&style=for-the-badge
   :alt: Language = Go
   :target: https://github.com/ElisStaaf/octal

Octal: A subpar LOC counter, written in Go. But I mean, it works! And you, yes, even *you*, can 
use it. Enjoy the project and contributions are welcome!

Requirements
------------
* `go`_
* `make`_
* `git`_ or `gh`_



Install
-------
Firstly, you'd want to install to your computer:

.. code:: bash

   # git
   git clone https://github.com/ElisStaaf/octal ~/octal

   # gh
   gh repo clone ElisStaaf/octal ~/octal

Then, you'd want to build the binary to your `/usr/bin`:

.. code:: bash

   sudo make install

Then, you're completely free to use it. *BUUUUUUUUUT*, right now it doesn't have support for counting comments, it can
only count empty lines and non-empty lines. |defstrike| \This *will* be fixed soon.\ |endstrike|\  This will *never* be fixed.
I promise.

Usage
-----
.. code:: bash

   octal [options] <file>

.. _`go`: https://go.dev/doc/install
.. _`git`: https://git-scm.com/downloads 
.. _`gh`: https://github.com/cli/cli#installation
.. _`make`: https://www.gnu.org/software/make

.. |defstrike| raw:: html

    <strike>

.. |endstrike| raw:: html

    </strike>
