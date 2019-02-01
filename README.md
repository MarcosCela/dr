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

You can enable shell autocomplete scripts (in fact, it is highly recommended!), read how here: [autocomplete].


## Usage

Once you have correctly downloaded and configured dr, you can start using it:

### Lis existing images:

```bash
$dr ls
{
  "images": [
    {
      "name": "alpine"
    },
    {
      "name": "another/busy/subpath"
    },
    {
      "name": "busybox"
    },
    {
      "name": "busybox/another/subpath/testing"
    },
    {
      "name": "busybox/testing"
    },
    {
      "name": "busybox/testing"
    }
  ]
}
```

Listing tags for a specific image:
```bash
dr tags busybox/testing 
{
  "tags": [
    "latest",
    "mytag"
  ]
}

```

Or getting information about a specific image and tag:
```bash
dr manifest busybox:1.29 
{
  "compressedSize": "729.5 kB",
  "digest": "sha256:758ec7f3a1ee85f8f08399b55641bfb13e8c1109287ddc5e22b68c3d653152ee",
  "image": "busybox",
  "layers": 1,
  "repository": "localhost:5000",
  "tag": "1.29"
}
```

You can also play with the flags to use different output formats:

```bash
dr manifest busybox:1.29 --output plain
repository: localhost:5000
image: busybox
name: "1.29"
digest: sha256:758ec7f3a1ee85f8f08399b55641bfb13e8c1109287ddc5e22b68c3d653152ee
layers: 1
compressedSize: 729.5 kB
```
```bash
dr ls --output yaml
images:
- name: alpine
- name: another/busy/subpath
- name: busybox
- name: busybox/another/subpath/testing
- name: busybox/testing
- name: busybox/testing
```

## Disclaimer

This is a personal project aimed at learning Golang, but perhaps it will
suit your needs, so feel free to use it, fork it or contribute as you please.

[kubectl]: https://kubernetes.io/docs/reference/kubectl/overview/
[configuration documentation]: doc/configuration-file.md
[Source install]: doc/install/SOURCE_INSTALL.md
[Binary download]: doc/install/BINARY_INSTALL.md
[autocomplete]: doc/autocomplete.md