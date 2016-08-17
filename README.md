## Realize

[![AUR](https://img.shields.io/aur/license/yaourt.svg?maxAge=2592000?style=flat-square)](https://raw.githubusercontent.com/tockins/realize/v1/LICENSE)

Run, build and watch file changes with custom paths

#### Features

- Build, Install and Run in the same time
- Live reload on file changes (re-build, re-install and re-run)
- Watch custom paths
- Watch specific file extensions
- Multi project support

#### Installation and usage

- Run this for get/install it:

    ```
    $ go get github.com/tockins/realize/...
    ```
- From the root of your project/projects:

    ```
    realize start 
    ```
    Will create a realize.config.yaml file with a sample project.
    
    You can pass additional parameters for your first project, such as the project name, the main file name and the base path. 
    
    ```
    realize start --name="Project Name" --main="main.go" --base="/"
    ```
- Add another project whenever you want    

    ```
    realize add --name="Project Name" --main="main.go" --base="/"
    ```
- Remove a project by his name

    ```
    realize remove --name="Project Name"
    ```
- Lists all projects

    ```
    realize list
    ```
- Build, Run and watch file changes. Realize will re-build and re-run your projects on each changes

    ```
    realize run 
    ```


#### To do
- [x] Command start - default config file
- [x] Command add - new project on the config file 
- [x] Command remove - remove project from the config file
- [x] Command watch - watch changes and rebuild 
- [x] Command list - print projects list
- [x] Remove duplicate projects
- [x] Support for multiples projects
- [x] Watcher files preview
- [x] Support for directories with duplicates names
- [ ] Unit test
- [ ] Documentation
- [x] Support for server start/stop 
- [x] Cli feedback


