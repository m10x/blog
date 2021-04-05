package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

func main() {

	strTheme := "hello-friend-ng"
	strThemePath := "themes/" + strTheme

	fmt.Println("----- 0) Updating theme")

	out, err := exec.Command("git", "-C", strThemePath, "fetch", "origin").Output()
	if err != nil {
		log.Fatal("git fetch: ", err)
	}
	fmt.Println("Fetch:", string(out))

	out, err = exec.Command("git", "-C", strThemePath, "reset", "--hard", "origin/master").Output()
	if err != nil {
		log.Fatal("git reset: ", err)
	}
	fmt.Println("Reset:", string(out))

	fmt.Println("----- 0) Finished")

	fmt.Println("----- 1) Removing everything in public but dont remove CNAME")

	dir, err := ioutil.ReadDir("public")
	for _, d := range dir {
		if d.Name() != "CNAME" {
			os.RemoveAll(path.Join([]string{"public", d.Name()}...))
		}
	}

	fmt.Println("----- 1) Finished")

	fmt.Println("----- 2) Running hugo")

	out, err = exec.Command("./hugo", "-t", strTheme).Output()
	if err != nil {
		log.Fatal("hugo: ", err)
	}
	fmt.Println("Hugo:", string(out))

	fmt.Println("----- 2) Finished")

	fmt.Println("----- 3) Pushing changes")

	/* max@max-pc MINGW64 ~/Documents/git/blog/m10x-blog/public (master)
	$ git add *
	fatal: in unpopulated submodule 'm10x-blog/public'


	git submodule add git@github.com:m10x/m10x.github.io.git public
	*/

	out, err = exec.Command("git", "-C", "public/", "add", "-A").Output()
	if err != nil {
		log.Fatal("git add: ", err)
	}
	fmt.Println("Git Add:", string(out))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a commit message: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("enter message: ", err)
	}

	out, err = exec.Command("git", "-C", "public", "commit", "-m", text).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Git Commit:", string(out))

	out, err = exec.Command("git", "-C", "public", "push").Output()
	if err != nil {
		log.Fatal("git push: ", (err))
	}
	fmt.Println("Git Push:", string(out))

	fmt.Println("----- 3) Finished")
}
