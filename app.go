package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

func init() {
	app.Commands = []*cli.Command{
		{
			Name:   "init",
			Usage:  "init password manager",
			Action: initDB,
			Flags:  []cli.Flag{},
		},
		{
			Name:   "add",
			Usage:  "add new password",
			Action: add,
			Flags: []cli.Flag{
				// required
				&cli.StringFlag{
					Name:     "name",
					Required: true,
					Usage:    "name for password",
				},
				&cli.StringFlag{
					Name:     "loginid",
					Required: true,
					Usage:    "login id",
				},
				&cli.StringFlag{
					Name:     "password",
					Required: true,
					Usage:    "password",
				},
				// optional
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
				&cli.StringFlag{
					Name:  "bind",
					Usage: "bind email/tel/third party account",
				},
				&cli.StringFlag{
					Name:  "comment",
					Usage: "notes for password",
				},
				&cli.StringFlag{
					Name:  "alias",
					Usage: "alias for name",
				},
				&cli.StringFlag{
					Name:  "site",
					Usage: "witch site to use this password",
				},
			},
		},
		{
			Name:   "get",
			Usage:  "get password",
			Action: get,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
			},
		},
		{
			Name:   "list",
			Usage:  "list all password",
			Action: list,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "update password",
			Action: update,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "name for password",
				},
				&cli.StringFlag{
					Name:  "loginid",
					Usage: "login id",
				},
				&cli.StringFlag{
					Name:  "password",
					Usage: "password",
				},
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
				&cli.StringFlag{
					Name:  "bind",
					Usage: "bind email/tel/third party account",
				},
				&cli.StringFlag{
					Name:  "comment",
					Usage: "notes for password",
				},
				&cli.StringFlag{
					Name:  "alias",
					Usage: "alias for name",
				},
				&cli.StringFlag{
					Name:  "site",
					Usage: "witch site to use this password",
				},
			},
		},
		{
			Name:   "remove",
			Usage:  "remove password",
			Action: remove,
		},
		{
			Name:   "import",
			Usage:  "import passwords from csv, the csv file format: name,loginId,password,bind,alias,site,comment",
			Action: importFromCsv,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
			},
		},
		{
			Name:   "export",
			Usage:  "export to csv",
			Action: export2Csv,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
				&cli.PathFlag{
					Name:  "output",
					Usage: "output file",
				},
			},
		},
		{
			Name:   "decrypt",
			Usage:  "decrypt all",
			Action: decryptAll,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:     "key",
					Usage:    "path to key file",
					Required: true,
				},
			},
		},
		{
			Name:   "encrypt",
			Usage:  "encrypt all",
			Action: encryptAll,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:     "key",
					Usage:    "path to key file",
					Required: true,
				},
			},
		},
		{
			Name:   "cover",
			Usage:  "encrypt local file",
			Action: cover,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:     "key",
					Usage:    "path to key file",
					Required: true,
				},
			},
		},
		{
			Name:   "decover",
			Usage:  "decrypt local file",
			Action: decover,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:     "key",
					Usage:    "path to key file",
					Required: true,
				},
			},
		},
		{
			Name:   "gen",
			Usage:  "generate password",
			Action: generate,
		},
	}

	app.Version = VERSION
	app.Usage = "password manager"
	app.EnableBashCompletion = true
	app.Name = "gokeeper"
	app.Authors = []*cli.Author{{Name: "Alctny", Email: "ltozvxe@gmail.com"}}
}

// create sqlite file and create table
func initDB(ctx *cli.Context) error {
	db.AutoMigrate(&Password{})
	return db.Error
}

// add new password
func add(ctx *cli.Context) error {
	keyFile := ctx.String("key")
	name := ctx.String("name")
	loginId := ctx.String("loginid")
	password := ctx.String("password")
	bind := ctx.String("bind")
	alias := ctx.String("alias")
	site := ctx.String("site")
	comment := ctx.String("comment")
	isEncrype := 0

	var err error
	if keyFile != "" {
		isEncrype = 1

		password, err = EncrypeString(password, keyFile)
		if err != nil {
			return err
		}

		loginId, err = EncrypeString(loginId, keyFile)
		if err != nil {
			return err
		}

		if bind != "" {
			bind, err = EncrypeString(bind, keyFile)
			if err != nil {
				return err
			}
		}
	}

	passwd := Password{
		Name:     name,
		LoginId:  loginId,
		Password: password,
		Bind:     bind,
		Alias:    alias,
		Site:     site,
		Comment:  comment,
		Encrypt:  isEncrype,
	}

	tx := db.Create(&passwd)
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println("add password success")
	return nil
}

