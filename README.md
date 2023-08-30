### [RootMe : write-up's]

This is my project for solving hacking problems

## Programming
- TCP - Back to school  [solution](./programming/back_to_school)
- TCP - Encoded string  [solution](./programming/encoded_string)
- TCP - The Roman wheel [solution](./programming/rot13)
- TCP - Uncompress Me   [solution](./programming/uncompress_me)

## Warning
To start the project, you must first install go
```shell
curl -L https://git.io/vQhTU | bash -s -- --version 1.21
```

## Utility
to start any relay, you can use the make commands
```shell
 make back_to_school
 run back_to_school...
 go run ./programming/back_to_school/main.go -a "challenge01.root-me.org:52002"
 flag: [+] Good job ! Here is your flag: **{***_*******_***_****}
```
To view the progress of the solved task, I added a simple utility
```shell
# add to .env
#ROOT_ME_API_KEY=<api_key>
#ROOT_ME_API_UIT=<uid>
make stats
./bin/stats -a https://api.www.root-me.org -k $(ROOT_ME_API_KEY) -u $(ROOT_ME_API_UID)
process...
UserName        Rank            Position
<username>      <rank>          <position>
.
├── Programmation
│   ├── TCP - Uncompress Me
│   ├── TCP - La roue romaine
│   ├── TCP - Chaîne encodée
│   └── TCP - Retour au collège
└── Cryptanalyse
    ├── Hash - SHA-2
    ├── Hash - Message Digest 5
    ├── Encodage - UU
    └── Encodage - ASCII
solved: 8
```