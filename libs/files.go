package libs

import (
    "path/filepath"
    "strings"
    "os/exec"
    "os"

)


func GetCurrentDirectory() string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
    if err != nil {
        log.Fatal(err)
    }

    return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
func GetAppPath() string {
    file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
    index := strings.LastIndex(path, string(os.PathSeparator))
    return path[:index]
}
