use clap::Args;
use std::ffi::OsString;

/// Change the current directory to a specific worktree
#[derive(Debug, Args)]
#[command(arg_required_else_help = false)]
pub(crate) struct CdCommand {
    #[arg(required = false)]
    pub(crate) branch: Option<OsString>,
}

pub(crate) fn subcommand(_args: CdCommand) {
    panic!("Subcommand NOT implemented");
}
