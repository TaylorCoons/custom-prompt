
## Installation

Add custom-prompt to path if not in default paths
```bash
PATH=$PATH:/path/to/custom-prompt
```

Set custom PS1 in bash
```bash
PROMPT_COMMAND='PS1=$(custom-prompt)'
```

Or run these in the cli (change your path still)

```bash
echo '# Custom Prompt PS1' >> ~/.bashrc
echo 'PATH=$PATH:$HOME/custom-prompt' >> ~/.bashrc
echo "PROMPT_COMMAND='PS1=\$(custom-prompt)'" >> ~/.bashrc
source ~/.bashrc
```