# Configuration file

## Location of the configuration file

By default, the configuration file is located at:

```bash
${HOME}/.config/dr/config.yaml
```
### Custom configuration file

If you want to change the configuration file location,
you can set up a full path in a ***DR_CONFIG*** environment variable,
like this:

```bash
export DR_CONFIG=${HOME}/my-custom-path/dr/config-for-dr.yaml
```

Just set this up in your ***.bashrc/.zshrc*** or any other automatically
sourced files.

## Format of the configuration file

An example configuration file is as follows:

```yaml
contexts:
  - name: default
    user: myuser
    url: http://my.domain.com:1234/repo
    trusted: true
currentContext: default
```

## Multiple contexts

You can have several contexts configured in the same file, and choose the active context
with the ***currentContext*** selector.

An example of a configuration file with multiple contexts should be as follows:



```yaml
contexts:
  - name: default
    user: myuser
    url: http://my.domain.com:1234/repo
    trusted: true
  - name: production
    user: admin
    url: https://my-secure.domain:1337/prod
currentContext: production
```

# Password-protected repository

DR leverages the power of the OS keyring to store credentials in a safe manner.
To correctly login into any repository, simply setup the configuration file with your private repository:

```yaml
currentContext: passprotected
contexts:
- name: passprotected
  url: https://repo.my-company.org:5000
  user: ""
```

And then "login":

```bazaar
dr login passprotected

# Takes the form
# dr login <name of context to login>

```

You will be asked a password, that will be then stored in the OS keyring. You can update this password with "login" again.

# Passwordless repository

If your repository does not require authentication of any kind,
simply set the user field to an empty string:

```yaml
currentContext: default
contexts:
- name: default
  url: http://localhost:5000
  user: ""
  trusted: true
```
