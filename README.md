# todo

`todo` scans your project for all TODO comments and generates a shitlist of things todo.

I'm unwell at the time of writing this, so my mental state is degraded: It's a great opportunity to teach myself programming again, in a way. 

Because of this mental state, I' m more able to [write code like it's my first time](https://prog21.dadgum.com/87.html) and focus on shipping, rather than beautiful software.

## Installation

To install `todo`, you need to have Go installed on your machine. Then, run the following command:

```bash
git clone github.com/hjfitz/todo
make install
```

## Usage

Simply run `todo` from the root of your project:

```plaintext
$ todo
/home/hjf/code/self/todo/main.go:
  41: check for ~/.config/todo.conf and load that
```

## Ignoring Files

`todo` ignores files and folders specified by your `.gitignore`.

