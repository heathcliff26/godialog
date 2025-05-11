[![CI](https://github.com/heathcliff26/godialog/actions/workflows/ci.yaml/badge.svg?event=push)](https://github.com/heathcliff26/godialog/actions/workflows/ci.yaml)
[![Coverage Status](https://coveralls.io/repos/github/heathcliff26/godialog/badge.svg)](https://coveralls.io/github/heathcliff26/godialog)
[![Editorconfig Check](https://github.com/heathcliff26/godialog/actions/workflows/editorconfig-check.yaml/badge.svg?event=push)](https://github.com/heathcliff26/godialog/actions/workflows/editorconfig-check.yaml)
[![Generate go test cover report](https://github.com/heathcliff26/godialog/actions/workflows/go-testcover-report.yaml/badge.svg)](https://github.com/heathcliff26/godialog/actions/workflows/go-testcover-report.yaml)
[![Renovate](https://github.com/heathcliff26/godialog/actions/workflows/renovate.yaml/badge.svg)](https://github.com/heathcliff26/godialog/actions/workflows/renovate.yaml)

# GoDialog

GoDialog is a golang API for opening OS native file dialogs on linux/windows. Additionally it allows to define a fallback implementation should the native dialog not work.

TODO:
- [x] Update all files and adapt for current use-case
- [x] Copy over files from go-minesweeper
- [x] Initialize go module
- [x] Create github repo
- [x] Update ci and app permissions
- [x] Create test app
- [x] Move fyne native implementation into separate package
- [x] Do not vendor deps?
- [x] Rename filedialog package to dodialog
- [x] Move dbus to linux dialog file
- [ ] Test with go-minesweeper
- [ ] Test infraspace-...
- [ ] Update README.md
- [ ] Release
- [ ] Use in other projects
