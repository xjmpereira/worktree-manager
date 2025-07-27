use clap::{Parser, Subcommand};

mod add;
mod cd;
mod clone;
mod create;
mod list;
mod rm;

/// A fictional versioning CLI
#[derive(Debug, Parser)] // requires `derive` feature
#[command(name = "wm")]
#[command(about = "A git worktree manager tool", long_about = None)]
pub(crate) struct Cli {
    #[command(subcommand)]
    command: Option<Commands>,
}

#[derive(Debug, Subcommand)]
pub(crate) enum Commands {
    Clone(clone::CloneCommand),
    List(list::ListCommand),
    Create(create::CreateCommand),
    Add(add::AddCommand),
    Rm(rm::RmCommand),
    Cd(cd::CdCommand),
}

pub(crate) fn run(cli: Cli) {
    // Not passing any command is the equivalent to a 'wm cd' command with not arguments
    let cmd = cli
        .command
        .unwrap_or(Commands::Cd(cd::CdCommand { branch: None }));
    println!("{cmd:?}");
    match cmd {
        Commands::Clone(clone_command) => {
            clone::subcommand(clone_command);
        }
        Commands::List(list_command) => {
            list::subcommand(list_command);
        }
        Commands::Create(create_command) => {
            create::subcommand(create_command);
        }
        Commands::Add(add_command) => {
            add::subcommand(add_command);
        }
        Commands::Rm(rm_command) => {
            rm::subcommand(rm_command);
        }
        Commands::Cd(cd_command) => {
            cd::subcommand(cd_command);
        }
    }
}
