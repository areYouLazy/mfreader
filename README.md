# mfreader

A tool inspired by [mfdread](https://github.com/zhovner/mfdread) to print MIFARE products dumps in a human readable format

## Index

- [Usage](#usage)
- [Examples](#examples)
- [References](#references)

## Usage

```console
root@localhost$ mfreader -h
mfreader, a tool inspired by mfdread ( https://github.com/zhovner/mfdread )
to print MIFARE dumps in a human readable format

Usage:
  mfreader [flags]

Flags:
  -f, --file string    dump file to analyze
  -h, --help           help for mfreader
  -j, --json           output in JSON format
  -m, --manufacturer   print manufacturer data
```

## Examples

```console
root@localhost$ mfreader -f dump.mfd -m
+----------+-----+-----+------+
| UID      | BCC | SAK | ATQA |
+----------+-----+-----+------+
| 33bd9d3f | 2c  | 98  | 02   |
+----------+-----+-----+------+
```

```console
root@localhost$ mfreader -f dump.mfd -m -j
{
  "UID": "33bd9d3f",
  "BCC": "2c",
  "SAK": "98",
  "ATQA": "02"
}
```

## References

- https://www.nxp.com/docs/en/data-sheet/MF1S50YYX_V1.pdf
- https://www.nxp.com/docs/en/application-note/AN10833.pdf
