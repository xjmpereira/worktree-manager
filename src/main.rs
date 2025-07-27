use clap::Parser;

mod cmds;
mod utils;

fn main() {
    cmds::run(cmds::Cli::parse());
}
