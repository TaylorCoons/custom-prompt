
## Installation

Download the deb file located in the github release

Install the deb 
```bash
sudo apt install ./path/to/deb/custom-prompt_{version}_{arch}.deb
```

Add environment configuration to the ~/.bashrc
```bash
# Add to ~/.bashrc
VIRTUAL_ENV_DISABLE_PROMPT=1
PROMPT_COMMAND='PS1=$(custom-prompt)'
# Re-source bash (run in terminal)
source ~/.bashrc
```

Or run these in the terminal

```bash
echo '# Custom Prompt PS1' >> ~/.bashrc
echo 'VIRTUAL_ENV_DISABLE_PROMPT=1' >> ~/.bashrc
echo "PROMPT_COMMAND='PS1=\$(custom-prompt)'" >> ~/.bashrc
source ~/.bashrc
```

## Uninstall

```bash
sudo apt remove custom-prompt
```

Remove environment variables from ~/.bashrc
