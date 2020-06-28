# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased]


## [1.4.0] - 2020-06-28
### Added
* Support for a "Superuser" defined in config that can use all commands
* Ability to define an additional "bot users" role

### Fixed
* Long lists of names are now split to avoid exceeding character limits
* Invalid mentions passed to .nuke and .restore no longer cause a crash
* Direct messages sent to the bot don't crash it any more


## [1.3.0] - 2020-06-18
### Added
* Quarantine role is now able to be configured per-guild

### Fixed
* Phrasing of quarantine errors now uses correct tense


## [1.2.1] - 2020-06-18
### Changed
* Message logs in console now include guild and channel names

### Fixed
* No longer displays extraneous nick-saving output when nick didn't change
* Quarantine role no longer gets removed instantly while running .nuke
* Members command will always show member count even if name list exceeds char limit


## [1.2.0] - 2020-06-13
### Added
* Guild stats now get printed when loading bot or joining servers
* First-run setup for guilds to save initial member names and guild config


## [1.1.0] - 2020-06-04
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


[Unreleased]: https://github.com/dshoreman/smegbot/compare/v1.4.0...develop
[1.4.0]: https://github.com/dshoreman/smegbot/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/dshoreman/smegbot/compare/v1.2.1...v1.3.0
[1.2.1]: https://github.com/dshoreman/smegbot/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/dshoreman/smegbot/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/dshoreman/smegbot/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/dshoreman/smegbot/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/dshoreman/smegbot/releases/tag/v1.0.0
