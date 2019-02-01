# Install DR from sources

To install the project from sources, clone it to your preferred directory:

```bash
git clone git@github.com:MarcosCela/dr.git --depth=1 --branch=stable
```

Then build it:

```bash
# Generate the binary called 'dr'
go build

# Copy the dr binary to a directory that is in your $PATH variable
cp dr ${DIRECTORY_IN_MY_PATH}/dr

# Make it executable!
chmod +x ${DIRECTORY_IN_MY_PATH}/dr

# Ensure it works!
dr help
```

Remember to set the basic [configuration]!

The project is built with the following go version:

```bash
go version go1.11 linux/amd64
```

[configuration]: ../configuration-file.md
