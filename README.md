cksumfly
=======

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](http://www.opensource.org/licenses/mit-license.php)

## Description

Print CRC checksum for flyway of each FILE.

## Installation

    % go get github.com/vvatanabe/cksumfly/cmd/cksumfly/

## Synopsis

    % cksumfly /path/to/foo.sql
    % cksumfly --path /path/to/foo.sql
    % cat /path/to/foo.sql | cksumfly
    
## Options

```
Options
  -p, --path= .sql file path
```

## Bugs and Feedback

For bugs, questions and discussions please use the GitHub Issues.

## Author

[vvatanabe](https://github.com/vvatanabe)