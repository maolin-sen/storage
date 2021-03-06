# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/)
and this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

## [v0.5.0] - 2019-12-30

### Added

- services: Add support for s3 (#41)
- services: Add basic oss support (#42)
- services: Add basic gcs support (#48)
- services: Add basic support for azblob (#50)

### Changed

- pkg/config: Allow emit host instead of credential
- pkg/credential: Implement proposal 4-credential-refactor

### Fixed

- sercices/s3: Fix error message for servicer (#44)

## [v0.4.0] - 2019-12-23

### Added

- servicer: Add String() for debug (#23)
- Implement proposal support service init via config string (#38)

### Changed

- internal: Refactor generator (#24)
- internal: Don't preserve files' metadata
- storager: Implement proposal 1-unify-storager-behavior (#30)
- *: Implement proposal 2-use-callback-in-list-operations (#31)
- services: Promote values into struct instead of metadata (#33)
- services: Split endpoint and credential into different pairs (#34)
- storager: Split Metadata to Metadata and Statistical (#39)

### Fixed

- services/posixfs: Fix std{in/out} support for Stat (#35)

## [v0.3.0] - 2019-11-11

### Added

- services: Add POSIX fs support (#1)
- services/posixfs: Add support for write stdout
- services/posixfs: Implement size and offset in Read (#8)
- services/qingstor: Add bucket name validate (#9)
- storager: Add String interface for debug print (#16)
- services/posixfs: Set updatedAt for regular file (#19)

### Changed

- storager: Merge Capablity and IsPairAvailable into Capable
- services: Unify behavior for ListDir recursively
- storager: All API now use relative path instead
- services/qingstor: Set default base value (#13)
- types: Rename Base to WorkDir (#17)

### Fixed

- services/qingstor: Fix segment ID used incorrectly
- services/posixfs: Fix ListDir not returned ErrDone
- pkg/segment: Fix data while update segments
- pkg/segment: Fix GetPartIndex bug in concurrent write segment
- pkg/segment: Return sorted parts
- services/posixfs: Fix ListDir recursive not work
- pkg/iterator: Fix next panic while under lying func return empty buf
- services/qingstor: Fix recursive not handled correctly
- services/posixfs: Fix relative path not returned correctly
- services/qingstor: Part number must be in ascending order (#14)
- services/qingstor: Fix get abs and rel path (#15)
- services/qingstor: Handle not found returned via head (#20)
- services/posixfs: Handle not found returned via os not exist error. (#21)

### Removed

- storager: Remove CreateDir (#18)

## [v0.2.0] - 2019-10-17

### Added

- services/qingstor: Implement ListSegments (#7)

### Changed

- segment: Replace identifier and add ListSegments support (#6)

### Fixed

- services/qingstor: Fix ListDir not handled correctly (#2)
- services/qingstor: Fix object size and last modified not filled (#4)
- services/qingstor: Add stat updated at support (#5)

## v0.1.0 - 2019-10-12

### Added

- Add storager and servicer interface.
- Add segment and iterator support.
- Add pair based option and metadata support.
- Add qingstor services.

[Unreleased]: https://github.com/Xuanwo/storage/compare/v0.5.0...HEAD
[v0.5.0]: https://github.com/Xuanwo/storage/compare/v0.4.0...v0.5.0
[v0.4.0]: https://github.com/Xuanwo/storage/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/Xuanwo/storage/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/Xuanwo/storage/compare/v0.1.0...v0.2.0
