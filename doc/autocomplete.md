# Autocomplete

To enable autocomplete of commands, you will need to source the appropriate file
in the [.autocomplete](../.autocomplete) folder.
As an example, if you are using bash, you will have to source the [.autocomplete/bash_autocomplete](../.autocomplete/bash_autocomplete).
If you are using zsh, you will source [.autocomplete/zsh_autocomplete](../.autocomplete/zsh_autocomplete). You can setup this in a place
where it will get automatically sourced (usually /etc/bash_completion.d/dr), or you can copy the file to somewhere else and source
it manually (in your ***.zshr*** for example).

As an example, we could have in our ***~/.zsh***:

```bash
source ~/.autocomplete/zsh_autocomplete
```

Obviously, you can change the name of the file to something more descriptive (e.g:
 from ***zsh_autocomplete** to ***dr_autocomplete***).