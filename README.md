
## Installation

Install the deb file
Append below to the ~/.bashrc
```bash
# Add to ~/.bashrc
VIRTUAL_ENV_DISABLE_PROMPT=1
PROMPT_COMMAND='PS1=$(custom-prompt)'
# Resource bash (run in cli)
source ~/.bashrc
```

Or run these in the cli

```bash
echo '# Custom Prompt PS1' >> ~/.bashrc
echo 'VIRTUAL_ENV_DISABLE_PROMPT=1' >> ~/.bashrc
echo "PROMPT_COMMAND='PS1=\$(custom-prompt)'" >> ~/.bashrc
source ~/.bashrc
```
