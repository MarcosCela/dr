![logo](doc/logo.png)

Docker remote is a CLI that attempts to provide a fast and intuitive way
to interact with a remote docker registry.

It is heavily inspired in [kubectl], and attempts to copy some concepts,
like the use of contexts, or different output formats.

## Install

You can install dr in your system in 2 different ways:

- [Source install]: you will clone the repository and build dr yourself.
- [Binary download]: you will download a pre-built binary of dr.


You will also need a basic configuration file, so you can tell ***dr*** where to find
the repositories. To learn more about how to set up the configuration file, read
the [configuration documentation], and it is greatly recommended to use the autocomplete files.

## Autocomplete

To enable autocomplete of commands, you will need to source the appropriate file
in the [.autocomplete](.autocomplete) folder.
As an example, if you are using bash, you will have to source the [.autocomplete/bash_autocomplete](.autocomplete/bash_autocomplete).
If you are using zsh, you will source [.autocomplete/zsh_autocomplete](.autocomplete/zsh_autocomplete). You can setup this in a place
where it will get automatically sourced (usually /etc/bash_completion.d/dr), or you can copy the file to somewhere else and source
it manually (in your ***.zshr*** for example).

As an example, we could have in our ***~/.zsh***:

```bash
source ~/.autocomplete/zsh_autocomplete
```

Obviously, you can change the name of the file to something more descriptive (e.g:
 from ***zsh_autocomplete** to ***dr_autocomplete***).

## Disclaimer

This is a personal project aimed at learning Golang, but perhaps it will
suit your needs, so feel free to use it, fork it or contribute as you please.

[kubectl]: https://kubernetes.io/docs/reference/kubectl/overview/
[configuration documentation]: doc/configuration-file.md
[Source install]: doc/install/SOURCE_INSTALL.md
[Binary download]: doc/install/BINARY_INSTALL.md