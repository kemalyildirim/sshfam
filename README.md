#### Goal

* I can connect to a remote server via `sshfam connect user@remote_ip`. After that, the ssh session datas are saved.
* I can list saved connections via `sshfam list`, and I can select one of the connections to connect them, with no password asked.
* Update any of the connections in the `sshfam list` output by the command `sshfam update <alias>`
* I can transfer files/folders without more easily than `scp`. i.g. `sshfam cp /home/testuser/myfile.zip .` (assuming there is one active session in sshfam)
* Delete saved sessions in `sshfam list`. i.g. `sshfam delete 'ubuntuserver1'`