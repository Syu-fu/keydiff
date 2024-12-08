# keydiff

[![lint](https://github.com/Syu-fu/keydiff/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/Syu-fu/keydiff/actions/workflows/lint.yml)
[![check license](https://github.com/Syu-fu/keydiff/actions/workflows/license-check.yml/badge.svg?branch=main)](https://github.com/Syu-fu/keydiff/actions/workflows/license-check.yml)
[![test](https://github.com/Syu-fu/keydiff/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/Syu-fu/keydiff/actions/workflows/test.yml)
[![Go Coverage](https://github.com/Syu-fu/keydiff/wiki/coverage.svg)](https://raw.githack.com/wiki/Syu-fu/keydiff/coverage.html)

**keydiff** is a tool for comparing structured data based on specified keys. It allows you to identify differences between datasets quickly and easily. This is especially useful for verifying updates, detecting changes, and ensuring data consistency across files or systems.

## Features

- **Key-based comparison**: Compare datasets using a specific key column.
- **Simple usage**: Just specify two CSV files and a key column.
- **Output**: Displays the differences in a human-readable format.

## Installation

There are two easy ways to install **Keydiff**:

### 1. Install via GitHub Releases

You can download the latest release from the [GitHub releases page](https://github.com/Syu-fu/keydiff/releases).

- Go to the releases page.
- Download the appropriate version for your operating system.
- Extract the archive and place the executable in a directory that’s in your system’s `PATH`.

### 2. Install via Homebrew (macOS/Linux)

If you're using Homebrew, you can install **Keydiff** directly from the Homebrew repository.

```bash
brew install Syu-fu/tap/keydiff
```

This will install the latest stable version of **Keydiff**.

### 3. Build from Source

If you'd like to build **Keydiff** from source, you can clone the repository and run the following commands:

```bash
git clone https://github.com/Syu-fu/keydiff.git
cd keydiff
go build
```

After building, you can place the `keydiff` executable in a directory that’s in your system’s `PATH` for easy access from the terminal.

## Usage

To compare two CSV files based on a specified key column, run the following command:

```bash
keydiff <original.csv> <modified.csv> --key 0
```

By default, the key column is the first column (index 0), but you can specify a different key column with the `--key` flag.

### Example

Let's say you have the following two CSV files, `original.csv` and `modified.csv`, and you want to compare them using the `name` column (index 1) as the key.

**original.csv**:

```
1,apple,1.50
2,banana,1.00
3,orange,1.25
```

**modified.csv**:

```
1,apple,1.50
2,banana,1.00
4,grape,1.00
3,orange,1.25
```

You can compare the files using the `name` column (index 1) as the key by running the following command:

```bash
keydiff --key 1 original.csv modified.csv
```

### Output

```
Key: grape
+: 4, grape, 1.00
```
