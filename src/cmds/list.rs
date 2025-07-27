use clap::Args;

/// List all the worktrees on the WM managed repository
#[derive(Debug, Args)]
#[command(arg_required_else_help = false)]
pub(crate) struct ListCommand {}

pub(crate) fn subcommand(_args: ListCommand) {
    panic!("Subcommand NOT implemented");
}
