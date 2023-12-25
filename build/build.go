package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// 获取当前项目目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("无法获取当前目录：", err)
		return
	}
	// 先删除bin目录
	err = os.RemoveAll(currentDir + "/build/bin")
	if err != nil {
		fmt.Println("删除文件出错：", err)
		return
	}
	fmt.Println("currentDir: ", currentDir)
	// 构建命令
	cmd := exec.Command("wails", "build") // 替换为你想要执行的命令
	// 设置命令的工作目录为当前项目目录
	// 设置命令的工作目录为当前项目目录
	cmd.Dir = currentDir
	// 设置命令的环境变量为当前系统的环境变量
	cmd.Env = os.Environ()
	// 执行命令并获取输出
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("命令执行出错：", err)
		return
	}
	// 打印命令输出结果
	fmt.Println("构建完毕：", string(output))
	fmt.Println("开始移动配置文件：")
	CopyDir(currentDir+"/data", currentDir+"/build/bin/data", true)
	fmt.Println("配置文件移动完成！开始压缩为zip文件")
	ZipDir(currentDir+"/build/bin", currentDir+"/build/bin.zip")
	fmt.Println("压缩完成！开始删除源文件")
	// 删除bin目录,但是保留bin.zip
	err = RemoveFiles(currentDir + "/build/bin")
	if err != nil {
		fmt.Println("删除文件出错：", err)
		return
	}
	os.RemoveAll(currentDir + "/build/bin")
	fmt.Println("文件夹压缩成功并删除源文件！开始移动zip文件")
	MoveFile(currentDir+"/build/bin.zip", currentDir+"/build/bin")
}

func MoveFile(filePath, destinationDir string) error {
	// 确保目标目录存在，如果不存在则创建
	err := os.MkdirAll(destinationDir, os.ModePerm)
	if err != nil {
		return err
	}
	// 获取文件名
	_, fileName := filepath.Split(filePath)
	// 构建目标路径
	destinationPath := filepath.Join(destinationDir, fileName)
	// 移动文件
	err = os.Rename(filePath, destinationPath)
	if err != nil {
		return err
	}

	return nil
}

/*
*
overwrite 是否覆盖源文件
*/
func CopyDir(sourceDir, destinationDir string, overwrite bool) error {
	// 创建目标目录
	err := os.MkdirAll(destinationDir, 0755)
	if err != nil {
		return err
	}
	// 遍历源目录
	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 构建目标文件或目录的路径
		destPath := filepath.Join(destinationDir, path[len(sourceDir):])
		if info.IsDir() {
			// 创建目标目录
			err = os.MkdirAll(destPath, 0755)
			if err != nil {
				return err
			}
		} else {
			// 复制文件
			err = copyFile(path, destPath, overwrite)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func copyFile(sourcePath, destinationPath string, overwrite bool) error {
	// 检查目标文件是否已存在
	if _, err := os.Stat(destinationPath); err == nil {
		if !overwrite {
			// 目标文件已存在且不允许覆盖
			return nil
		}
	}
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}

func ZipDir(sourceDir, destinationFile string) error {
	// 创建目标zip文件
	zipFile, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 创建zip.Writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历源目录
	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		// 构建zip文件中的路径
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}
		// 创建zip文件中的文件或目录
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath
		if info.IsDir() {
			// 创建目录
			header.Name += "/"
			_, err = zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
		} else {
			// 创建文件
			fileWriter, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			// 打开源文件
			file, err := os.Open(path)

			if err != nil {
				return err
			}
			defer file.Close()
			// 将源文件内容复制到zip文件中
			_, err = io.Copy(fileWriter, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func RemoveFiles(sourceDir string) error {
	// 遍历源目录
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 删除文件
		if !info.IsDir() && !EndsWith(info.Name(), ".zip") {
			err = os.Remove(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func EndsWith(str, suffix string) bool {
	if len(str) < len(suffix) {
		return false
	}
	return str[len(str)-len(suffix):] == suffix
}
