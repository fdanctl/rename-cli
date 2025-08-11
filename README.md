Batch file renaming utility with append, prepend, and replace options

### Installation

1. Clone the repository:

```sh
git clone git@github.com:fdanctl/rename-cli.git
```

2. Navigate to the project directory:

```sh
cd rename-cli
```

3. Install with Go:

```sh
go install
```

Make sure your GOPATH is set as an environment variable and pointing to a valid Go workspace.
Once installed, you can use the rename-cli command to append, prepend, or replace filenames according to your needs.

### Commands

#### Append

Append adds a string to the end of filenames using optional numbering and date formatting:

```sh
rename-cli append -n 1 --dry-run ' - %YYYY-%MM-%DD' *.jpg
```

#### Prepend

Prepend adds a string to the beginning of filenames using optional numbering and date formatting:

```sh
rename-cli prepend -n 1 --dry-run '%YYYY-%MM-%DD_%hh:%mm:%ss' *.jpg
```

#### Replace

Replace patterns in filenames with a new string using optional numbering and date formatting:

```sh
rename-cli replace -n 1 'old_pattern' 'new_pattern_%0n' \*.txt
```

#### Flags

| Flag              | Description                                       |
| ----------------- | ------------------------------------------------- |
| `-n` or `--start` | Starting value for %n numbering (default: `1`)    |
| `--dry-run`       | Show what would be renamed without making changes |

### Explanation of % Placeholders

| Placeholder              | Description                                                   | Example              |
| ------------------------ | ------------------------------------------------------------- | -------------------- |
| `%n`, `%0n`, `%00n`, ... | Sequential number starting from the --start value (default 1) | `%0n → 01, %n → 1`   |
| `%YYYY`                  | Four-digit year                                               | Current year: `2025` |
| `%YY`                    | Two-digit year                                                | Current year: `25`   |
| `%MM`                    | Two-digit month                                               | Current month: `08`  |
| `%DD`                    | Two-digit day                                                 | Current day: `11`    |
| `%hh`                    | Two-digit hour                                                | Current hour: `21`   |
| `%mm`                    | Two-digit minute                                              | Current minute: `45` |
| `%ss`                    | Two-digit second                                              | Current second: `30` |