// get password
func get(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return fmt.Errorf("query is empty")
	}
	query := fmt.Sprintf("%%%s%%", ctx.Args().First())
	passwords := []Password{}
	tx := db.Where("name LIKE ? OR site LIKE ? OR comment LIKE ? OR alias LIKE ?", query, query, query, query).Find(&passwords)
	if tx.Error != nil {
		return tx.Error
	}

	keyFile := ctx.Path("key")
	if keyFile != "" {
		var err error
		passwords, err = DecryptPasswords(passwords, keyFile)
		if err != nil {
			return err
		}
	}
	ShowPasswords(passwords)
	return nil
}

// list all passwords
func list(ctx *cli.Context) error {
	passwords := []Password{}
	keyFile := ctx.String("key")

	tx := db.Find(&passwords)
	if tx.Error != nil {
		return tx.Error
	}
	if keyFile != "" {
		var err error
		passwords, err = DecryptPasswords(passwords, keyFile)
		if err != nil {
			return err
		}
	}

	ShowPasswords(passwords)
	return nil
}

// update password infomaion
// this method will not attempt to automatically update the encryption status
func update(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return fmt.Errorf("id is empty")
	}
	idStr := ctx.Args().First()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	name := ctx.String("name")
	loginId := ctx.String("loginid")
	password := ctx.String("password")
	bind := ctx.String("bind")
	alias := ctx.String("alias")
	site := ctx.String("site")
	comment := ctx.String("comment")
	isEncrype := 0

	keyFile := ctx.Path("key")
	if keyFile != "" {
		isEncrype = 1
		if password != "" {
			password, err = EncrypeString(password, keyFile)
			if err != nil {
				return err
			}
		}

		if loginId != "" {
			loginId, err = EncrypeString(loginId, keyFile)
			if err != nil {
				return err
			}
		}

		if bind != "" {
			bind, err = EncrypeString(bind, keyFile)
			if err != nil {
				return err
			}
		}
	}

	p := Password{
		Id:       id,
		Name:     name,
		LoginId:  loginId,
		Password: password,
		Bind:     bind,
		Alias:    alias,
		Site:     site,
		Comment:  comment,
		Encrypt:  isEncrype,
	}
	// 此处必须添加 encrypt = isEncrype 条件，保证修改前后加密状态未发生改变
	// 避免出现所有信息都没有加密但 encrypt = 1 或所有信息都有加密但 encrypt = 0 的情况
	tx := db.Model(&Password{}).Where("id = ? AND encrypt = ?", id, isEncrype).Updates(p)
	if tx.Error != nil {
		return tx.Error
	}

	fmt.Println("update password success")
	return nil
}

// remove password
func remove(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return fmt.Errorf("id is empty")
	}
	idStr := ctx.Args().First()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	tx := db.Delete(&Password{Id: id})
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println("remove password success")
	return nil
}

// import passwords from csv, csv format(no header)):
// name,loginId,password,bind,alias,site,comment
func importFromCsv(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return fmt.Errorf("csv file is empty")
	}
	csvFile := ctx.Args().First()
	f, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	if len(data) < 1 {
		return nil
	}

	keyFile := ctx.Path("key")

	if len(data[0]) < 7 {
		return fmt.Errorf("csv format error, should be 'name,loginId,password,bind,alias,site,comment', and no header")
	}
	passwords := []Password{}
	for _, line := range data[1:] {
		p := Password{
			Name:     line[0],
			LoginId:  line[1],
			Password: line[2],
			Bind:     line[3],
			Alias:    line[4],
			Site:     line[5],
			Comment:  line[6],
		}
		if keyFile != "" {
			p.Encrypt = 1
			p.Password, err = EncrypeString(p.Password, keyFile)
			if err != nil {
				return err
			}
			p.LoginId, err = EncrypeString(p.LoginId, keyFile)
			if err != nil {
				return err
			}
			p.Bind, err = EncrypeString(p.Bind, keyFile)
			if err != nil {
				return err
			}
		}
		passwords = append(passwords, p)
	}
	tx := db.Create(&passwords)
	if err != nil {
		return err
	}
	fmt.Println("import passwords success: ", tx.RowsAffected)
	return nil
}

