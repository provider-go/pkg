package file

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// ReadLines reads all lines of the file.
// An error is returned if the specified file does not exist.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ReadLinesV2 reads all lines of the file.
// An error is returned if the specified file does not exist.
func ReadLinesV2(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	r := bufio.NewReader(file)
	for {
		// ReadString reads until the first occurrence of delim in the input,
		// returning a string containing the data up to and including the delimiter.
		line, err := r.ReadString('\n')
		if err == io.EOF {
			lines = append(lines, line)
			break
		}
		if err != nil {
			return lines, err
		}
		lines = append(lines, line[:len(line)-1])
	}
	return lines, nil
}

// ReadLinesV3 reads all lines of the file.
// An error is returned if the specified file does not exist.
func ReadLinesV3(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	r := bufio.NewReader(f)
	for {
		// ReadLine is a low-level line-reading primitive.
		// Most callers should use ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
		bytes, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return lines, err
		}
		lines = append(lines, string(bytes))
	}
	return lines, nil
}

// ListDir 列出指定目录中的所有文件目录名，该方法不会递归遍历。
func ListDir(dir string) ([]string, error) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(infos))
	for i, info := range infos {
		names[i] = info.Name()
	}
	return names, nil
}

// ListFilenames 列出指定目录中的所有文件名。
func ListFilenames(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var names []string
	for i := 0; i < len(entries); i++ {
		if !entries[i].IsDir() {
			names = append(names, entries[i].Name())
		}
	}
	return names, nil
}

// FileMd5 获取文件MD5。
func FileMd5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return FileMd5Reader(file)
}

// FileMd5Reader 从io.Reader获取文件MD5。
func FileMd5Reader(r io.Reader) (string, error) {
	hash := md5.New()
	_, err := io.Copy(hash, r)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// FileSize 获取以字节为单位的文件大小。
func FileSize(path string) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return FileSizeFile(file)
}

// FileSizeFile 从操作系统获取文件大小。文件（以字节为单位）。
func FileSizeFile(file *os.File) (int64, error) {
	info, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// ListDirEntryPaths 递归地列出目录中的所有文件或目录路径。
// 如果cur为true，则结果将包括当前目录。
// 请注意，如果子目录是符号链接，则GetDirAllEntryPaths不会跟随符号链接。
func ListDirEntryPaths(dir string, cur bool) ([]string, error) {
	// Remove the trailing path separator if dirname has.
	dir = strings.TrimSuffix(dir, string(os.PathSeparator))

	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	paths := make([]string, 0, len(infos))

	// Include current directory.
	if cur {
		paths = append(paths, dir)
	}

	for _, info := range infos {
		path := dir + string(os.PathSeparator) + info.Name()
		if info.IsDir() {
			tmp, err := ListDirEntryPaths(path, cur)
			if err != nil {
				return nil, err
			}
			paths = append(paths, tmp...)
			continue
		}
		paths = append(paths, path)
	}
	return paths, nil
}

// ListDirEntryPathsSymlink lists all the file or dir paths in the directory recursively.
// If the cur is true result will include current directory.
func ListDirEntryPathsSymlink(dirname string, cur bool) ([]string, error) {
	// Remove the trailing path separator if dirname has.
	dirname = strings.TrimSuffix(dirname, string(os.PathSeparator))

	infos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	paths := make([]string, 0, len(infos))

	// Include current directory.
	if cur {
		paths = append(paths, dirname)
	}

	for _, info := range infos {
		path := dirname + string(os.PathSeparator) + info.Name()
		realInfo, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		if realInfo.IsDir() {
			tmp, err := ListDirEntryPathsSymlink(path, cur)
			if err != nil {
				return nil, err
			}
			paths = append(paths, tmp...)
			continue
		}
		paths = append(paths, path)
	}
	return paths, nil
}
