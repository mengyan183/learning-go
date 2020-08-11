package file_test

import (
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestReadFile(t *testing.T) {
	fi, e := os.Stat("./file_test.go")
	if e != nil {
		t.Error(e)
	} else {
		// 文件的权限
		fileMode := fi.Mode()
		t.Log(fileMode.String())
	}
}

func TestCreateFileAndDir(t *testing.T) {
	// 创建新的文件夹权限为0777
	err := os.Mkdir("new_dir", os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	// 创建一个文件
	create, err := os.Create("new_dir/t.txt")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(create)
	}
}

func TestReadAndWriteFile(t *testing.T) {
	// 以只读方式读取的文件
	open, err := os.Open("new_dir/t.txt")
	if err != nil {
		t.Error(err)
	} else {
		defer open.Close()
		t.Log(open.Name())
		// 对于已只读方式读取的文件进行写入操作,会抛出异常
		writeString, err := open.WriteString("写入数据")
		if err != nil {
			t.Error(err)
		} else {
			t.Log(writeString)
		}
	}
	// 以读写方式读取文件
	file, err := os.OpenFile("new_dir/t.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		t.Error(err)
	} else {
		defer file.Close()
		writeString, err := file.WriteString("以读写方式写入数据")
		if err != nil {
			t.Error(err)
		} else {
			t.Log(writeString)
		}
	}
	// 删除文件
	err = os.Remove("new_dir/t.txt")
	if err != nil {
		t.Error(err)
	}
	// 删除空的文件夹
	err = os.Remove("new_dir")
	if err != nil {
		t.Error(err)
	}
	//删除文件夹以及文件夹下的所有文件
	err = os.RemoveAll("new_dir")
	if err != nil {
		t.Error(err)
	}
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandString 生成随机字符串
func RandString(len int) string {
	byteArray := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		byteArray[i] = byte(b)
	}
	return string(byteArray)
}

func TestReadAndWrite(t *testing.T) {
	file, err := os.OpenFile("read_write.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Error(err)
	} else {
		defer file.Close()
		for i := 0; i < 10; i++ {
			// 写入数据; 输入写入默认从头开始覆盖写入,只有在os.O_APPEND模式下才为数据追加写入的方式
			writeString, err := file.WriteString(RandString(5) + "\n")
			if err != nil {
				t.Error("写入数据报错", err)
			} else {
				t.Log("写入数据字节数", writeString)
			}
		}
		file.Close()
		file, err = os.OpenFile("read_write.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			t.Error(err)
		} else {
			defer file.Close()
			// 读取数据,对于数据读取要求,必须在数据写入释放后才能重新执行数据读取操作
			// 按照字节进行数据读取
			// 定义切片
			readDataSlice := make([]byte, 1024, 1024)
			read, err := file.Read(readDataSlice)
			if err != nil {
				t.Error("数据读取报错,", err)
			} else {
				t.Log("读取到字节数量,", read)
				// 将字节转换为字符串
				s := string(readDataSlice)
				t.Log(s)
			}
		}
	}

}

func TestCopyFile(t *testing.T) {
	// 采用边读边写的方式
	file, err := os.Open("read_write.txt")
	if err != nil {
		t.Error(err)
	} else {
		defer file.Close()
		targetFile, err := os.OpenFile("read_write_target.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			t.Error(err)
		} else {
			defer targetFile.Close()
			readData := make([]byte, 1024, 1024)
			for {
				read, err := file.Read(readData)
				if read == 0 || err != nil {
					break
				}
				_, err = targetFile.Write(readData)
				if err != nil {
					t.Error("数据写入失败", err)
				}
			}
		}
	}
}

// 使用io包下的复制功能
func TestCopyFileWithInternalFunction(t *testing.T) {
	file, err := os.Open("read_write.txt")
	if err != nil {
		t.Error(err)
	} else {
		defer file.Close()
		targetFile, err := os.OpenFile("read_write_target1.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			t.Error(err)
		} else {
			defer targetFile.Close()
			_, err = io.Copy(targetFile, file)
			if err != nil {
				t.Error("数据copy失败", err)
			}
		}
	}
}

// 使用ioUtils包进行文件复制
func TestIoUtilsCopy(t *testing.T) {
	file, err := ioutil.ReadFile("read_write.txt")
	if err != nil {
		t.Error(err)
	} else {
		err := ioutil.WriteFile("read_write_target2.txt", file, os.ModePerm)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestDeleteFile(t *testing.T) {
	_ = os.Remove("read_write.txt")
	_ = os.Remove("read_write_target.txt")
	_ = os.Remove("read_write_target1.txt")
	_ = os.Remove("read_write_target2.txt")
}
