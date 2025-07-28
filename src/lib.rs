pub mod utils;

// let tmp_dirname = format!("wm-{}", generate(8, ALPHA_LOWER));
// let tmp_dir = TempDir::new(tmp_dirname.as_str()).unwrap();
// let mut opts = RepositoryInitOptions::new();
// opts.bare(true);
// match Repository::init_opts(tmp_dir.path(), &opts) {
//     Ok(repo) => {
//         let remote = repo
//             .remote("origin", remote.as_str())
//             .unwrap();
//         let name = remote.name();
//         println!("{name:?}");
//     },
//     Err(error) => {
//         panic!("{error:?}")
//     }
// };
// let root_files = fs::read_dir(tmp_dir.path()).unwrap()
//     .map(|entry| entry.unwrap())
//     .collect::<Vec<DirEntry>>();
// println!("{:?}", root_files)
