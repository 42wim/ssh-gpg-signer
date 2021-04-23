# ssh-gpg-signer

Companion for <https://github.com/42wim/ssh-agentx>

Can not be used stand-alone.

Drop-in replacement for `gpg.program` config in `git`.

e.g.

```bash
git config --global gpg.program /home/user/bin/ssh-gpg-signer
git config --global commit.gpgSign true
```