// export all passwords to csv
// TODO: export to csv
func export2Csv(ctx *cli.Context) error {
	keyFile := ctx.Path("key")
	passwords := []Password{}
	if keyFile != "" {
		tx := db.Find(&passwords)
		if tx.Error != nil {
			return tx.Error
		}
		var err error
		passwords, err = DecryptPasswords(passwords, keyFile)
		if err != nil {
			return err
		}
	} else {
		tx := db.Where("encrypt = 0").Find(&passwords)
		if tx.Error != nil {
			return tx.Error
		}
	}

	outputFile := ctx.Path("output")
	if outputFile == "" {
		outputFile = "output.csv"
	}
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()
	csvWriter := csv.NewWriter(f)
	defer csvWriter.Flush()
	for _, p := range passwords {
		csvWriter.Write([]string{p.Name, p.LoginId, p.Password, p.Bind, p.Alias, p.Site, p.Comment})
	}
	fmt.Println("export passwords success to: ", outputFile)
	return nil
}

// decrypt all passwors, loginId, bind
func decryptAll(ctx *cli.Context) error {
	keyFile := ctx.Path("key")
	passwords := []Password{}
	tx := db.Where(&Password{Encrypt: 1}).Find(&passwords)
	if tx.Error != nil {
		return tx.Error
	}
	var err error
	passwords, err = DecryptPasswords(passwords, keyFile)
	if err != nil {
		return err
	}
	tx = db.Begin()
	for _, p := range passwords {
		result := tx.Save(&p)
		if result.Error != nil {
			return result.Error
		}
	}
	tx = tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println("decrypt success")

	return nil
}

// encrypt all passwors, loginId, bind
func encryptAll(ctx *cli.Context) error {
	keyFile := ctx.Path("key")
	passwords := []Password{}
	tx := db.Where("encrypt = 0").Find(&passwords)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Begin()
	var err error
	for _, p := range passwords {
		p.Password, err = EncrypeString(p.Password, keyFile)
		if err != nil {
			return err
		}

		p.LoginId, err = EncrypeString(p.LoginId, keyFile)
		if err != nil {
			return err
		}

		p.Bind, err = EncrypeString(p.Bind, keyFile)
		if err != nil {
			return err
		}
		p.Encrypt = 1
		result := tx.Save(p)
		if result.Error != nil {
			return result.Error
		}
	}
	tx = tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println("encrypt success")

	return nil
}

// generate password
// 生成长度固定为 13 同时包含小写字母，大写字母，数字，特殊字符的随机字符串
func generate(ctx *cli.Context) error {
	const passwordLen = 13
	sed := rand.New(rand.NewSource(time.Now().UnixNano()))
	charSet := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789",
		"!@#$%^&*()_+",
	}

	// 生成 4 个非 0 随机数，和为 13
	charSetLen := len(charSet)
	charCont := make([]int, charSetLen)
	var sum = passwordLen
	for i := 0; i < charSetLen; i++ {
		// charSetLen-1-i 待生成的随机数个数 +1 保证生成的数不为 0
		charCont[i] = sed.Intn(sum-(charSetLen-1-i)) + 1
		sum -= charCont[i]
	}

	// 按照生成的随机数依次使用小写，大写，数字，特殊字符进行填充
	passwords := make([]byte, passwordLen)
	v := 0
	for i := 0; i < charSetLen; i++ {
		length := len(charSet[i])
		for j := 0; j < charCont[i]; j++ {
			passwords[v] = charSet[i][sed.Intn(length)]
			v++
		}
	}

	// 打乱密码顺序，进一步提高随机性，非必要
	for i := 0; i < 6; i++ {
		ri := sed.Intn(13)
		passwords[i], passwords[ri] = passwords[ri], passwords[i]
	}

	fmt.Println("gen: ", string(passwords))
	return nil
}

func cover(ctx *cli.Context) error { return nil }

func decover(ctx *cli.Context) error { return nil }
