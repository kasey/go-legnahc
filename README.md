## Changelog generation tool

The purpose of this tool is to enable a more automated changelog management workflow built around
a directory of changelog "fragments", which can be independently added in PRs and merged together
into the previous changelog at release time.

The repo also demonstrates using a github action together with the `changelog check` subcommand 
to automate ensuring that all PRs include a changelog fragment. 

#### What is a changelog fragment?

A changelog fragment is a tiny markdown file consisting only of section headers, matching the set 
of headers defined by the [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) format, followed by
bullet points of changelog entries that should be merged into their respect sections into
a combined release changelog. The following sections are supported:
```
### Added
- for new features.
### Changed
- for changes in existing functionality.
### Deprecated
- for soon-to-be removed features.
### Removed
- for now removed features.
### Fixed
- for any bug fixes.
### Security
- in case of vulnerabilities.

### Ignored
- see below
```

Note that in addition to the sections supported by "keep a changelog" we have added a new section called
"Ignored". This section was added because sometimes we have a valid reason for not adding to the changelog,
like a bug fix for an unreleased/incomplete feature, or a PR that reverts a previously merged but unreleased PR.
When the `Ignored` section is used, the bullet point should be a reason that this PR should be excluded
from the changelog.

### Why manage the changelog this way

There are a few benefits to using separate files over other options (simple file; deriving from pr titles):
- Standardize changelog entry style; punctuation and automatically include backrefs to pull requests.
- Reduce the possibility for human error. The release subcommand uses the range of commits between the
  previous release tag (as parsed from the previously released changelog in the repo) and the newly taggeed
  release.
- Enables github automated checks for the existence of a changelog edit.
- Avoids frustrating conflicts on the changelog itself.
- Enables more control and flexibility compared to trying to parse everything out of PR metadata. We prototyped
  that approach, but it had a few issues like single-commit PRs, PRs that needed to add multiple sections of changelog.

### Best Practices
- Put files in the `changelog` directory.
- Pick a unique name for the changelog file - an obvious choice would be `<github user name>_<branch name>.md`.
- Avoid editing changelogs; only add new files.
- The tool supports adding as many bullet points as you want to any section you want. But don't make up sections,
  unknown sections will be treated as invalid in the github workflow check.

### Using the tool to make a released changelog
- Tag release commit and make note of the exact tag string, eg `v5.3.0`.
```
$ git fetch && git checkout origin/develop
changelog release 
```
