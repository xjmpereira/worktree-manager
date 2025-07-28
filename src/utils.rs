use anyhow::{anyhow, Result};
use regex::Regex;
use std::option::Option;
use std::str::FromStr;
use std::string::String;

#[derive(Debug)]
pub enum GitUriScheme {
    Ssh,
    File,
    Http,
    Https,
    Ftp,
    Ftps,
}

impl GitUriScheme {
    pub fn as_str(&self) -> &'static str {
        match self {
            GitUriScheme::Ssh => "ssh://",
            GitUriScheme::File => "file://",
            GitUriScheme::Http => "http://",
            GitUriScheme::Https => "https://",
            GitUriScheme::Ftp => "ftp://",
            GitUriScheme::Ftps => "ftps://",
        }
    }
}

impl FromStr for GitUriScheme {
    type Err = anyhow::Error;
    fn from_str(input: &str) -> Result<GitUriScheme, Self::Err> {
        match input {
            "ssh://" => Ok(GitUriScheme::Ssh),
            "file://" => Ok(GitUriScheme::File),
            "http://" => Ok(GitUriScheme::Http),
            "https://" => Ok(GitUriScheme::Https),
            "ftp://" => Ok(GitUriScheme::Ftp),
            "ftps://" => Ok(GitUriScheme::Ftps),
            _ => Err(anyhow!("Invalid scheme.")),
        }
    }
}

#[derive(Debug)]
pub struct GitUri {
    pub uri: String,
}

impl FromStr for GitUri {
    type Err = anyhow::Error;
    fn from_str(input: &str) -> anyhow::Result<GitUri, Self::Err> {
        Ok(GitUri {
            uri: input.to_string(),
        })
    }
}
impl GitUri {
    pub fn parse(&self) -> anyhow::Result<GitUriDetails> {
        let re = Regex::new(r"^(?P<scheme>git@|git://|file://|ssh://|ftps?://|https?://)?((?P<user>[a-zA-Z0-9]\w*)(?P<password>:.*)?@)?((?P<domain>[a-zA-Z0-9][\w\.]*)(?P<port>:[0-9]*)?)?(?P<path>(\.?/?~?|~).*)/(?P<name>[\w\.-]+(\.git)?)/?$").unwrap();
        let groups = re.captures(self.uri.as_str()).unwrap();

        let scheme: GitUriScheme = GitUriScheme::from_str(&groups["scheme"])?;
        let domain: GitUriDomain = GitUriDomain::from_str(&groups["domain"])?;
        let user: Option<GitUriUser> = match groups.name("user") {
            Some(val) => GitUriUser::from_str(val.as_str()).ok(),
            None => None,
        };
        let password: Option<GitUriPassword> = match groups.name("password") {
            Some(val) => GitUriPassword::from_str(val.as_str().strip_prefix(":").unwrap()).ok(),
            None => None,
        };
        let port: Option<GitUriPort> = match groups.name("port") {
            Some(val) => GitUriPort::from_str(val.as_str().strip_prefix(":").unwrap()).ok(),
            None => None,
        };
        let path: GitUriPath = GitUriPath::from_str(&groups["path"]).unwrap();
        let name: GitUriName =
            GitUriName::from_str(groups["name"].strip_suffix(".git").unwrap()).unwrap();

        Ok(GitUriDetails {
            scheme,
            domain,
            user,
            password,
            port,
            path,
            name,
        })
    }
}

pub type GitUriDomain = String;
pub type GitUriUser = String;
pub type GitUriPassword = String;
pub type GitUriPort = u32;
pub type GitUriPath = String;
pub type GitUriName = String;

#[derive(Debug)]
pub struct GitUriDetails {
    pub scheme: GitUriScheme,
    pub domain: GitUriDomain,
    pub user: Option<GitUriUser>,
    pub password: Option<GitUriPassword>,
    pub port: Option<GitUriPort>,
    pub path: GitUriPath,
    pub name: GitUriName,
}

// pub fn ls_remote(remote: String) {

// let repo = Repository::open(".")?;
// let mut remote = repo
//     .find_remote(remote.to_str().unwrap())
//     .or_else(|_| repo.remote_anonymous(remote.to_str().unwrap()))?;

// // Connect to the remote and call the printing function for each of the
// // remote references.
// let connection = remote.connect_auth(Direction::Fetch, None, None)?;

// // Get the list of references on the remote and print out their name next to
// // what they point to.
// for head in connection.list()?.iter() {
//     println!("{}\t{}", head.oid(), head.name());
// }
// Ok(())
// }
