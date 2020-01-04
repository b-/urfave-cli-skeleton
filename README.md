# urfave/cli skeleton
Basic skeleton to start a urfave cli project with debug and quiet flags and structured ./cmd directory.

# Usage
This will be an example of how you would start a new project with this skeleton. Example project will be in the repo 
github.com/luthermonson/example in the directory ~/go/example.

```shell
cd ~/go
git clone --depth=1 --branch=master git@github.com:luthermonson/urfave-cli-sekeleton.git example
rm -rf ./example/.git
cd example
git init
go mod init github.com/luthermonson/example
go mod tidy && go mod vendor
sed -i 's/github.com\/luthermonson\/urfave-cli-skeleton/github.com\/luthermonson\/example/' main.go
git add -A
git commit -m "initializing"
git remote add origin git@github.com:luthermonson/example.git
git push -u origin master
```