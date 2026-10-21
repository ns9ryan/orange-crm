package utils

import (
	"fmt"
	"orange-crm/global"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
)

func getPythonCmd() string {
	if runtime.GOOS == "windows" {
		return "python"
	}

	wd, err := os.Getwd()
	if err != nil {
		return "python3"
	}

	return filepath.Join(wd, "lib", "myenv", "bin", "python")
}

// MakeQrcode 生成qrcode
func MakeQrcode(employeeNumber string) (output string, err error) {
	wd, _ := os.Getwd()
	// python脚本路径
	pyPath := filepath.Join(GetDirByName("lib"), "main.py")
	// 输出路劲
	outputPath := filepath.Join(wd, global.CONFIG.Upload.Qrcode.SavePath, fmt.Sprintf("%s.png", employeeNumber)) // 输出路劲
	//outputPath := filepath.Join(ProjectRoot(), global.CONFIG.Upload.Qrcode.SavePath, fmt.Sprintf("%s.png", employeeNumber)) // 输出路劲
	pyCmd := getPythonCmd()
	// 构建命令，例：python D:\AA\GO\orange-crm\lib\main.py 2c97c9c4-e00e-4910-b30b-85a7a430b7d2 0 2c97c9c4-e00e-4910-b30b-85a7a430b7d2 D:\AA\GO\orange-crm\upload\qrcode\2c97c9c4-e00e-4910-b30b-85a7a430b7d2.png
	cmd := exec.Command(pyCmd, pyPath, employeeNumber, employeeNumber, strconv.Itoa(0), employeeNumber, outputPath)
	res, err := cmd.CombinedOutput()
	output = string(res)
	return
}
