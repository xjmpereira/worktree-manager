use clap::Args;
use std::ffi::OsString;

/// Remove a worktree from the WM managed repository
#[derive(Debug, Args)]
#[command(arg_required_else_help = false)]
pub(crate) struct RmCommand {
    #[arg(required = false)]
    branch: Option<OsString>,
}

pub(crate) fn subcommand(_args: RmCommand) {
    panic!("Subcommand NOT implemented");
}
