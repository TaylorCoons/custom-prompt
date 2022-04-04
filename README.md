
## Installation

```bash
# Add to ~/.bashrc
VIRTUAL_ENV_DISABLE_PROMPT=1
PATH=$PATH:path/to/custom-prompt
PROMPT_COMMAND='PS1=$(custom-prompt)'
# Resource bash (run in cli)
source ~/.bashrc
```

Or run these in the cli (change your path still)

```bash
echo '# Custom Prompt PS1' >> ~/.bashrc
echo 'VIRTUAL_ENV_DISABLE_PROMPT=1' >> ~/.bashrc
echo 'PATH=$PATH:$HOME/custom-prompt' >> ~/.bashrc
echo "PROMPT_COMMAND='PS1=\$(custom-prompt)'" >> ~/.bashrc
source ~/.bashrc
```
