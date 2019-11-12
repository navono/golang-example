package gogit

import (
	"fmt"
	"time"

	"github.com/urfave/cli"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
	"gopkg.in/src-d/go-git.v4/storage/memory"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "git",
		Aliases: []string{"git"},

		Usage: "demonstration of go-git",
		Subcommands: []cli.Command{
			{
				Name:   "clone",
				Usage:  "basic git clone example",
				Action: cloneAction,
			}, {
				Name:   "init",
				Usage:  "basic git init example",
				Action: initAction,
			}, {
				Name:   "commit",
				Usage:  "basic git commit example",
				Action: commitAction,
			}, {
				Name:   "reset",
				Usage:  "basic git reset example",
				Action: resetAction,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "hash",
						Required: true,
					},
				},
			},
		},
		Category: "Git",
	})
}

const path = "F:\\test-data"

// Example of how to:
// - Clone a repository into memory
// - Get the HEAD reference
// - Using the HEAD reference, obtain the commit this reference is pointing to
// - Using the commit, obtain its history and print it
func cloneAction(c *cli.Context) error {

	// Clones the given repository, creating the remote, the local branches
	// and fetching the objects, everything in memory:
	Info("git clone https://github.com/src-d/go-siva")

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/src-d/go-siva",
	})
	CheckIfError(err)

	// Gets the HEAD history from HEAD, just like this command:
	Info("git log")

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)

	// ... retrieves the commit history
	// since := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	// until := time.Date(2019, 7, 30, 0, 0, 0, 0, time.UTC)
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})
	CheckIfError(err)

	return err
}

func initAction(c *cli.Context) error {
	gitFs := osfs.New(path)
	dot, _ := gitFs.Chroot(".git")

	storage := filesystem.NewStorage(dot, cache.NewObjectLRUDefault())
	repo, err := git.Init(storage, gitFs)
	if err != nil {
		return err
	}

	w, _ := repo.Worktree()
	_, err = w.Add(".")
	if err != nil {
		return err
	}

	return nil
}

func commitAction(c *cli.Context) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	wt, err := repo.Worktree()
	if err != nil {
		return err
	}

	_, err = wt.Add(".")
	if err != nil {
		return err
	}

	h, err := wt.Commit("commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "user1",
			Email: "",
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	fmt.Printf("commit hash %s", h.String())

	return nil
}

func resetAction(c *cli.Context) error {
	h := c.String("hash")
	fmt.Printf("reset hash: %s", h)

	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	wt, err := repo.Worktree()
	if err != nil {
		return err
	}

	return wt.Reset(&git.ResetOptions{
		Commit: plumbing.NewHash(h),
		Mode:   git.HardReset,
	})
}
