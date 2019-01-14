package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("goinstall repo version [sub_dir]")
        fmt.Println("eg. goinstall github.com/oxequa/realize 2.0.2")
        return
    }
    repo := os.Args[1]
    version := os.Args[2]
    //optional sub directory
    var optPath string
    if len(os.Args) > 3 {
        optPath = os.Args[3]
    }

    goPath := os.Getenv("GOPATH")
    os.RemoveAll(goPath + "/src/" + repo)
    fmt.Printf("simulating: go get -u %s@%s\n", repo, version)
    cmd := exec.Command("go", "get", "-u", repo)
    cmd.Dir = goPath
    if err := cmd.Run(); err != nil {
        panic(err)
    }

    cmd = exec.Command("git", "checkout", fmt.Sprintf("tags/v%s", version))
    cmd.Dir = goPath + "/src/" + repo
    fmt.Printf("git checkout tags/v%s (%s)\n", version, cmd.Dir)
    if err := cmd.Run(); err != nil {
        panic(err)
    }

    cmd = exec.Command("go", "get")
    cmd.Dir = goPath + "/src/" + repo
    fmt.Printf("go get\n")
    if err := cmd.Run(); err != nil {
        panic(err)
    }

    subPath := optPath
    if subPath != "" {
        subPath = "/" + subPath
    }

    cmd = exec.Command("go", "install")
    cmd.Dir = goPath + "/src/" + repo + subPath
    fmt.Printf("go install (%s)...\n", cmd.Dir)
    if err := cmd.Run(); err != nil {
        panic(err)
    }

    fmt.Printf("Installed %s@%s\n", repo+subPath, version)
}
