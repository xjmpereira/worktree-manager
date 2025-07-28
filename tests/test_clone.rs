#![allow(missing_docs)]

use assert_cmd::Command;
use random_string::{charsets::ALPHA_LOWER, generate};
use std::fs::{self, read, DirEntry};
use std::path::Path;
use tempdir::TempDir;
use worktree_manager::utils;

#[test]
fn clone_non_existing_dir() {
    let uri = utils::GitUri {
        uri: String::from("https://github.com/xjmpereira/test-worktree-manager.git"),
    };
    let details = uri.parse().unwrap();
    println!("{:?}", details);

    // Prepare directory that gets deleted at end of test
    let tmp_dirname = format!("wm-{}", generate(8, ALPHA_LOWER));
    let tmp_dir = TempDir::new(tmp_dirname.as_str()).unwrap();
    // Prepare path test where repository will be downloaded to
    let repo_dir = tmp_dir.path().join(details.name.clone());
    let tmp_path = repo_dir.as_os_str().to_str().unwrap();

    // Actually run the CLI
    Command::cargo_bin("worktree-manager")
        .unwrap()
        .args([
            "clone",
            "https://github.com/xjmpereira/test-worktree-manager.git",
            tmp_path,
        ])
        .assert()
        .success();
    let root_files = fs::read_dir(tmp_path)
        .unwrap()
        .map(|entry| entry.unwrap())
        .collect::<Vec<DirEntry>>();

    // Test that the expected files exist in the remote directory
    let license_file = Path::new(&details.name.clone()).join("LICENSE");
    let readme_file = Path::new(&details.name.clone()).join("README.md");
    let git_dir = Path::new(&details.name.clone()).join(".git");
    let mut expected: Vec<&Path> = vec![&license_file, &readme_file, &git_dir];
    for file in root_files {
        let file_path = file.path();
        let stripped_path = file_path.strip_prefix(tmp_dir.path()).unwrap();
        let index = expected.iter().position(|&path| path == stripped_path);
        match index {
            Some(idx) => {
                println!("Found file {:?}", stripped_path);
                expected.remove(idx);
            }
            None => {
                panic!("File {:?} not found in {:?}", file_path, expected);
            }
        }
    }
}
