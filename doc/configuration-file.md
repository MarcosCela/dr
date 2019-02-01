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
    pass: mypassword
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
    pass: mypassword
    url: http://my.domain.com:1234/repo
    trusted: true
  - name: production
    user: admin
    pass: my-super-s3cure-p@ss
    url: https://my-secure.domain:1337/prod
currentContext: production
```

# Passwordless repository

If your repository does not require authentication of any kind,
simply set the user and password fields to an emtpy string:

```yaml
currentContext: default
contexts:
- name: default
  url: http://localhost:5000
  user: ""
  pass: ""
  trusted: true
```
