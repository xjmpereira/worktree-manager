#![allow(missing_docs)]

use assert_cmd::Command;

const HELP_STR: &str = "A git worktree manager tool

Usage: worktree-manager [COMMAND]

Commands:
  clone   Clone a repository and prepare the folder as a WM managed repository
  list    List all the worktrees on the WM managed repository
  create  Create a worktree and add it to the WM managed repository
  add     Add a worktree from the WM managed repository
  rm      Remove a worktree from the WM managed repository
  cd      Change the current directory to a specific worktree
  help    Print this message or the help of the given subcommand(s)

Options:
  -h, --help  Print help
";

#[test]
fn help() {
    Command::cargo_bin("worktree-manager")
        .unwrap()
        .args(["--help"])
        .assert()
        .success()
        .stdout(HELP_STR);
}
