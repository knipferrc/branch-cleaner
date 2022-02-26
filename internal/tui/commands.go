package tui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type errorMsg error
type repoDataMsg []list.Item

func readCurrentGitBranchesCmd() tea.Cmd {
	return func() tea.Msg {
		r, err := git.PlainOpen(git.GitDirName)
		if err != nil {
			return errorMsg(err)
		}

		branches, err := r.Branches()
		if err != nil {
			return errorMsg(err)
		}

		var branchNames []*plumbing.Reference
		err = branches.ForEach(func(b *plumbing.Reference) error {
			branchNames = append(branchNames, b)

			return nil
		})

		if err != nil {
			return errorMsg(err)
		}

		var items []list.Item

		for _, branch := range branchNames {
			items = append(items, item{
				title: branch.Name().Short(),
				desc:  branch.Hash().String(),
			})
		}

		return repoDataMsg(items)
	}
}

func deleteSelectedBranchCmd(branchName string) tea.Cmd {
	return func() tea.Msg {
		r, err := git.PlainOpen(git.GitDirName)
		if err != nil {
			return errorMsg(err)
		}

		log.Output(2, "here")

		headRef, err := r.Head()
		if err != nil {
			return errorMsg(err)
		}

		referenceName := plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branchName))
		ref := plumbing.NewHashReference(referenceName, headRef.Hash())
		err = r.Storer.SetReference(ref)
		if err != nil {
			return errorMsg(err)
		}

		err = r.Storer.RemoveReference(ref.Name())
		if err != nil {
			return errorMsg(err)
		}

		return nil
	}
}
