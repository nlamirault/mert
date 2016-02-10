# Mert

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fmert.svg)](https://badge.fury.io/gh/nlamirault%2Fmert)

Mert is my own terminal based on VTE, a library implementing a terminal emulator widget for GTK+.


Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/mert/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/mert/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/mert/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/mert/tree/develop)


## Features

*nothing*


## Configuration

Configuration use [toml][] format. File is located into `$HOME/.config/mert/mert.toml`.

Example:

```toml
font = "Monospace 10"

[theme]
foreground = "#ffffff"
foreground_bold = "#ffffff"
background = "#000000"
cursor = "#999999"
palette = "#666666;#CC6699;#99CC66;#CC9966;#6699CC;#9966CC;#66CC99;#CCCCCC;#999999;#FF99CC;#CF062B;#CCFF99;#FFCC99;#99CCFF;#CC99FF;#99FFCC;#FFFFFF"

```

## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test


## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>




[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat

[toml]: https://github.com/toml-lang/toml
