# gh-backup

A simple CLI utility, written in GoLang, to backup all of your public GitHub repos.

# Installing gh-backup

There are a few options for installing gh-backup. You can use [Homebrew](https://brew.sh) (recommended), you can install it from the "releases" page here or you can build it manually.

## Install from Homebrew

Installing from [Homebrew](https://brew.sh) is the recommended approach on both Mac and Linux as it allows for automatic updates when needed. To install from Homebrew run the following command:

```
brew install ChrisWiegman/gh-backup/gh-backup
```

## Download from GitHub releases

Simply download the latest release from our [release page](https://github.com/ChrisWiegman/gh-backup/releases) and extract the CLI to a location accessible by your system PATH

**Note for Mac users** I have not signed the download copy so you'll need to manually allow it in your Mac settings if you download it from the releases page. Install it via Homebrew to avoid this step.

## Build manually

You will need [Go](https://go.dev) installed locally to build the application for now.

1. Clone this repo `git clone git@github.com:ChrisWiegman/gh-backup.git`
2. CD into the repo and run `make install`

Assuming you have Go properly setup with `GOBIN` in your system path, you should now be able to use Kana. Run `gh-backup --version` to test.

# Using

Run the backup in the directory where you want your backups to live

```gh-backup <my GitHub username>```
