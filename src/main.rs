use clap::Parser;

mod cmds;

fn main() {
    cmds::run(cmds::Cli::parse());
}
