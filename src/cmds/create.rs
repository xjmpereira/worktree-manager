use clap::Args;
use std::ffi::OsString;

/// Create a worktree and add it to the WM managed repository
#[derive(Debug, Args)]
#[command(arg_required_else_help = false)]
pub(crate) struct CreateCommand {
    #[arg(required = false)]
    branch: Option<OsString>,
}

pub(crate) fn subcommand(_args: CreateCommand) {
    panic!("Subcommand NOT implemented");
}
