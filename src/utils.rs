// pub(crate) fn ls_remote(remote: OsString) -> Result<(), git2::Error> {
//     let repo = Repository::open(".")?;
//     let mut remote = repo
//         .find_remote(remote.to_str().unwrap())
//         .or_else(|_| repo.remote_anonymous(remote.to_str().unwrap()))?;

//     // Connect to the remote and call the printing function for each of the
//     // remote references.
//     let connection = remote.connect_auth(Direction::Fetch, None, None)?;

//     // Get the list of references on the remote and print out their name next to
//     // what they point to.
//     for head in connection.list()?.iter() {
//         println!("{}\t{}", head.oid(), head.name());
//     }
//     Ok(())
// }
