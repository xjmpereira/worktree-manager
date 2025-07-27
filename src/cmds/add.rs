use clap::Args;
use std::ffi::OsString;

/// Add a worktree from the WM managed repository
#[derive(Debug, Args)]
#[command(arg_required_else_help = false)]
pub(crate) struct AddCommand {
    #[arg(required = false)]
    branch: Option<OsString>,
}

pub(crate) fn subcommand(_args: AddCommand) {
    panic!("Subcommand NOT implemented");
}
