# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased]
### Added
* Nicks (or usernames) are now saved when a member is updated
* Join and part messages now show nickname or previous name where applicable
* New `.config` command for getting/setting guild options
* Target channel for join and part messages can now be set by admins


## [1.0.1] - 2020-06-03
### Fixed
* Compile error caused by missing brace thanks to .todo command in local copy
* Bad link to v1.0.0 tag in changelog


## [1.0.0] - 2020-06-03
### Added
* `.roles` command to list the roles a mentioned user has
* `.members` to list the members with a mentioned role
* `.nuke` and `.restore` to temporarily replace a member's roles with @Quarantine
* Welcome and leave messages when users join, leave or get kicked from the server
* Global permission checks - messages from members without Administrator are ignored


[Unreleased]: https://github.com/dshoreman/smegbot/compare/v1.0.1...develop
[1.0.1]: https://github.com/dshoreman/smegbot/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/dshoreman/smegbot/releases/tag/v1.0.0
