package main

import (
	"flag"
	"fmt"
	pwd "password/HashMethod"
	"sync"
)

func pwdHash(wg *sync.WaitGroup) {
	var (
		g    string
		t    string
		s    int
		H    string
		hash string
		r    bool
		c    bool
		help bool
		pg   sync.WaitGroup
	)
	flag.StringVar(&g, "g", "", "生成密码")
	flag.StringVar(&t, "t", "", "选择加密方式, 默认md5")
	flag.BoolVar(&r, "r", false, "随机生成6长度的盐")
	flag.IntVar(&s, "s", 0, "指定盐")
	flag.StringVar(&H, "H", "", "指定hash算法")
	flag.BoolVar(&c, "c", false, "检查密码")
	flag.StringVar(&hash, "hash", "", "传入以加密的hash值")
	flag.BoolVar(&help, "help", false, "帮助")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	p := pwd.NewPwdHash()

	defer wg.Done()

	if help {
		flag.Usage()
		return
	}

	if g == "" {
		fmt.Println("密码不能为空")
		return
	}

	if c {
		if hash == "" && t == "" {
			fmt.Println("hash 检查密码, hash和加密方法不为空")
			return
		}
		pg.Add(1)
		go func() {
			if !p.CheckPassword(g, hash, t, &pg) {
				fmt.Println("密码不正确")
				return
			} else {
				fmt.Println("密码正确")
				return
			}
		}()
		pg.Wait()
		return
	}

	if r && s > 0 {
		fmt.Println("参数 r 和 参数 s 不能同时指定")
		return
	} else if !r && s == 0 {
		fmt.Println("需选择一个种加盐方式")
		return
	}

	if !r {
		s = p.SaltLength
		pg.Add(1)
		go func() {
			password, err := p.HashPassword(g, t, s, &pg)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(password)
		}()
		pg.Wait()
	}
}

func fileHash(wg *sync.WaitGroup) {
	var f string
	var t string
	var h bool
	var pg sync.WaitGroup

	flag.StringVar(&f, "f", "", "文件路径")
	flag.StringVar(&t, "t", "", "加密方法")
	flag.BoolVar(&h, "h", false, "")

	flag.Usage = func() {
		fmt.Println("如不指定加密方法默认使用md5")
		flag.PrintDefaults()
	}
	flag.Parse()

	defer wg.Done()

	file := pwd.NewFileHash()

	if h {
		flag.Usage()
	}

	if t == "" {
		t = "md5"
	}
	if f != "" {
		pg.Add(1)
		go func() {
			if md, err := file.FileMd5(f, t, &pg); err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println(md)
			}
		}()
		pg.Wait()
	} else {
		fmt.Println("请指定文件路径")
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go fileHash(&wg)
	wg.Wait()
	fmt.Println("完成")
}
